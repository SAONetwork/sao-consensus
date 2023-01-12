package keeper

import (
	"context"
	"github.com/SaoNetwork/sao-did/parser"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	did := msg.Did
	newDocId := msg.NewDocId

	// check creator
	if !k.CheckCreator(ctx, msg.Creator, did) {
		logger.Error("invalid Creator", "creator", msg.Creator, "did", did)
		return nil, types.ErrInvalidCreator
	}

	if len(msg.RemoveAccountDid) == 0 {
		logger.Error("remove list should not be empty", "did", did)
		return nil, types.ErrNoNeedToUpdate
	}

	if len(msg.UpdateAccountAuth) == 0 {
		logger.Error("update list should include a payment account auth", "did", did)
		return nil, types.ErrUpdateAccAuthEmpty
	}

	payAddr, found := k.GetPaymentAddress(ctx, did)
	if !found {
		logger.Error("payment address not set", "did", did)
		return nil, types.ErrPayAddrNotSet
	}

	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		logger.Error("accountList is not found with did %v", did)
		return nil, types.ErrAccountListNotFound
	}

	// check auth count
	if len(accountList.AccountDids) != len(msg.RemoveAccountDid)+len(msg.UpdateAccountAuth) {
		logger.Error("account auth count dose not match",
			"len(exist)", len(accountList.AccountDids),
			"len(update)", len(msg.UpdateAccountAuth),
			"len(remove)", len(msg.RemoveAccountDid))
		return nil, types.ErrInvalidAuthCount
	}

	// ensure all accountDids are handled
	for _, accountDid := range accountList.AccountDids {
		if !inList(accountDid, msg.RemoveAccountDid) && !inUpdateList(accountDid, msg.UpdateAccountAuth) {
			logger.Error("accountDid %v is not handled, put AccountAuth with new sidDocument in update to keep it alive, put accountDid in remove to drop it", accountDid)
			return nil, types.ErrUnhandledAccountDid
		}
	}

	// check and unbind account
	for _, accDid := range msg.RemoveAccountDid {
		accountId, found := k.GetAccountId(ctx, accDid)
		if !found {
			logger.Error("accountId not found", "accountDid", accDid)
			return nil, types.ErrAccountIdNotFound
		}
		accIdSplits := strings.Split(accountId.AccountId, ":")
		if len(accIdSplits) == 3 &&
			accIdSplits[0] == "cosmos" &&
			accIdSplits[1] == ctx.ChainID() &&
			accIdSplits[2] == payAddr.Address {
			logger.Error("cannot unbind payment address", "did", did, "accountDid", accDid, "accountId", accountId.AccountId)
			return nil, types.ErrUnbindPayAddr
		}
		k.RemoveDidBindingProof(ctx, accountId.AccountId)
		k.RemoveAccountId(ctx, accDid)
	}

	// add new SidDocument
	parsedDid, err := parser.Parse(did)
	if err != nil {
		logger.Error("failed to parse did", "did", did)
		return nil, types.ErrInvalidDid
	}

	versions, found := k.GetSidDocumentVersion(ctx, parsedDid.ID)
	if inList(newDocId, versions.VersionList) {
		logger.Error("newDocId is exists in version list", "did", did, "newDocId", newDocId)
		return nil, types.ErrDocExists
	}

	_, found = k.GetSidDocument(ctx, newDocId)
	if found {
		logger.Error("docId exists", "doc_id", newDocId, "did", did)
		return nil, types.ErrDocExists
	}

	calDocId, err := CalculateDocId(msg.Keys, msg.Timestamp)
	if err != nil {
		logger.Error("failed to calculate doc Id", "did", did, "err", err)
		return nil, types.ErrInvalidKeys
	}

	// verify and set sid document if sid is new
	if newDocId != calDocId {
		logger.Error("inconsistent docId", "calculatedDocId", calDocId, "newDocId", newDocId, "did", did)
		return nil, types.ErrInconsistentDocId
	}

	versions.VersionList = append(versions.VersionList, newDocId)

	k.SetSidDocument(ctx, types.SidDocument{
		VersionId: newDocId,
		Keys:      msg.Keys,
	})

	k.SetSidDocumentVersion(ctx, versions)

	// update AccountAuth
	// update
	for _, accAuth := range msg.UpdateAccountAuth {
		k.SetAccountAuth(ctx, *accAuth)
	}
	// remove
	for _, toRemove := range msg.RemoveAccountDid {
		k.RemoveAccountAuth(ctx, toRemove)
		for i, ad := range accountList.AccountDids {
			if ad == toRemove {
				accountList.AccountDids = append(accountList.AccountDids[:i], accountList.AccountDids[i+1:]...)
				break
			}
		}
	}
	// payment address cannot be removed, so accountList cannot be empty
	k.SetAccountList(ctx, accountList)

	// update PastSeed
	ps, found := k.GetPastSeeds(ctx, did)
	if found {
		if inList(msg.PastSeed, ps.Seeds) {
			logger.Error("past seed exists", "did", did)
			return nil, types.ErrSeedExists
		}
		ps.Seeds = append(ps.Seeds, msg.PastSeed)
	} else {
		ps = types.PastSeeds{
			Did:   did,
			Seeds: []string{msg.PastSeed},
		}
	}
	k.SetPastSeeds(ctx, ps)

	return &types.MsgUpdateResponse{}, nil
}

package keeper

import (
	"context"
	"github.com/SaoNetwork/sao-did/parser"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

func (k msgServer) Update(goCtx context.Context, msg *types.MsgUpdate) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	did := msg.Did
	newDocId := msg.NewDocId
	removeList := msg.RemoveAccountDid
	updateList := msg.UpdateAccountAuth

	// check creator
	if !k.CheckCreator(ctx, msg.Creator, did) {
		logger.Error("invalid Creator", "creator", msg.Creator, "did", did)
		return nil, types.ErrInvalidCreator
	}

	now := time.Now().Unix()
	if msg.Timestamp+EXPIRE_DURATION < uint64(now) {
		logger.Error("timestamp is too old", "proof.Timestamp", msg.Timestamp, "now", now)
		return nil, types.ErrOutOfDate
	}

	if len(removeList) == 0 {
		logger.Error("remove list should not be empty", "did", did)
		return nil, types.ErrNoNeedToUpdate
	}

	if len(updateList) == 0 {
		logger.Error("update list should include a payment account auth", "did", did)
		return nil, types.ErrUpdateAccAuthEmpty
	}

	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		// unreachable
		logger.Error("accountList is not found with did %v", did)
		return nil, types.ErrAccountListNotFound
	}

	if len(accountList.AccountDids) != len(removeList)+len(updateList) {
		logger.Error("account auth count dose not match",
			"len(stored)", len(accountList.AccountDids),
			"len(update)", len(updateList),
			"len(remove)", len(removeList))
		return nil, types.ErrInvalidAuthCount
	}

	// ensure all accountDids are handled
	for _, accountDid := range accountList.AccountDids {
		if !inList(accountDid, removeList) && !inUpdateList(accountDid, updateList) {
			logger.Error("accountDid %v is not handled, put AccountAuth with new sidDocument in update to keep it alive, put accountDid in remove to drop it", accountDid)
			return nil, types.ErrUnhandledAccountDid
		}
	}

	ps, foundPastSeeds := k.GetPastSeeds(ctx, did)
	if foundPastSeeds && inList(msg.PastSeed, ps.Seeds) {
		logger.Error("past seed exists", "did", did)
		return nil, types.ErrSeedExists
	}

	payAddr, found := k.GetPaymentAddress(ctx, did)
	if !found {
		logger.Error("payment address not set", "did", did)
		return nil, types.ErrPayAddrNotSet
	}

	// check remove account
	removeAccId := make([]string, 0)
	for _, accDid := range removeList {
		accountId, found := k.GetAccountId(ctx, accDid)
		if !found {
			logger.Error("accountId not found", "accountDid", accDid)
			return nil, types.ErrAccountIdNotFound
		}
		caip10, err := parseAcccountId(accountId.AccountId)
		if err != nil {
			logger.Error("failed to parse accountId!!", "accountId", accountId.AccountId, "did", did, "err", err)
			return nil, types.ErrInvalidAccountId
		}
		if caip10.Network == "cosmos" &&
			caip10.Chain == ctx.ChainID() &&
			caip10.Address == payAddr.Address {
			logger.Error("cannot unbind payment address", "did", did, "accountDid", accDid, "accountId", accountId.AccountId)
			return nil, types.ErrUnbindPayAddr
		}
		removeAccId = append(removeAccId, accountId.AccountId)
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

	// update database

	for _, accId := range removeAccId {
		k.RemoveDid(ctx, accId)
	}

	for _, accDid := range removeList {
		k.RemoveAccountId(ctx, accDid)
	}

	versions.VersionList = append(versions.VersionList, newDocId)

	k.SetSidDocument(ctx, types.SidDocument{
		VersionId: newDocId,
		Keys:      msg.Keys,
	})

	k.SetSidDocumentVersion(ctx, versions)

	// update AccountAuth
	// update
	for _, accAuth := range updateList {
		k.SetAccountAuth(ctx, *accAuth)
	}
	// remove
	for _, toRemove := range removeList {
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
	if foundPastSeeds {
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

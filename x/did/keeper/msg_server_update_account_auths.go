package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAccountAuths(goCtx context.Context, msg *types.MsgUpdateAccountAuths) (*types.MsgUpdateAccountAuthsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	did := msg.Did
	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		logger.Error("accountList is not found with did %v", did)
		return nil, types.ErrAccountListNotFound
	}

	// check all past account dids are handled
	for _, accountDid := range accountList.AccountDids {
		if !inList(accountDid, msg.Remove) && !inUpdateList(accountDid, msg.Update) {
			logger.Error("accountDid %v is not handled, put AccountAuth with new sidDocument in update to keep it alive, put accountDid in remove to drop it", accountDid)
			return nil, types.ErrUnhandledAccountDid
		}
	}

	// add
	for _, accAuth := range msg.Update {
		// continue if account is already in account list
		if inList(accAuth.AccountDid, accountList.AccountDids) {
			continue
		}

		_, found = k.GetAccountAuth(ctx, accAuth.AccountDid)
		// store account and add it to list if account not exists
		if !found {
			k.SetAccountAuth(ctx, *accAuth)
			accountList.AccountDids = append(accountList.AccountDids, accAuth.AccountDid)
		}
	}

	// remove
	for _, toRemove := range msg.Remove {
		k.RemoveAccountAuth(ctx, toRemove)
		for i, ad := range accountList.AccountDids {
			if ad == toRemove {
				accountList.AccountDids = append(accountList.AccountDids[:i], accountList.AccountDids[i+1:]...)
				break
			}
		}
	}

	if len(accountList.AccountDids) == 0 {
		k.RemoveAccountList(ctx, did)
	}
	if len(msg.Remove) != 0 || len(msg.Update) != 0 {
		k.SetAccountList(ctx, accountList)
	}

	return &types.MsgUpdateAccountAuthsResponse{}, nil
}

func inUpdateList(did string, list []*types.AccountAuth) bool {
	for _, v := range list {
		if v.AccountDid == did {
			return true
		}
	}
	return false
}

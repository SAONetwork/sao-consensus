package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAccountAuths(goCtx context.Context, msg *types.MsgUpdateAccountAuths) (*types.MsgUpdateAccountAuthsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	did := msg.Did
	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		return nil, types.ErrAccountListNotFound
	}

outer:
	// add
	for _, accAuth := range msg.Update {
		// continue if account is already in account list
		for _, accDid := range accountList.AccountDids {
			if accDid == accAuth.AccountDid {
				continue outer
			}
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

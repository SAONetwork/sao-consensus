package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAccountAuths(goCtx context.Context, msg *types.MsgUpdateAccountAuths) (*types.MsgUpdateAccountAuthsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	did := msg.Did
	// add
	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		return nil, types.ErrAccountListNotFound
	}

outer:
	for _, aa := range msg.Update {

		for _, ad := range accountList.AccountDids {
			if ad == aa.AccountDid {
				continue outer
			}
		}
		accountList.AccountDids = append(accountList.AccountDids, aa.AccountDid)

		_, found = k.GetAccountAuth(ctx, aa.AccountDid)
		if !found {
			k.SetAccountAuth(ctx, *aa)
		}
		k.SetAccountList(ctx, accountList)
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

	return &types.MsgUpdateAccountAuthsResponse{}, nil
}

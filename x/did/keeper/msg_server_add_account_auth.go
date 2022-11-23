package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAccountAuth(goCtx context.Context, msg *types.MsgAddAccountAuth) (*types.MsgAddAccountAuthResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	did := msg.Did
	aa := *msg.AccountAuth
	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		accountList = types.AccountList{
			Did:         did,
			AccountDids: []string{aa.AccountDid},
		}
	} else {
		for _, ad := range accountList.AccountDids {
			if ad == aa.AccountDid {
				return nil, types.ErrAuthExists
			}
		}
		accountList.AccountDids = append(accountList.AccountDids, aa.AccountDid)
	}

	_, found = k.GetAccountAuth(ctx, aa.AccountDid)
	if found {
		return nil, types.ErrAuthExists
	} else {
		k.SetAccountAuth(ctx, aa)
	}
	k.SetAccountList(ctx, accountList)

	return &types.MsgAddAccountAuthResponse{}, nil
}

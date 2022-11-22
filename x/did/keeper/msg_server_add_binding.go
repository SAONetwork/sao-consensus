package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBinding(goCtx context.Context, msg *types.MsgAddBinding) (*types.MsgAddBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	_, exist := k.GetDidBindingProofs(ctx, msg.GetAccountId())
	if exist {
		return &types.MsgAddBindingResponse{}, nil
	}
	newDidBindingProofs := types.DidBindingProofs{
		AccountId: msg.GetAccountId(),
		Proof:     msg.GetProof(),
	}
	k.SetDidBindingProofs(ctx, newDidBindingProofs)

	return &types.MsgAddBindingResponse{}, nil
}

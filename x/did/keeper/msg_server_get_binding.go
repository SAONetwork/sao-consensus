package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) GetBinding(goCtx context.Context, msg *types.MsgGetBinding) (*types.MsgGetBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	didBindingProofs, exist := k.GetDidBindingProofs(ctx, msg.GetAccountId())
	if !exist {
		return nil, sdkerrors.Wrapf(types.ErrProofNotFound, "No Binding Proofs of %s", msg.AccountId)
	}
	binding := types.Binding{
		AccountId: msg.GetAccountId(),
		Proof:     didBindingProofs.Proof,
	}

	return &types.MsgGetBindingResponse{&binding}, nil
}

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

	return &types.MsgAddBindingResponse{}, nil
}

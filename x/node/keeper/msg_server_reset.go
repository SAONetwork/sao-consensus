package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Reset(goCtx context.Context, msg *types.MsgReset) (*types.MsgResetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgResetResponse{}, nil
}

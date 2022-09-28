package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Logout(goCtx context.Context, msg *types.MsgLogout) (*types.MsgLogoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgLogoutResponse{}, nil
}

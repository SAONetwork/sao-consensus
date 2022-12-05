package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdataPermission(goCtx context.Context, msg *types.MsgUpdataPermission) (*types.MsgUpdataPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdataPermissionResponse{}, nil
}

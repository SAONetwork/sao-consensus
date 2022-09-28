package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Logout(goCtx context.Context, msg *types.MsgLogout) (*types.MsgLogoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.GetNode(ctx, msg.Creator)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	k.RemoveNode(ctx, msg.Creator)

	return &types.MsgLogoutResponse{}, nil
}

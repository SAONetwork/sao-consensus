package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Login(goCtx context.Context, msg *types.MsgLogin) (*types.MsgLoginResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_, found := k.GetNode(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyRegistered, "%s", msg.Creator)
	}

	var node = types.Node{
		Peer:            "",
		Creator:         msg.Creator,
		Reputation:      10000.0,
		Status:          types.NODE_STATUS_NA,
		LastAliveHeight: ctx.BlockHeight(),
	}

	k.SetNode(ctx, node)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.LoginEventType,
			sdk.NewAttribute(types.NodeEventCreator, node.Creator),
			sdk.NewAttribute(types.NodeEventPeer, node.Peer),
		),
	)

	return &types.MsgLoginResponse{}, nil
}

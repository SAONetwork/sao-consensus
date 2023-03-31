package keeper

import (
	"context"
	"strings"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/multiformats/go-multiaddr"
)

func (k msgServer) Reset(goCtx context.Context, msg *types.MsgReset) (*types.MsgResetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.GetNode(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNodeNotFound, "%s", msg.Creator)
	}

	if node.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrOnlyOwner, "only node owner can execute this action")
	}

	if msg.Status != types.NODE_STATUS_NA && node.Status != msg.Status {
		node.Status = msg.Status
	}

	if msg.Peer != "" && node.Peer != msg.Peer {
		for _, peerInfo := range strings.Split(msg.Peer, ",") {

			_, err := multiaddr.NewMultiaddr(peerInfo)

			if err != nil {
				return nil, sdkerrors.Wrapf(types.ErrInvalidPeer, "%s", peerInfo)
			}
		}
		node.Peer = msg.Peer
	}

	if msg.Description != nil && node.Description != msg.Description {
		node.Description = msg.Description
	}

	node.LastAliveHeight = ctx.BlockHeight()

	k.SetNode(ctx, node)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ResetEventType,
			sdk.NewAttribute(types.NodeEventCreator, node.Creator),
			sdk.NewAttribute(types.NodeEventPeer, node.Peer),
		),
	)

	return &types.MsgResetResponse{}, nil
}

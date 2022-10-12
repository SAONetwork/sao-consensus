package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/libp2p/go-libp2p/core/peer"
)

func (k msgServer) Reset(goCtx context.Context, msg *types.MsgReset) (*types.MsgResetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.GetNode(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNodeNotFound, "%s", msg.Creator)
	}

	_, err := peer.AddrInfoFromString(msg.Peer)

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPeer, "%s", msg.Peer)
	}

	signers := msg.GetSigners()

	if len(signers) != 1 || signers[0].String() != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrSignerAndCreator, "signer shoud equal to creator")
	}

	if len(signers) != 1 || signers[0].String() != node.Creator {
		return nil, sdkerrors.Wrapf(types.ErrOnlyOwner, "only node owner can execute this action")
	}

	node.Peer = msg.Peer

	k.SetNode(ctx, node)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ResetEventType,
			sdk.NewAttribute(types.NodeEventCreator, node.Creator),
			sdk.NewAttribute(types.NodeEventPeer, node.Peer),
		),
	)

	return &types.MsgResetResponse{}, nil
}

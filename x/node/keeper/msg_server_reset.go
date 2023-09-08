package keeper

import (
	"context"
	"strings"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/multiformats/go-multiaddr"
)

const StorageThreshold = 10 << 30

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

	if msg.Validator != "" && node.Validator != msg.Validator {
		valAddr, err := sdk.ValAddressFromBech32(msg.Validator)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidValidator, "%s", msg.Validator)
		}
		_, found := k.staking.GetValidator(ctx, valAddr)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrValidatorNotFound, "%s", msg.Validator)
		}
		node.Validator = msg.Validator
	}

	if msg.Description != nil && node.Description != msg.Description {
		node.Description = msg.Description
	}

	if len(msg.TxAddresses) != 0 {
		node.TxAddresses = msg.TxAddresses
	}

	node.LastAliveHeight = ctx.BlockHeight()

	// super check
	node.Role = types.NODE_NORMAL
	if msg.Status&types.NODE_STATUS_SUPER_REQUIREMENT == types.NODE_STATUS_SUPER_REQUIREMENT {
		pledge, found := k.GetPledge(ctx, msg.Creator)
		if found && pledge.TotalStorage >= k.VstorageThreshold(ctx) {
			k.CheckNodeShare(ctx, &node, msg.Creator)
		}
	}

	k.SetNode(ctx, node)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ResetEventType,
			sdk.NewAttribute(types.NodeEventCreator, node.Creator),
			sdk.NewAttribute(types.NodeEventPeer, node.Peer),
		),
	)

	return &types.MsgResetResponse{}, nil
}

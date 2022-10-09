package keeper

import (
	"context"
	"fmt"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check provider
	node, found := k.node.GetNode(ctx, msg.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	// check cid
	_, err := cid.Decode(msg.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invali cid: %s", msg.Cid)
	}

	var order = types.Order{
		Creator:  msg.Creator,
		Provider: node.Creator,
		Cid:      msg.Cid,
		Expire:   int32(ctx.BlockHeight()) + 86400,
		Duration: msg.Duration,
		Status:   types.OrderPending,
		Replica:  msg.Replica,
	}

	orderId := k.AppendOrder(ctx, order)

	// choose node
	sps := k.node.RandomSP(ctx, int(msg.Replica))

	// check replica
	if msg.Replica <= 0 || int(msg.Replica) > len(sps) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
	}

	shards := make(map[string]*types.Shard, 0)
	for _, sp := range sps {
		shards[sp.Creator] = &types.Shard{
			OrderId: orderId,
			Status:  types.ShardWaiting,
			Cid:     msg.Cid,
		}
	}

	order.Shards = shards

	k.SetOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	for provider, shard := range order.Shards {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.NewShardEventType,
				sdk.NewAttribute(types.ShardEventProvider, provider),
				sdk.NewAttribute(types.EventCid, shard.Cid),
			),
		)
	}

	return &types.MsgStoreResponse{OrderId: orderId}, nil
}

package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) newRandomShard(goCtx context.Context, order *types.Order) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// choose node
	sps := k.node.RandomSP(ctx, int(order.Replica))

	logger := k.Logger(ctx)

	logger.Info("random sps", "sps", sps)

	// check replica
	if order.Replica <= 0 || int(order.Replica) > len(sps) {
		return sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
	}

	shards := make(map[string]*types.Shard, 0)
	for _, sp := range sps {
		shards[sp.Creator] = &types.Shard{
			OrderId: order.Id,
			Status:  types.ShardWaiting,
			Cid:     order.Cid,
		}
	}

	order.Shards = shards

	order.Status = types.OrderDataReady

	for provider, shard := range order.Shards {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.NewShardEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute(types.OrderEventProvider, order.Provider),
				sdk.NewAttribute(types.ShardEventProvider, provider),
				sdk.NewAttribute(types.EventCid, shard.Cid),
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			),
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.OrderDataReadyEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	return nil
}

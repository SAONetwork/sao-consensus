package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewShardTask(ctx sdk.Context, order *types.Order, provider string) *types.Shard {

	shard := types.Shard{
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     order.Cid,
		Size_:   order.Size_,
		Sp:      provider,
	}

	shard.Id = k.AppendShard(ctx, shard)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.ShardEventProvider, provider),
			sdk.NewAttribute(types.EventCid, shard.Cid),
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventOperation, fmt.Sprintf("%d", order.Operation)),
		),
	)
	return &shard
}

func (k Keeper) FulfillShard(ctx sdk.Context, shard *types.Shard, sp string, cid string) {

	shard.Status = types.ShardCompleted
	shard.Cid = cid

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", shard.OrderId)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
		),
	)
}

func (k Keeper) TerminateShard(ctx sdk.Context, shard *types.Shard, sp string, owner string, orderId uint64) error {

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TerminateShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
			sdk.NewAttribute(types.EventCid, shard.Cid),
		),
	)
	return nil
}

func (k Keeper) GetOrderShardBySP(ctx sdk.Context, order *types.Order, sp string) *types.Shard {

	for _, id := range order.Shards {
		shard, found := k.GetShard(ctx, id)
		if found && shard.Sp == sp {
			return &shard
		}
	}
	return nil
}

func (k Keeper) RenewShard(ctx sdk.Context, order *types.Order, sp string) error {

	shard := k.GetOrderShardBySP(ctx, order, sp)
	if shard == nil {
		return status.Errorf(codes.NotFound, "shard of %s not found", sp)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
		),
	)

	k.SetOrder(ctx, *order)
	return nil
}

func (k Keeper) MigrateShard(ctx sdk.Context, oldShard *types.Shard, order *types.Order, from string, to string) *types.Shard {

	shard := types.Shard{
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     oldShard.Cid,
		From:    from,
		Size_:   oldShard.Size_,
		Sp:      to,
	}

	shard.Id = k.AppendShard(ctx, shard)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", oldShard.OrderId)),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.ShardEventProvider, from),
			sdk.NewAttribute(types.EventCid, shard.Cid),
			sdk.NewAttribute(types.OrderEventOperation, fmt.Sprintf("%d", order.Operation)),
		),
	)

	return &shard
}

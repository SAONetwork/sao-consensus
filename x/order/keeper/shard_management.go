package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NewShardTask(ctx sdk.Context, order *types.Order, provider string) *types.Shard {

	shard := &types.Shard{
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     order.Cid,
		//TODO: use the same type as Order.Size_
		Size_: order.Size_,
	}

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
	return shard
}

func (k Keeper) FulfillShard(ctx sdk.Context, order *types.Order, sp string, cid string, _ uint64) error {

	shard := order.Shards[sp]

	shard.Status = types.ShardCompleted
	shard.Cid = cid

	order.Shards[sp] = shard

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
		),
	)

	k.SetOrder(ctx, *order)
	return nil
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

func (k Keeper) RenewShard(ctx sdk.Context, order *types.Order, sp string) error {

	shard := order.Shards[sp]

	order.Shards[sp] = shard

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
		),
	)

	k.SetOrder(ctx, *order)
	return nil
}

func (k Keeper) MigrateShard(ctx sdk.Context, order *types.Order, from string, to string) *types.Shard {

	shard := &types.Shard{
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     order.Cid,
		From:    from,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.ShardEventProvider, from),
			sdk.NewAttribute(types.EventCid, shard.Cid),
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventOperation, fmt.Sprintf("%d", order.Operation)),
		),
	)

	return shard
}

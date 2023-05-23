package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleExpiredShard(ctx sdk.Context, shardId uint64) {
	shard, found := k.order.GetShard(ctx, shardId)
	if !found {
		return
	}

	order, found := k.order.GetOrder(ctx, shard.OrderId)
	if !found {
		return
	}

	k.node.ShardRelease(ctx, sdk.MustAccAddressFromBech32(shard.Sp), &shard)
	k.market.WorkerRelease(ctx, &order, &shard)
	k.order.RemoveShard(ctx, shardId)
	if len(order.Shards) == 1 {
		if order.Shards[0] == shardId {
			k.order.RemoveOrder(ctx, order.Id)
		}
	} else {
		for i, id := range order.Shards {
			if id == shardId {
				order.Shards = append(order.Shards[:i], order.Shards[i+1:]...)
				break
			}
		}
		k.order.SetOrder(ctx, order)
	}
}

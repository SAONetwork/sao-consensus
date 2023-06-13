package keeper

import (
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleTimeoutOrder(ctx sdk.Context, orderId uint64) {
	log := k.Logger(ctx)
	order, found := k.order.GetOrder(ctx, orderId)
	if !found {
		return
	}

	if uint64(ctx.BlockHeight())+order.Timeout >= order.CreatedAt+order.Duration {
		return
	}

	if order.Status == ordertypes.OrderCompleted || order.Status == ordertypes.OrderMigrating {
		return
	}

	shards := k.FindShardsByOrderId(ctx, order.Id)
	var sps []string
	for _, shard := range shards {
		sps = append(sps, shard.Sp)
	}
	log.Debug("order timeout", "orderId", order.Id, "sps", sps)

	var newTimeoutBlock uint64

	var updateShards []uint64
	for _, shard := range shards {
		if shard.Status == ordertypes.ShardCompleted {
			updateShards = append(updateShards, shard.Id)
			continue
		}
		if shard.Status == ordertypes.ShardWaiting {
			// TODO: sp punishment
			//k.node.DecreaseReputation(ctx, shard.Sp, types.TimeoutReputationPunishment)
		}

		randSp := k.node.RandomSP(ctx, 1, sps, int64(shard.Size_))
		log.Debug("fix shard", "shardId", shard.Id, "oldSP", shard.Sp, "newSp", randSp)
		if len(randSp) != 0 {
			// remove old shard
			k.order.RemoveShard(ctx, shard.Id)

			// create new shard
			newShard := k.order.NewShardTask(ctx, &order, randSp[0].Creator)
			updateShards = append(updateShards, newShard.Id)
			sps = append(sps, randSp[0].Creator)
		} else {
			updateShards = append(updateShards, shard.Id)
		}
		newTimeoutBlock = uint64(ctx.BlockHeight()) + order.Timeout
	}

	if newTimeoutBlock > uint64(ctx.BlockHeight()) {
		order.Shards = updateShards
		k.order.SetOrder(ctx, order)
		k.SetTimeoutOrderBlock(ctx, order, newTimeoutBlock)
	}

	return
}

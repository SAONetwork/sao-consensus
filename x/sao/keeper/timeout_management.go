package keeper

import (
	"fmt"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) HandleTimeoutOrder(ctx sdk.Context, orderId uint64) {
	log := k.Logger(ctx)
	order, found := k.order.GetOrder(ctx, orderId)
	if !found {
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

	for _, shard := range shards {
		if shard.Status == ordertypes.ShardCompleted {
			continue
		}
		if shard.Status == ordertypes.ShardWaiting {
			k.node.DecreaseReputation(ctx, shard.Sp, types.TimeoutReputationPunishment)
		}

		randSp := k.node.RandomSP(ctx, 1, sps)
		log.Debug("fix shard", "shardId", shard.Id, "oldSP", shard.Sp, "newSp", randSp)
		if len(randSp) != 0 {
			// remove old shard
			k.order.RemoveShard(ctx, shard.Id)
			for i, id := range order.Shards {
				if id == shard.Id {
					order.Shards = append(order.Shards[:i], order.Shards[i+1:]...)
					break
				}
			}

			// create new shard
			newShard := k.order.NewShardTask(ctx, &order, randSp[0].Creator)
			order.Shards = append(order.Shards, newShard.Id)
			sps = append(sps, randSp[0].Creator)
		}
		newTimeoutBlock = uint64(ctx.BlockHeight()) + order.Timeout
	}

	if newTimeoutBlock >= order.CreatedAt+order.Duration {
		return
	} else if newTimeoutBlock > uint64(ctx.BlockHeight()) {
		k.order.SetOrder(ctx, order)
		k.SetTimeoutOrderBlock(ctx, order, newTimeoutBlock)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(ordertypes.OrderTimeoutEventType,
			sdk.NewAttribute(ordertypes.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	return
}

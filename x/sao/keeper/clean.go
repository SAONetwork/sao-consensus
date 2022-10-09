package keeper

import (
	"github.com/SaoNetwork/sao/x/sao/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CheckOrder(ctx sdk.Context) {
	orders := k.GetAllOrder(ctx)
	for _, order := range orders {
		switch order.Status {
		case types.OrderCompleted:
		case types.OrderTerminated:
		case types.OrderCanceled:
			continue
		default:
			// case order expired
			if order.Expire < int32(ctx.BlockHeight()) {
				// check shard
				for sp, shard := range order.Shards {
					if shard.Status == types.ShardWaiting {
						k.node.DecreaseReputation(ctx, sp, 1000)
					}
				}
				order.Status = types.OrderExpired
				k.SetOrder(ctx, order)
			}
		}
	}
}

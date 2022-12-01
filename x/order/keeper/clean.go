package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"

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
						fmt.Println(sp)
					}
				}
				order.Status = types.OrderExpired
				k.SetOrder(ctx, order)
			}
		}
	}
}

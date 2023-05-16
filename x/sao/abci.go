package sao

import (
	"fmt"
	"github.com/SaoNetwork/sao/x/sao/keeper"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	TimeoutOrder, found := k.GetTimeoutOrder(ctx, uint64(ctx.BlockHeight()))
	if found {

		for _, orderId := range TimeoutOrder.OrderList {
			k.HandleTimeoutOrder(ctx, orderId)
		}

		k.RemoveTimeoutOrder(ctx, TimeoutOrder.Height)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderTimeoutEventType,
				sdk.NewAttribute(types.EventTimeoutOrderList, fmt.Sprintf("%v", TimeoutOrder.OrderList)),
			),
		)

	}
	ExpiredShard, found := k.GetExpiredShard(ctx, uint64(ctx.BlockHeight()))
	if found {

		for _, shardId := range ExpiredShard.ShardList {
			k.HandleExpiredShard(ctx, shardId)
		}

		k.RemoveExpiredShard(ctx, ExpiredShard.Height)
	}
}

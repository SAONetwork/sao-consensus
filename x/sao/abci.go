package sao

import (
	"github.com/SaoNetwork/sao/x/sao/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	TimeoutOrder, found := k.GetTimeoutOrder(ctx, uint64(ctx.BlockHeight()))
	if found {

		for _, orderId := range TimeoutOrder.OrderList {
			k.HandleTimeoutOrder(ctx, orderId)
		}

		k.RemoveTimeoutOrder(ctx, TimeoutOrder.Height)
	}
}

package order

import (
	"github.com/SaoNetwork/sao/x/order/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// clean expired order
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	expiredOrder, found := k.GetExpiredOrder(ctx, uint64(ctx.BlockHeight()))
	if !found {
		return
	}

	for _, orderId := range expiredOrder.Data {
		k.RefundExpiredOrder(ctx, orderId)
	}

	k.RemoveExpiredOrder(ctx, expiredOrder.Height)
}

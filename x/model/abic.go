package model

import (
	"github.com/SaoNetwork/sao/x/model/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// clean expired data
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	expiredData, foundExpired := k.GetExpiredData(ctx, uint64(ctx.BlockHeight()))
	if foundExpired {

		for _, dataId := range expiredData.Data {
			k.DeleteMeta(ctx, dataId)
		}

		k.RemoveExpiredData(ctx, expiredData.Height)
	}

	expiredOrder, found := k.GetExpiredOrder(ctx, uint64(ctx.BlockHeight()))
	if !found {
		return
	}

	for _, orderId := range expiredOrder.Data {
		k.RefundExpiredOrder(ctx, orderId)
	}

	k.RemoveExpiredOrder(ctx, expiredOrder.Height)

	orderFinish, foundFinished := k.GetOrderFinish(ctx, uint64(ctx.BlockHeight()))
	if foundFinished {

		for _, orderId := range orderFinish.Data {
			k.OrderSettlement(ctx, orderId)
		}

		k.RemoveOrderFinish(ctx, orderFinish.Height)
	}

}

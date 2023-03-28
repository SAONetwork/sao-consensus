package model

import (
	"github.com/SaoNetwork/sao/x/model/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// clean expired data
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	expiredData, foundExpired := k.GetExpiredData(ctx, uint64(ctx.BlockHeight()))
	if !foundExpired {
		return
	}

	for _, dataId := range expiredData.Data {
		k.DeleteMeta(ctx, dataId)
	}

	k.RemoveExpiredData(ctx, expiredData.Height)
}

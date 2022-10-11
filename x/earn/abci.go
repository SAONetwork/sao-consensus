package earn

import (
	"github.com/SaoNetwork/sao/x/earn/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	pool, found := k.GetPool(ctx)
	if !found {
		return
	}

	if pool.Denom.Amount.IsZero() {
		pool.LastRewardBlock = ctx.BlockHeight()
		return
	}

	params := k.GetParams(ctx)

	if params.BlockReward < 0 {
		return
	}
	rewardCoin := sdk.NewCoin(params.GetRewardDenom(), sdk.NewInt(int64(params.GetBlockReward())))

	rewardCoins := sdk.NewCoins(rewardCoin)

	err := k.MintCoins(ctx, rewardCoins)

	if err == nil {
		// update reward per share
		pool.CoinPerShare = pool.CoinPerShare + uint64(params.BlockReward)*1e12/uint64(pool.Denom.Amount.Int64())
	}

	pool.LastRewardBlock = ctx.BlockHeight()

}

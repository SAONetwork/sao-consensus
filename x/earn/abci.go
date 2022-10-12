package earn

import (
	"time"

	"github.com/SaoNetwork/sao/x/earn/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	logger := k.Logger(ctx)

	pool, found := k.GetPool(ctx)
	if !found {
		return
	}

	params := k.GetParams(ctx)

	if params.BlockReward <= 0 {
		return
	}

	rewardCoin := sdk.NewInt64Coin(params.EarnDenom, int64(params.BlockReward))

	rewardCoins := sdk.NewCoins(rewardCoin)

	logger.Info("mint new storage reward", "reward", rewardCoin.String())

	err := k.MintCoins(ctx, rewardCoins)

	pool.TotalReward.Add(rewardCoin)

	logger.Info("mint new storage reward ", "reward", rewardCoin.String())

	if pool.Denom.Amount.IsZero() {
		pool.LastRewardBlock = ctx.BlockHeight()
		k.SetPool(ctx, pool)
		return
	}

	if err == nil {
		// update reward per share
		pool.CoinPerShare = pool.CoinPerShare + uint64(rewardCoin.Amount.Int64())*1e12/uint64(pool.Denom.Amount.Int64())
	}

	pool.LastRewardBlock = ctx.BlockHeight()
	k.SetPool(ctx, pool)
}

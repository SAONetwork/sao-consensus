package node

import (
	"math/big"
	"time"

	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	logger := k.Logger(ctx)

	pool, found := k.GetPool(ctx)
	if !found {
		logger.Error("pool not found")
		return
	}

	if pool.Denom.Amount.Int64() <= 0 {
		logger.Error("not pledged ye")
		return
	}

	params := k.GetParams(ctx)

	if params.BlockReward <= 0 {
		logger.Error("invalid block reward")
		return
	}

	rewardCoin := sdk.NewInt64Coin(params.EarnDenom, int64(ctx.BlockHeight()-pool.LastRewardBlock)*int64(params.BlockReward))

	rewardCoins := sdk.NewCoins(rewardCoin)

	logger.Debug("mint node incentive coins", "coin", rewardCoin)

	k.Stats(ctx)

	err := k.MintCoins(ctx, rewardCoins)

	if pool.Denom.Amount.IsZero() {
		pool.LastRewardBlock = ctx.BlockHeight()
		k.SetPool(ctx, pool)
		return
	}

	if err == nil {
		pool.TotalReward.Add(rewardCoin)
		// update reward per share
		reward := new(big.Int).Mul(rewardCoin.Amount.BigInt(), big.NewInt(1e12))
		coinPerShare, _ := new(big.Int).SetString(pool.CoinPerShare, 10)
		coinPerShare = new(big.Int).Add(coinPerShare, new(big.Int).Div(reward, pool.Denom.Amount.BigInt()))
		pool.CoinPerShare = coinPerShare.String()
		logger.Debug("pool reward coin per share", "share", pool.CoinPerShare)
	}

	pool.LastRewardBlock = ctx.BlockHeight()
	k.SetPool(ctx, pool)
}

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

	err := k.MintCoins(ctx, rewardCoins)

	if pool.Denom.Amount.IsZero() {
		pool.LastRewardBlock = ctx.BlockHeight()
		k.SetPool(ctx, pool)
		return
	}

	if err == nil {
		pool.TotalReward.Add(rewardCoin)
		// update reward per share
		coinPerShare, _ := new(big.Int).SetString(pool.CoinPerShare, 10)
		reward := new(big.Int).Mul(rewardCoin.Amount.BigInt(), big.NewInt(1e12))
		coinPerShare = new(big.Int).Add(coinPerShare, new(big.Int).Div(reward, pool.Denom.Amount.BigInt()))
		pool.CoinPerShare = coinPerShare.String()
	}

	pool.LastRewardBlock = ctx.BlockHeight()
	k.SetPool(ctx, pool)
}

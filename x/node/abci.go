package node

import (
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

	if pool.TotalPledged.IsZero() {
		logger.Error("not pledged ye")
		return
	}

	params := k.GetParams(ctx)

	if params.BlockReward.IsZero() {
		logger.Error("invalid block reward")
		return
	}

	rewardCoin := sdk.NewCoin(params.BlockReward.Denom, params.BlockReward.Amount.MulRaw(int64(ctx.BlockHeight()-pool.LastRewardBlock)))

	rewardCoins := sdk.NewCoins(rewardCoin)

	logger.Debug("mint node incentive coins", "coin", rewardCoin)

	err := k.MintCoins(ctx, rewardCoins)

	if err == nil {
		pool.TotalReward = pool.TotalReward.Add(rewardCoin)
		pool.AccRewardPerByte.Amount = pool.AccRewardPerByte.Amount.Add(sdk.NewDecFromInt(rewardCoin.Amount).QuoInt64(pool.TotalStorage))
		pool.AccPledgePerByte.Amount = pool.AccRewardPerByte.Amount
	}

	if pool.TotalPledged.IsZero() {
		pool.LastRewardBlock = ctx.BlockHeight()
		k.SetPool(ctx, pool)
		return
	}

	pool.LastRewardBlock = ctx.BlockHeight()
	k.SetPool(ctx, pool)
}

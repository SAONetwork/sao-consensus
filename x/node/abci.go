package node

import (
	"math"
	"math/big"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TOTAL_REWARD = "400000000000000sao"

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {

	// 1st age 6.25 * 32000000 blocks
	// 16000000 blocks per year

	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	logger := k.Logger(ctx)

	pool, found := k.GetPool(ctx)
	if !found {
		logger.Error("pool not found")
		return
	}

	if pool.TotalPledged.IsZero() {
		return
	}

	params := k.GetParams(ctx)

	if params.BlockReward.IsZero() {
		logger.Error("invalid block reward")
		return
	}

	pool.TotalStorage += pool.PendingStorage
	pool.PendingStorage = 0

	var rewardCoin sdk.Coin

	subsidy := params.BlockReward.Amount.BigInt()
	subsidy.Rsh(subsidy, GetRewardAge(pool))
	rewardCoin = sdk.NewCoin(params.BlockReward.Denom, sdkmath.NewIntFromBigInt(subsidy))
	logger.Debug("reward age", "age", GetRewardAge(pool))

	if pool.TotalPledged.IsLT(params.Baseline) {
		apy, err := sdk.NewDecFromStr(params.AnnualPercentageYield)
		if err != nil {
			logger.Error("error node apy params", "err", err.Error())
			return
		}

		reward := sdk.NewDecCoinFromCoin(pool.TotalPledged).Amount.Mul(apy).QuoInt64(16000000).TruncateInt()
		logger.Debug("baseline mint", "reward", reward)
		if reward.LT(rewardCoin.Amount) {
			rewardCoin = sdk.NewCoin(params.BlockReward.Denom, reward)
		}
	}

	if rewardCoin.IsZero() {
		logger.Debug("waiting for more storage pledge", "current", pool.TotalPledged)
		return
	}

	if pool.NextRewardPerBlock.IsZero() {
		pool.NextRewardPerBlock = sdk.NewDecCoinFromCoin(rewardCoin)
	}

	// reset reward accumulation every 2000 blocks

	if ctx.BlockHeight()%100 == 0 {
		pool.RewardPerBlock = pool.NextRewardPerBlock
		pool.NextRewardPerBlock = sdk.NewDecCoinFromCoin(rewardCoin)
	}

	pool.NextRewardPerBlock.Amount = pool.NextRewardPerBlock.Amount.Add(sdk.NewDecFromInt(rewardCoin.Amount)).Quo(sdk.NewDec(2))

	rewardCoins := sdk.NewCoins(rewardCoin)

	logger.Debug("mint node incentive coins", "coin", rewardCoin)
	err := k.MintCoins(ctx, rewardCoins)
	if err == nil {
		pool.TotalReward = pool.TotalReward.Add(rewardCoin)
		pool.AccRewardPerByte.Amount = pool.AccRewardPerByte.Amount.Add(sdk.NewDecFromInt(rewardCoin.Amount).QuoInt64(pool.TotalStorage))
		pool.AccPledgePerByte.Amount = pool.AccRewardPerByte.Amount
	}

	pool.RewardedBlockCount += 1
	k.SetPool(ctx, pool)
}

func GetRewardAge(pool types.Pool) uint {
	totalReward, _ := sdk.ParseCoinNormalized(TOTAL_REWARD)
	remain := totalReward.Sub(pool.TotalReward)
	t, _ := new(big.Float).SetInt(totalReward.Amount.Quo(remain.Amount).BigInt()).Float64()
	return uint(math.Log2(t))
}

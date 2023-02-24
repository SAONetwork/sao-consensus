package node

import (
	"math/big"
	"time"

	sdkmath "cosmossdk.io/math"
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
		return
	}

	params := k.GetParams(ctx)
	if params.BlockReward.IsZero() {
		logger.Error("invalid block reward")
		return
	}

	bitLen := uint(params.BlockReward.Amount.BigInt().BitLen())
	halvings := uint(pool.RewardedBlockCount / types.NODE_SUBSIDY_HALVING_INTERVAL)
	if halvings >= bitLen {
		return
	}

	subsidy := params.BlockReward.Amount.BigInt()
	subsidy.Rsh(subsidy, halvings)
	if subsidy.Cmp(big.NewInt(0)) <= 0 {
		return
	}

	rewardCoin := sdk.NewCoin(params.BlockReward.Denom, sdkmath.NewIntFromBigInt(subsidy))
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

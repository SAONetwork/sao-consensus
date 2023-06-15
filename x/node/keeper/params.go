package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	apy, _ := sdk.NewDecFromStr(k.AnnualPercentageYield(ctx))
	return types.NewParams(
		k.BlockReward(ctx),
		k.Baseline(ctx),
		apy,
		k.HalvingPeriod(ctx),
		k.AdjustmentPeriod(ctx),
		k.FishmenInfo(ctx),
		k.PenaltyBase(ctx),
		k.MaxPenalty(ctx),
		k.ShareThreshold(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// BlockReward returns the BlockReward param
func (k Keeper) BlockReward(ctx sdk.Context) (res sdk.Coin) {
	k.paramstore.Get(ctx, types.KeyBlockReward, &res)
	return
}

func (k Keeper) Baseline(ctx sdk.Context) (res sdk.Coin) {
	k.paramstore.Get(ctx, types.KeyBaseLine, &res)
	return
}

func (k Keeper) AnnualPercentageYield(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyAPY, &res)
	return
}

func (k Keeper) HalvingPeriod(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyHalvingPeriod, &res)
	return
}

func (k Keeper) ShareThreshold(ctx sdk.Context) (res sdk.Dec) {
	var val string
	k.paramstore.Get(ctx, types.KeyShareThreshold, &val)
	res, _ = sdk.NewDecFromStr(val)
	return
}

func (k Keeper) AdjustmentPeriod(ctx sdk.Context) (res int64) {
	k.paramstore.Get(ctx, types.KeyAdjustmentPeriod, &res)
	return
}

func (k Keeper) FishmenInfo(ctx sdk.Context) (fishmenInfo string) {
	k.paramstore.Get(ctx, types.KeyFishmenInfo, &fishmenInfo)
	return
}

func (k Keeper) PenaltyBase(ctx sdk.Context) (penaltyBase uint64) {
	k.paramstore.Get(ctx, types.KeyPenaltyBase, &penaltyBase)
	return
}

func (k Keeper) MaxPenalty(ctx sdk.Context) (maxPenalty uint64) {
	k.paramstore.Get(ctx, types.KeyMaxPenalty, &maxPenalty)
	return
}

func (k Keeper) SetAnnualPercentageYield(ctx sdk.Context, apy string) {
	k.paramstore.Set(ctx, types.KeyAPY, &apy)
}

func (k Keeper) SetHalvingPeriod(ctx sdk.Context, halving int64) {
	k.paramstore.Set(ctx, types.KeyAPY, &halving)
}

func (k Keeper) SetAdjustmentPeriod(ctx sdk.Context, adjustment int64) {
	k.paramstore.Set(ctx, types.KeyAPY, &adjustment)
}

func (k Keeper) SetBaseline(ctx sdk.Context, baseline int64) {
	_baseline := k.Baseline(ctx)
	newBaseline := sdk.NewInt64Coin(_baseline.Denom, baseline)
	k.paramstore.Set(ctx, types.KeyBaseLine, newBaseline)
}

func (k Keeper) SetShareThreshold(ctx sdk.Context, threshold string) {
	k.paramstore.Set(ctx, types.KeyShareThreshold, threshold)
}

func (k Keeper) SetFishmenInfo(ctx sdk.Context, fishmenInfo string) {
	k.paramstore.Set(ctx, types.KeyFishmenInfo, &fishmenInfo)
}

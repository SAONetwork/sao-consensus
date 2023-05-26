package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.LoanInterest(ctx),
		k.MinLiquidityRatio(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// LoanInterest returns the LoanInterest param
func (k Keeper) LoanInterest(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyLoanInterest, &res)
	return
}

// MinLiquidityRatio returns the MinLiquidityRatio param
func (k Keeper) MinLiquidityRatio(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMinLiquidityRatio, &res)
	return
}

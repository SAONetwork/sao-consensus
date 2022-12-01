package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.BlockReward(ctx),
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

// EarnDenom returns the EarnDenom param
func (k Keeper) EarnDenom(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyEarnDenom, &res)
	return
}

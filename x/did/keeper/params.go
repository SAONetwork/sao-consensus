package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.GetBuiltinDids(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) GetBuiltinDids(ctx sdk.Context) (builtinDid string) {
	k.paramstore.Get(ctx, types.KeyBuiltinDid, &builtinDid)
	return
}

func (k Keeper) SetBuiltinDids(ctx sdk.Context, builtinDid string) {
	k.paramstore.Set(ctx, types.KeyBuiltinDid, &builtinDid)
}

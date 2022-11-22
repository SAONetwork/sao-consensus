package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) GetCosmosPaymentAddress(ctx sdk.Context, did string) sdk.AccAddress {
	return sdk.MustAccAddressFromBech32("cosmos1r33rtwtgak2erkwq2462l3ed2ry2q0p0427eu9")
}

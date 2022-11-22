package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) GetCosmosPaymentAddress(ctx sdk.Context, did string) sdk.AccAddress {
	return sdk.MustAccAddressFromBech32("cosmos1vpe97e0y80438pp3xpdrr3qs9v2g8f58u2rz2h")
}

package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) GetCosmosPaymentAddress(ctx sdk.Context, did string) sdk.AccAddress {
	return sdk.MustAccAddressFromBech32("cosmos1g68ahtuuxq8grzf6w8ns6tg7dzgh4ta55pqffy")
}

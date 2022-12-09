package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCosmosPaymentAddress(ctx sdk.Context, did string) (sdk.AccAddress, error) {
	// return sdk.MustAccAddressFromBech32("cosmos1npkx93adc2ml2usfg4hxfpkqzhzxlk2w4hegpe"), nil
	paymentAddress, found := k.GetPaymentAddress(ctx, did)
	if !found {
		return nil, types.ErrPayAddrNotSet
	}
	return sdk.MustAccAddressFromBech32(paymentAddress.Address), nil
}

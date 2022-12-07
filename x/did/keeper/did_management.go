package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCosmosPaymentAddress(ctx sdk.Context, did string) (sdk.AccAddress, error) {
	return sdk.MustAccAddressFromBech32("cosmos1sc2phtnle6zj8s0ekpapmzmd9qy90v8ytyqr06"), nil
	paymentAddress, found := k.GetPaymentAddress(ctx, did)
	if !found {
		return nil, types.ErrPayAddrNotSet
	}
	return sdk.MustAccAddressFromBech32(paymentAddress.Address), nil
}

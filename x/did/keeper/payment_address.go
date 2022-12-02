package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPaymentAddress set a specific paymentAddress in the store from its index
func (k Keeper) SetPaymentAddress(ctx sdk.Context, paymentAddress types.PaymentAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentAddressKeyPrefix))
	b := k.cdc.MustMarshal(&paymentAddress)
	store.Set(types.PaymentAddressKey(
		paymentAddress.Did,
	), b)
}

// GetPaymentAddress returns a paymentAddress from its index
func (k Keeper) GetPaymentAddress(
	ctx sdk.Context,
	did string,

) (val types.PaymentAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentAddressKeyPrefix))

	b := store.Get(types.PaymentAddressKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePaymentAddress removes a paymentAddress from the store
func (k Keeper) RemovePaymentAddress(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentAddressKeyPrefix))
	store.Delete(types.PaymentAddressKey(
		did,
	))
}

// GetAllPaymentAddress returns all paymentAddress
func (k Keeper) GetAllPaymentAddress(ctx sdk.Context) (list []types.PaymentAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaymentAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PaymentAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

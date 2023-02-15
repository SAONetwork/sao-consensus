package keeper

import (
	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetExpiredOrder set a specific expiredOrder in the store from its index
func (k Keeper) SetExpiredOrder(ctx sdk.Context, expiredOrder types.ExpiredOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredOrderKeyPrefix))
	b := k.cdc.MustMarshal(&expiredOrder)
	store.Set(types.ExpiredOrderKey(
		expiredOrder.Height,
	), b)
}

// GetExpiredOrder returns a expiredOrder from its index
func (k Keeper) GetExpiredOrder(
	ctx sdk.Context,
	height uint64,

) (val types.ExpiredOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredOrderKeyPrefix))

	b := store.Get(types.ExpiredOrderKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExpiredOrder removes a expiredOrder from the store
func (k Keeper) RemoveExpiredOrder(
	ctx sdk.Context,
	height uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredOrderKeyPrefix))
	store.Delete(types.ExpiredOrderKey(
		height,
	))
}

// GetAllExpiredOrder returns all expiredOrder
func (k Keeper) GetAllExpiredOrder(ctx sdk.Context) (list []types.ExpiredOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredOrderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExpiredOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

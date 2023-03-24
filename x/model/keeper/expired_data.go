package keeper

import (
	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetExpiredData set a specific expiredData in the store from its index
func (k Keeper) SetExpiredData(ctx sdk.Context, expiredData types.ExpiredData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredDataKeyPrefix))
	b := k.cdc.MustMarshal(&expiredData)
	store.Set(types.ExpiredDataKey(
		expiredData.Height,
	), b)
}

// GetExpiredData returns a expiredData from its index
func (k Keeper) GetExpiredData(
	ctx sdk.Context,
	height uint64,
) (val types.ExpiredData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredDataKeyPrefix))

	b := store.Get(types.ExpiredDataKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExpiredData removes a expiredData from the store
func (k Keeper) RemoveExpiredData(
	ctx sdk.Context,
	height uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredDataKeyPrefix))
	store.Delete(types.ExpiredDataKey(
		height,
	))
}

// GetAllExpiredData returns all expiredData
func (k Keeper) GetAllExpiredData(ctx sdk.Context) (list []types.ExpiredData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExpiredData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

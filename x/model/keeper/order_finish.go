package keeper

import (
	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetOrderFinish set a specific orderFinish in the store from its index
func (k Keeper) SetOrderFinish(ctx sdk.Context, orderFinish types.OrderFinish) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderFinishKeyPrefix))
	b := k.cdc.MustMarshal(&orderFinish)
	store.Set(types.OrderFinishKey(
		orderFinish.Height,
	), b)
}

// GetOrderFinish returns a orderFinish from its index
func (k Keeper) GetOrderFinish(
	ctx sdk.Context,
	timestamp uint64,
) (val types.OrderFinish, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderFinishKeyPrefix))

	b := store.Get(types.OrderFinishKey(
		timestamp,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrderFinish removes a orderFinish from the store
func (k Keeper) RemoveOrderFinish(
	ctx sdk.Context,
	timestamp uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderFinishKeyPrefix))
	store.Delete(types.OrderFinishKey(
		timestamp,
	))
}

// GetAllOrderFinish returns all orderFinish
func (k Keeper) GetAllOrderFinish(ctx sdk.Context) (list []types.OrderFinish) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderFinishKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OrderFinish
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

package keeper

import (
	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetModel set a specific model in the store from its index
func (k Keeper) SetModel(ctx sdk.Context, model types.Model) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	b := k.cdc.MustMarshal(&model)
	store.Set(types.ModelKey(
		model.Key,
	), b)
}

// GetModel returns a model from its index
func (k Keeper) GetModel(
	ctx sdk.Context,
	key string,

) (val types.Model, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))

	b := store.Get(types.ModelKey(
		key,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveModel removes a model from the store
func (k Keeper) RemoveModel(
	ctx sdk.Context,
	key string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	store.Delete(types.ModelKey(
		key,
	))
}

// GetAllModel returns all model
func (k Keeper) GetAllModel(ctx sdk.Context) (list []types.Model) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ModelKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Model
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

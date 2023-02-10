package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetShard set a specific shard in the store from its index
func (k Keeper) SetShard(ctx sdk.Context, shard types.Shard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKeyPrefix))
	b := k.cdc.MustMarshal(&shard)
	store.Set(types.ShardKey(
		shard.Idx,
	), b)
}

// GetShard returns a shard from its index
func (k Keeper) GetShard(
	ctx sdk.Context,
	idx string,

) (val types.Shard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKeyPrefix))

	b := store.Get(types.ShardKey(
		idx,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShard removes a shard from the store
func (k Keeper) RemoveShard(
	ctx sdk.Context,
	idx string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKeyPrefix))
	store.Delete(types.ShardKey(
		idx,
	))
}

// GetAllShard returns all shard
func (k Keeper) GetAllShard(ctx sdk.Context) (list []types.Shard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

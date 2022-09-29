package keeper

import (
	"encoding/binary"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetShardCount get the total number of shard
func (k Keeper) GetShardCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ShardCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetShardCount set the total number of shard
func (k Keeper) SetShardCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ShardCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendShard appends a shard in the store with a new id and update the count
func (k Keeper) AppendShard(
	ctx sdk.Context,
	shard types.Shard,
) uint64 {
	// Create the shard
	count := k.GetShardCount(ctx)

	// Set the ID of the appended value
	shard.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKey))
	appendedValue := k.cdc.MustMarshal(&shard)
	store.Set(GetShardIDBytes(shard.Id), appendedValue)

	// Update shard count
	k.SetShardCount(ctx, count+1)

	return count
}

// SetShard set a specific shard in the store
func (k Keeper) SetShard(ctx sdk.Context, shard types.Shard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKey))
	b := k.cdc.MustMarshal(&shard)
	store.Set(GetShardIDBytes(shard.Id), b)
}

// GetShard returns a shard from its id
func (k Keeper) GetShard(ctx sdk.Context, id uint64) (val types.Shard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKey))
	b := store.Get(GetShardIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShard removes a shard from the store
func (k Keeper) RemoveShard(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKey))
	store.Delete(GetShardIDBytes(id))
}

// GetAllShard returns all shard
func (k Keeper) GetAllShard(ctx sdk.Context) (list []types.Shard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ShardKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetShardIDBytes returns the byte representation of the ID
func GetShardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetShardIDFromBytes returns ID in uint64 format from a byte array
func GetShardIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

package keeper

import (
	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetExpiredShard set a specific expiredShard in the store from its index
func (k Keeper) SetExpiredShard(ctx sdk.Context, expiredShard types.ExpiredShard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredShardKeyPrefix))
	b := k.cdc.MustMarshal(&expiredShard)
	store.Set(types.ExpiredShardKey(
		expiredShard.Height,
	), b)
}

// GetExpiredShard returns a expiredShard from its index
func (k Keeper) GetExpiredShard(
	ctx sdk.Context,
	height uint64,

) (val types.ExpiredShard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredShardKeyPrefix))

	b := store.Get(types.ExpiredShardKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExpiredShard removes a expiredShard from the store
func (k Keeper) RemoveExpiredShard(
	ctx sdk.Context,
	height uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredShardKeyPrefix))
	store.Delete(types.ExpiredShardKey(
		height,
	))
}

// GetAllExpiredShard returns all expiredShard
func (k Keeper) GetAllExpiredShard(ctx sdk.Context) (list []types.ExpiredShard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExpiredShardKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExpiredShard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetExpiredShardBlock(ctx sdk.Context, shardId uint64, expiredAt uint64) {

	expiredShard, found := k.GetExpiredShard(ctx, expiredAt)
	if found {
		expiredShard.ShardList = append(expiredShard.ShardList, shardId)
	} else {
		expiredShard = types.ExpiredShard{
			Height:    expiredAt,
			ShardList: []uint64{shardId},
		}
	}

	k.SetExpiredShard(ctx, expiredShard)
}

func (k Keeper) SetExpiredShardsBlock(ctx sdk.Context, expiredShardsMap map[uint64][]uint64) {

	for expiredAt, shards := range expiredShardsMap {
		expiredShard, found := k.GetExpiredShard(ctx, expiredAt)
		if found {
			expiredShard.ShardList = append(expiredShard.ShardList, shards...)
		} else {
			expiredShard = types.ExpiredShard{
				Height:    expiredAt,
				ShardList: shards,
			}
		}

		k.SetExpiredShard(ctx, expiredShard)
	}
}

package v4

import (
	"encoding/binary"

	v3 "github.com/SaoNetwork/sao/x/node/migrations/v3/types"
	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, orderStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	pledgeList := GetAllV3Pledge(ctx, storeKey, cdc)
	price := sdk.NewDecWithPrec(1, 6)
	totalSize := int64(0)
	shardList := NewShardPledge(ctx, orderStoreKey, cdc)
	totalPledged := sdk.NewInt(0)
	for _, pledge := range pledgeList {
		amount := sdk.NewDecFromInt(pledge.TotalStoragePledged.Amount)
		if shardPledge, ok := shardList[pledge.Creator]; ok {
			amount = amount.Sub(shardPledge)
		}
		totalPledged = totalPledged.Add(amount.TruncateInt())
		size := amount.Quo(price).TruncateInt().Int64()
		newPledge := types.Pledge{
			Creator:             pledge.Creator,
			TotalStoragePledged: sdk.NewCoin(pledge.TotalStoragePledged.Denom, amount.TruncateInt()),
			Reward:              pledge.Reward,
			RewardDebt:          pledge.RewardDebt,
			UsedStorage:         pledge.TotalStorage,
			TotalStorage:        size,
		}
		totalSize += size
		SetPledge(ctx, newPledge, storeKey, cdc)
	}

	pool, _ := GetPool(ctx, storeKey, cdc)
	pool.TotalStorage = totalSize
	pool.TotalPledged.Amount = totalPledged
	SetPool(ctx, pool, storeKey, cdc)

	return nil
}

func NewShardPledge(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) map[string]sdk.Dec {

	shardPledgeList := make(map[string]sdk.Dec, 0)

	store := ctx.KVStore(storeKey)

	orderStore := prefix.NewStore(store, types.KeyPrefix(ordertypes.OrderKey))

	iterator := sdk.KVStorePrefixIterator(orderStore, []byte{})

	shardStore := prefix.NewStore(store, types.KeyPrefix(ordertypes.ShardKey))

	for ; iterator.Valid(); iterator.Next() {
		orderKey := iterator.Key()
		oldVal := orderStore.Get(orderKey)
		var order ordertypes.Order
		cdc.MustUnmarshal(oldVal, &order)

		shardPledge := order.UnitPrice.Amount.
			MulInt64(int64(order.Size_)).
			MulInt64(int64(order.Duration)).
			MulInt64(1).
			QuoInt64(10)

		for _, shardId := range order.Shards {
			shardKey := GetShardIDBytes(shardId)
			val := shardStore.Get(shardKey)
			var shard ordertypes.Shard
			cdc.MustUnmarshal(val, &shard)

			if _, ok := shardPledgeList[shard.Sp]; !ok {
				shardPledgeList[shard.Sp] = sdk.NewDec(0)
			}
			shard.Pledge.Amount = shardPledge.TruncateInt()
			shardPledgeList[shard.Sp] = shardPledgeList[shard.Sp].Add(sdk.NewDecFromInt(shard.Pledge.Amount))
			newVal := cdc.MustMarshal(&shard)
			shardStore.Set(shardKey, newVal)
		}
	}

	return shardPledgeList
}

func GetAllV3Pledge(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (list []v3.Pledge) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val v3.Pledge
		cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}

func SetPledge(ctx sdk.Context, pledge types.Pledge, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	b := cdc.MustMarshal(&pledge)
	store.Set(types.PledgeKey(
		pledge.Creator,
	), b)
}

func SetPool(ctx sdk.Context, pool types.Pool, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))
	b := cdc.MustMarshal(&pool)
	store.Set([]byte{0}, b)
}

// GetPool returns pool
func GetPool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetShardIDBytes returns the byte representation of the ID
func GetShardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

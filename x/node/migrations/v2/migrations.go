package v2

import (
	"encoding/binary"

	v1 "github.com/SaoNetwork/sao/x/node/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/node/types"
	v1order "github.com/SaoNetwork/sao/x/order/migrations/v1/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MigrateStore(ctx sdk.Context, bank types.BankKeeper, storeKey storetypes.StoreKey, orderStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	logger := ctx.Logger()
	// migrate shards

	shardStore := prefix.NewStore(ctx.KVStore(orderStoreKey), types.KeyPrefix(v1order.ShardKey))
	shardIterator := sdk.KVStorePrefixIterator(shardStore, []byte{})

	defer shardIterator.Close()
	logger.Debug("migrating shards")

	for ; shardIterator.Valid(); shardIterator.Next() {
		pool, _ := GetV1Pool(ctx, storeKey, cdc)
		var val v1order.Shard
		shardKey := shardIterator.Key()
		cdc.MustUnmarshal(shardIterator.Value(), &val)
		order, found := GetOrder(ctx, val.OrderId, orderStoreKey, cdc)
		if val.Status != ordertypes.ShardCompleted {
			shardStore.Delete(shardKey)
			continue
		}
		pledge, found := GetPledge(ctx, val.Sp, storeKey, cdc)
		if !found {
			shardStore.Delete(shardKey)
			continue
		}
		logger.Debug("migratign shard", "id", val.Id)
		if found && order.Status != ordertypes.OrderCompleted {
			var coins sdk.Coins

			if val.Pledge.IsZero() {
				shardStore.Delete(shardKey)
				continue
			}
			coins = coins.Add(val.Pledge)

			err := bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(val.Sp), coins)

			if err != nil {
				shardStore.Delete(shardKey)
				continue
			}

			pledge.TotalStorage -= int64(val.Size_)

			pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(val.Pledge)

			pool.TotalStorage -= int64(val.Size_)

			pool.TotalPledged = pool.TotalPledged.Sub(val.Pledge)

			SetV1Pool(ctx, storeKey, cdc, pool)

			shardStore.Delete(shardKey)
		}
	}

	// migrate pledge

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	pool, found := GetPool(ctx, storeKey, cdc)

	if !found {
		return status.Error(codes.NotFound, "pool not found")
	}

	totalReward := sdk.NewDecCoinFromCoin(pool.TotalReward)

	pool.AccRewardPerByte.Amount = totalReward.Amount.QuoInt64(pool.TotalStorage)

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		var val v1.Pledge
		cdc.MustUnmarshal(iterator.Value(), &val)
		pledge := types.Pledge{
			Creator:             val.Creator,
			TotalStoragePledged: val.TotalStoragePledged,
			Reward:              sdk.NewInt64DecCoin(val.Reward.Denom, 0),
			TotalStorage:        val.TotalStorage,
			RewardDebt:          sdk.NewInt64DecCoin(val.Reward.Denom, 0),
		}

		newVal := cdc.MustMarshal(&pledge)
		store.Set(key, newVal)
	}

	SetPool(ctx, storeKey, cdc, pool)

	return nil
}

func GetOrderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func SetPledge(ctx sdk.Context, pledge v1.Pledge, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	b := cdc.MustMarshal(&pledge)
	store.Set(types.PledgeKey(
		pledge.Creator,
	), b)
}

func GetPledge(ctx sdk.Context, creator string, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (val v1.Pledge, found bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PledgeKeyPrefix))

	b := store.Get(types.PledgeKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

func GetOrder(ctx sdk.Context, id uint64, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (val v1order.Order, found bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(v1order.OrderKey))
	b := store.Get(GetOrderIDBytes(id))
	if b == nil {
		return val, false
	}
	cdc.MustUnmarshal(b, &val)
	return val, true
}

func SetV1Pool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))
	b := cdc.MustMarshal(&pool)
	store.Set([]byte{0}, b)
}

func SetPool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))
	b := cdc.MustMarshal(&pool)
	store.Set([]byte{0}, b)
}

func GetV1Pool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

func GetPool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (val types.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

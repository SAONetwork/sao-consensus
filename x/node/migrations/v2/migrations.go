package v2

import (
	v1 "github.com/SaoNetwork/sao/x/node/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

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
			TotalOrderPledged:   val.TotalOrderPledged,
			TotalStoragePledged: val.TotalStoragePledged,
			Reward:              sdk.NewInt64DecCoin(val.Reward.Denom, 0),
			TotalStorage:        val.TotalStorage,
			RewardDebt:          sdk.NewInt64DecCoin(val.Reward.Denom, 0),
		}

		//pledge.RewardDebt.Amount = pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage)

		newVal := cdc.MustMarshal(&pledge)
		store.Set(key, newVal)
	}

	SetPool(ctx, storeKey, cdc, pool)

	return nil
}

func SetPool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))
	b := cdc.MustMarshal(&pool)
	store.Set([]byte{0}, b)
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

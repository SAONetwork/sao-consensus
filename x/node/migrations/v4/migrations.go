package v4

import (
	v3 "github.com/SaoNetwork/sao/x/node/migrations/v3/types"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	pledgeList := GetAllV3Pledge(ctx, storeKey, cdc)
	price := sdk.NewDecWithPrec(1, 6)
	totalSize := int64(0)
	for _, pledge := range pledgeList {
		amount := sdk.NewDecFromInt(pledge.TotalStoragePledged.Amount)
		size := amount.Quo(price).TruncateInt().Int64()
		newPledge := types.Pledge{
			Creator:             pledge.Creator,
			TotalStoragePledged: pledge.TotalStoragePledged,
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
	SetPool(ctx, pool, storeKey, cdc)

	return nil
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

package v3

import (
	"errors"

	v1 "github.com/SaoNetwork/sao/x/node/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	pool, found := GetV1Pool(ctx, storeKey, cdc)
	if !found {
		return errors.New("pool not found")
	}

	rewardPerBlock := pool.TotalReward.Amount.QuoRaw(pool.RewardedBlockCount)
	newPool := types.Pool{
		TotalPledged:       pool.TotalPledged,
		TotalReward:        pool.TotalReward,
		TotalStorage:       pool.TotalStorage,
		AccPledgePerByte:   pool.AccPledgePerByte,
		AccRewardPerByte:   pool.AccRewardPerByte,
		RewardPerBlock:     sdk.NewDecCoin(pool.TotalReward.Denom, rewardPerBlock),
		RewardedBlockCount: pool.RewardedBlockCount,
		NextRewardPerBlock: sdk.NewInt64DecCoin(pool.TotalReward.Denom, 0),
	}

	SetPool(ctx, storeKey, cdc, newPool)

	return nil
}

func UpdateNodeParams(ctx sdk.Context, paramStore *paramtypes.Subspace) error {
	var blockReward sdk.Coin
	//blockReward = paramstore
	paramStore.GetIfExists(ctx, types.KeyBlockReward, &blockReward)

	// set baseLine
	baseLine := sdk.NewInt64Coin(blockReward.Denom, 2000000)
	paramStore.Set(ctx, types.KeyBaseLine, &baseLine)

	// blockReward = 6.25 SAO
	blockReward = sdk.NewInt64Coin(blockReward.Denom, 6250000)
	paramStore.Set(ctx, types.KeyBlockReward, &blockReward)

	paramStore.Set(ctx, types.KeyHalvingPeriod, int64(32000000))
	paramStore.Set(ctx, types.KeyAdjustmentPeriod, int64(2000))

	// set apy
	paramStore.Set(ctx, types.KeyAPY, "8.0")

	return nil
}

func GetV1Pool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) (val v1.Pool, found bool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

func SetPool(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PoolKey))
	b := cdc.MustMarshal(&pool)
	store.Set([]byte{0}, b)
}

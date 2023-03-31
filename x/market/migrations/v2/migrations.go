package v2

import (
	v1 "github.com/SaoNetwork/sao/x/market/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/market/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.WorkerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	logger := ctx.Logger()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		var val v1.Worker

		cdc.MustUnmarshal(iterator.Value(), &val)
		logger.Debug("migrate worker", "name", val.Workername, "reward", val.Reward)
		worker := types.Worker{
			Workername:      val.Workername,
			Storage:         val.Storage,
			Reward:          val.Reward,
			IncomePerSecond: val.IncomePerSecond,
			LastRewardAt:    val.LastRewardAt,
		}

		if worker.Storage > 0 {
			reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockTime().Unix() - worker.LastRewardAt)
			worker.Reward.Amount = worker.Reward.Amount.Add(reward)
		}
		worker.LastRewardAt = ctx.BlockHeight()

		newVal := cdc.MustMarshal(&worker)
		store.Set(key, newVal)
	}

	return nil
}

package v2

import (
	"github.com/SaoNetwork/sao/x/market/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WorkerAppend func(ctx sdk.Context, order *ordertypes.Order, shard *ordertypes.Shard) error

func MigrateStore(ctx sdk.Context, workerAppend WorkerAppend, storeKey storetypes.StoreKey, orderKeeper types.OrderKeeper) error {

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.WorkerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		store.Delete(key)
	}

	orders := orderKeeper.GetAllOrder(ctx)
	for _, order := range orders {
		if order.Status != ordertypes.OrderCompleted {
			continue
		}

		for _, shardId := range order.Shards {
			shard, _ := orderKeeper.GetShard(ctx, shardId)
			workerAppend(ctx, &order, &shard)
		}
	}

	return nil
}

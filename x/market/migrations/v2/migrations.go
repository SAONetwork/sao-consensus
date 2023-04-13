package v2

import (
	"github.com/SaoNetwork/sao/x/market/types"
	v1order "github.com/SaoNetwork/sao/x/order/migrations/v1/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WorkerAppend func(ctx sdk.Context, order *ordertypes.Order, shard *ordertypes.Shard) error

func MigrateStore(ctx sdk.Context, workerAppend WorkerAppend, storeKey storetypes.StoreKey, orderStoreKey storetypes.StoreKey, cdc codec.BinaryCodec, orderKeeper types.OrderKeeper) error {

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.WorkerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		store.Delete(key)
	}

	orderStore := prefix.NewStore(ctx.KVStore(orderStoreKey), types.KeyPrefix(ordertypes.OrderKey))
	orderIterator := sdk.KVStorePrefixIterator(orderStore, []byte{})
	defer orderIterator.Close()

	for ; orderIterator.Valid(); orderIterator.Next() {

		var orderOld v1order.Order
		var order ordertypes.Order
		err := cdc.Unmarshal(orderIterator.Value(), &orderOld)

		if err != nil {
			cdc.MustUnmarshal(orderIterator.Value(), &order)
		} else {
			order = ordertypes.Order{
				Id:       orderOld.Id,
				Amount:   orderOld.Amount,
				Shards:   orderOld.Shards,
				Status:   orderOld.Status,
				Duration: orderOld.Duration,
				Replica:  orderOld.Replica,
			}
		}

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

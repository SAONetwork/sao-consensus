package v2

import (
	"encoding/binary"
	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	logger := ctx.Logger()
	logger.Debug("migrating market")

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.OrderKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	shardStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ShardKey))

	for ; iterator.Valid(); iterator.Next() {

		var order types.Order
		err := cdc.Unmarshal(iterator.Value(), &order)
		if err != nil {
			return err
		}

		if order.Status == types.OrderCompleted {
			count := order.Replica
			timeoutShards := []uint64{}
			normalShards := []uint64{}
			for _, shardId := range order.Shards {
				shardBytes := shardStore.Get(GetShardIDBytes(shardId))
				var shard types.Shard
				err := cdc.Unmarshal(shardBytes, &shard)
				if err != nil {
					return err
				}
				if shard.Status == types.ShardTimeout {
					timeoutShards = append(timeoutShards, shardId)
				} else if shard.Status == types.ShardCompleted || shard.Status == types.ShardMigrating {
					normalShards = append(normalShards, shardId)
					count--
				}
			}

			if count == 0 {
				for _, shardId := range timeoutShards {
					shardStore.Delete(GetShardIDBytes(shardId))
				}
				order.Shards = normalShards
				b := cdc.MustMarshal(&order)
				store.Set(GetOrderIDBytes(order.Id), b)
			}
		}
	}

	return nil
}

func GetShardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetOrderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

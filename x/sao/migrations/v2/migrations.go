package v2

import (
	v2order "github.com/SaoNetwork/sao/x/order/migrations/v2/types"
	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SetExpiredShardsBlock func(ctx sdk.Context, expiredShardsMap map[uint64][]uint64)

func MigrateStore(ctx sdk.Context, setExpiredShardsBlock SetExpiredShardsBlock, orderStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	orderstore := ctx.KVStore(orderStoreKey)

	orderStore := prefix.NewStore(orderstore, types.KeyPrefix(types.OrderKey))

	orderIterator := sdk.KVStorePrefixIterator(orderStore, []byte{})

	expiredShardsMap := make(map[uint64][]uint64)

	for ; orderIterator.Valid(); orderIterator.Next() {
		orderKey := orderIterator.Key()
		oldVal := orderStore.Get(orderKey)
		var order v2order.Order
		cdc.MustUnmarshal(oldVal, &order)

		expiredAt := order.CreatedAt + order.Duration
		expiredShardsMap[expiredAt] = append(expiredShardsMap[expiredAt], order.Shards...)
	}

	setExpiredShardsBlock(ctx, expiredShardsMap)

	return nil
}

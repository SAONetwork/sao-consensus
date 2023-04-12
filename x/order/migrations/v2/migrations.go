package v2

import (
	v1 "github.com/SaoNetwork/sao/x/model/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type RefundOrder func(ctx sdk.Context, orderId uint64) error

func MigrateStore(ctx sdk.Context, refund RefundOrder, storeKey storetypes.StoreKey, modelStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))

	iterator := sdk.KVStorePrefixIterator(orderStore, []byte{})

	logger := ctx.Logger()

	for ; iterator.Valid(); iterator.Next() {
		orderKey := iterator.Key()
		oldVal := orderStore.Get(orderKey)
		var order types.Order
		cdc.MustUnmarshal(oldVal, &order)

		_, found := GetMetadata(ctx, modelStoreKey, order.Metadata.DataId, cdc)
		if found {
			order.CreatedAt = uint64(ctx.BlockHeight())
			order.Timeout = 86400
			order.DataId = order.Metadata.DataId
			order.Commit = order.Metadata.Commit
			newVal := cdc.MustMarshal(&order)
			orderStore.Set(orderKey, newVal)
			logger.Debug("migrate order created_at", "order", order.Id, "created_at", order.CreatedAt)
		} else {
			logger.Debug("remote order", "order", order.Id)
			orderStore.Delete(orderKey)
			refund(ctx, order.Id)
		}

	}

	return nil
}

func GetMetadata(ctx sdk.Context, storeKey storetypes.StoreKey, dataId string, cdc codec.BinaryCodec) (val v1.Metadata, found bool) {
	//storeKey := storetypes.NewKVStoreKey("model")
	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(v1.MetadataKeyPrefix))

	b := store.Get(v1.MetadataKey(
		dataId,
	))
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

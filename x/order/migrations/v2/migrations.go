package v2

import (
	v1 "github.com/SaoNetwork/sao/x/model/migrations/v1/types"
	v1order "github.com/SaoNetwork/sao/x/order/migrations/v1/types"
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
		var order v1order.Order
		cdc.MustUnmarshal(oldVal, &order)

		newOrder := types.Order{
			Creator:   order.Creator,
			Owner:     order.Owner,
			Id:        order.Id,
			Provider:  order.Provider,
			Cid:       order.Cid,
			Duration:  order.Duration,
			Status:    order.Status,
			Replica:   order.Replica,
			Metadata:  (*types.Metadata)(order.Metadata),
			Shards:    order.Shards,
			Amount:    order.Amount,
			Size_:     order.Size_,
			Operation: order.Operation,
		}

		_, found := GetMetadata(ctx, modelStoreKey, order.Metadata.DataId, cdc)
		if found {
			newOrder.CreatedAt = uint64(ctx.BlockHeight())
			newOrder.Timeout = 86400
			newOrder.DataId = order.Metadata.DataId
			newOrder.Commit = order.Metadata.Commit
			newVal := cdc.MustMarshal(&newOrder)
			orderStore.Set(orderKey, newVal)
			logger.Debug("migrate order created_at", "order", order.Id, "created_at", order.CreatedAt)
		} else {
			refund(ctx, order.Id)
			logger.Debug("remove order", "order", order.Id)
			orderStore.Delete(orderKey)
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

package v3

import (
	v2 "github.com/SaoNetwork/sao/x/model/migrations/v2/types"
	"github.com/SaoNetwork/sao/x/model/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, orderStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	orderStore := prefix.NewStore(ctx.KVStore(orderStoreKey), types.KeyPrefix(ordertypes.OrderKey))
	orderIterator := sdk.KVStorePrefixIterator(orderStore, []byte{})
	defer orderIterator.Close()

	orderIdMap := make(map[string][]uint64)
	for ; orderIterator.Valid(); orderIterator.Next() {
		var val ordertypes.Order
		cdc.MustUnmarshal(orderIterator.Value(), &val)

		if _, ok := orderIdMap[val.DataId]; !ok {
			orderIdMap[val.DataId] = []uint64{val.Id}
		} else {
			orderIdMap[val.DataId] = append(orderIdMap[val.DataId], val.Id)
		}
	}

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.MetadataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		var val v2.Metadata
		cdc.MustUnmarshal(iterator.Value(), &val)

		metadata := types.Metadata{
			DataId:        val.DataId,
			Owner:         val.Owner,
			Alias:         val.Alias,
			GroupId:       val.GroupId,
			OrderId:       val.OrderId,
			Tags:          val.Tags,
			Cid:           val.Cid,
			Commits:       val.Commits,
			ExtendInfo:    val.ExtendInfo,
			Update:        val.Update,
			Commit:        val.Commit,
			Rule:          val.Rule,
			Duration:      val.Duration,
			CreatedAt:     val.CreatedAt,
			ReadonlyDids:  val.ReadonlyDids,
			ReadwriteDids: val.ReadwriteDids,
			Status:        val.Status,
			Orders:        orderIdMap[val.DataId],
		}

		newVal := cdc.MustMarshal(&metadata)
		store.Set(key, newVal)
	}

	return nil
}

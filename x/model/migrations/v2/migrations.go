package v2

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	v1 "github.com/SaoNetwork/sao/x/model/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/model/types"
	v1order "github.com/SaoNetwork/sao/x/order/migrations/v1/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SetDataExpireBlock = func(ctx sdk.Context, dataId string, expiredAt uint64)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, orderStoreKey storetypes.StoreKey, cdc codec.BinaryCodec, setDataExpireBlock SetDataExpireBlock) error {
	edStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.ExpiredDataKeyPrefix))
	edIterator := sdk.KVStorePrefixIterator(edStore, []byte{})
	defer edIterator.Close()

	for ; edIterator.Valid(); edIterator.Next() {
		key := edIterator.Key()
		edStore.Delete(key)
	}

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.MetadataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	orderStore := prefix.NewStore(ctx.KVStore(orderStoreKey), types.KeyPrefix(ordertypes.OrderKey))

	defer iterator.Close()

	currentHeight := ctx.BlockHeight()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		var val v1.Metadata
		cdc.MustUnmarshal(iterator.Value(), &val)

		var duration uint64
		var valOrder v1order.Order
		orderBytes := orderStore.Get(GetOrderIDBytes(val.OrderId))
		err := cdc.Unmarshal(orderBytes, &valOrder)
		if err != nil {
			var valOrder ordertypes.Order
			cdc.MustUnmarshal(orderBytes, &valOrder)
			duration = valOrder.Duration
		} else {
			duration = valOrder.Duration
		}

		for i, _ := range val.Commits {
			val.Commits[i] = Version(CommitFromVersion(val.Commits[i]), currentHeight)
		}

		meta := types.Metadata{
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
			Duration:      duration,
			CreatedAt:     uint64(currentHeight),
			ReadonlyDids:  val.ReadonlyDids,
			ReadwriteDids: val.ReadwriteDids,
			Status:        types.MetaComplete,
		}

		setDataExpireBlock(ctx, meta.DataId, meta.CreatedAt+meta.Duration)
		newVal := cdc.MustMarshal(&meta)
		store.Set(key, newVal)
	}

	return nil
}

func Version(commit string, height int64) string {
	version := bytes.NewBufferString(commit)
	version.WriteByte(26)
	version.WriteString(fmt.Sprintf("%d", height))
	return version.String()
}

func CommitFromVersion(version string) string {
	splited := strings.Split(version, string([]uint8{26}))
	return splited[0]
}

// GetOrderIDBytes returns the byte representation of the ID
func GetOrderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

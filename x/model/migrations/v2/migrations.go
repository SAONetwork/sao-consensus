package v2

import (
	"bytes"
	"fmt"
	v1 "github.com/SaoNetwork/sao/x/model/migrations/v1/types"
	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type SetDataExpireBlock = func(ctx sdk.Context, dataId string, expiredAt uint64)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, orderKeeper types.OrderKeeper, setDataExpireBlock SetDataExpireBlock) error {

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.MetadataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	currentHeight := ctx.BlockHeight()

	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		var val v1.Metadata
		cdc.MustUnmarshal(iterator.Value(), &val)

		order, _ := orderKeeper.GetOrder(ctx, val.OrderId)

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
			Duration:      order.Duration,
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

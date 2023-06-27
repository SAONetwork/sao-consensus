package v4

import (
	"bytes"
	"encoding/binary"
	"fmt"
	markettypes "github.com/SaoNetwork/sao/x/market/types"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Deposit = func(coin sdk.Coin) error

func MigrateStore(ctx sdk.Context, storeKey, modelStoreKey, marketStoreKey storetypes.StoreKey, cdc codec.BinaryCodec, deposit Deposit) error {
	store := ctx.KVStore(storeKey)
	modelStore := ctx.KVStore(modelStoreKey)

	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))
	shardStore := prefix.NewStore(store, types.KeyPrefix(types.ShardKey))

	metadataStore := prefix.NewStore(modelStore, modeltypes.KeyPrefix(modeltypes.MetadataKeyPrefix))

	workerStore := prefix.NewStore(ctx.KVStore(marketStoreKey), markettypes.KeyPrefix(markettypes.WorkerKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(orderStore, []byte{})
	for ; iterator.Valid(); iterator.Next() {
		orderKey := iterator.Key()
		oldVal := orderStore.Get(orderKey)
		var order types.Order
		cdc.MustUnmarshal(oldVal, &order)

		// merge InProgress into Complete
		if order.Status == types.OrderInProgress {
			// order status update
			order.Status = types.OrderCompleted

			// deposit
			err := deposit(order.Amount)
			if err != nil {
				return err
			}

			// worker append
			for _, shardId := range order.Shards {
				shard, foundShard := GetShard(shardStore, shardId, cdc)
				if foundShard && shard.Status == types.ShardCompleted {
					WorkerAppend(ctx, workerStore, order, shard, cdc)
				}
			}

			// meta status update
			meta, found := GetMetadata(metadataStore, order.DataId, cdc)
			if found {
				if meta.OrderId < order.Id {
					meta.OrderId = order.Id
				}
				if meta.Commit == order.Commit && meta.Status != modeltypes.MetaComplete {
					meta.Cid = order.Cid
					meta.Commits = append(meta.Commits, Version(order.Commit, ctx.BlockHeight()))
					meta.Orders = append(meta.Orders, order.Id)
					meta.Status = modeltypes.MetaComplete
				}
			}
			SetOrder(orderStore, order, cdc)
			SetMetadata(metadataStore, meta, cdc)
		}
	}

	return nil
}

func SetOrder(store prefix.Store, order types.Order, cdc codec.BinaryCodec) {
	b := cdc.MustMarshal(&order)
	store.Set(GetOrderIDBytes(order.Id), b)
}

func GetOrderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func GetMetadata(store prefix.Store, dataId string, cdc codec.BinaryCodec) (val modeltypes.Metadata, found bool) {
	b := store.Get(modeltypes.MetadataKey(
		dataId,
	))
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

func SetMetadata(store prefix.Store, metadata modeltypes.Metadata, cdc codec.BinaryCodec) {
	b := cdc.MustMarshal(&metadata)
	store.Set(modeltypes.MetadataKey(
		metadata.DataId,
	), b)
}

func GetShard(store prefix.Store, id uint64, cdc codec.BinaryCodec) (val types.Shard, found bool) {
	b := store.Get(GetShardIDBytes(id))
	if b == nil {
		return val, false
	}
	cdc.MustUnmarshal(b, &val)
	return val, true
}

func GetShardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func Version(commit string, height int64) string {
	version := bytes.NewBufferString(commit)
	version.WriteByte(26)
	version.WriteString(fmt.Sprintf("%d", height))
	return version.String()
}

func WorkerAppend(ctx sdk.Context, store prefix.Store, order types.Order, shard types.Shard, cdc codec.BinaryCodec) {

	amount := sdk.NewDecCoinFromCoin(order.Amount)

	workerName := fmt.Sprintf("%s-%s", amount.Denom, shard.Sp)
	worker, found := GetWorker(store, workerName, cdc)
	if !found {
		worker = markettypes.Worker{
			Workername:      workerName,
			Storage:         0,
			Reward:          sdk.NewInt64DecCoin(amount.Denom, 0),
			IncomePerSecond: sdk.NewInt64DecCoin(amount.Denom, 0),
		}
	}

	IncomePerBlock := order.UnitPrice.Amount.MulInt64(int64(shard.Size_))
	reward := IncomePerBlock.MulInt64(ctx.BlockHeight() - int64(shard.CreatedAt))
	if worker.Storage > 0 {
		reward = reward.Add(worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt))
	}
	worker.Reward.Amount = worker.Reward.Amount.Add(reward)
	worker.LastRewardAt = ctx.BlockHeight()

	worker.Storage += shard.Size_
	worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Add(IncomePerBlock)

	SetWorker(store, worker, cdc)
}

func GetWorker(store prefix.Store, workername string, cdc codec.BinaryCodec) (val markettypes.Worker, found bool) {
	b := store.Get(markettypes.WorkerKey(
		workername,
	))
	if b == nil {
		return val, false
	}

	cdc.MustUnmarshal(b, &val)
	return val, true
}

func SetWorker(store prefix.Store, worker markettypes.Worker, cdc codec.BinaryCodec) {
	b := cdc.MustMarshal(&worker)
	store.Set(markettypes.WorkerKey(
		worker.Workername,
	), b)
}

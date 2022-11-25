package keeper

import (
	"github.com/SaoNetwork/sao/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetWorker set a specific worker in the store from its index
func (k Keeper) SetWorker(ctx sdk.Context, worker types.Worker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorkerKeyPrefix))
	b := k.cdc.MustMarshal(&worker)
	store.Set(types.WorkerKey(
		worker.Workername,
	), b)
}

// GetWorker returns a worker from its index
func (k Keeper) GetWorker(
	ctx sdk.Context,
	workername string,

) (val types.Worker, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorkerKeyPrefix))

	b := store.Get(types.WorkerKey(
		workername,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWorker removes a worker from the store
func (k Keeper) RemoveWorker(
	ctx sdk.Context,
	workername string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorkerKeyPrefix))
	store.Delete(types.WorkerKey(
		workername,
	))
}

// GetAllWorker returns all worker
func (k Keeper) GetAllWorker(ctx sdk.Context) (list []types.Worker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorkerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Worker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

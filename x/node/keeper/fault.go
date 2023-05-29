package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	uuid "github.com/satori/go.uuid"
)

const NS_URL = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"

// SetFault set a specific fault in the store from its index
func (k Keeper) SetFault(ctx sdk.Context, fault *types.Fault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))

	if fault.Status == types.FaultStatusConfirming && fault.FaultId == "" {
		fault.FaultId = generateFaultId(fault)
	}
	b := k.cdc.MustMarshal(fault)
	store.Set(types.FaultKey(
		fault.Provider,
		fault.ShardId,
	), b)
}

// GetFault returns a fault from its index
func (k Keeper) GetFault(
	ctx sdk.Context,
	provider string,
	shardId string,
) (val *types.Fault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))

	b := store.Get(types.FaultKey(
		provider,
		shardId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

// RemoveFault removes a fault from the store
func (k Keeper) RemoveFault(
	ctx sdk.Context,
	provider string,
	shardId string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	store.Delete(types.FaultKey(
		provider,
		shardId,
	))
}

// GetAllFaults returns all faults
func (k Keeper) GetAllFault(ctx sdk.Context) (list []types.Fault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fault
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllFaultsByStatus returns all faults with the expected status
func (k Keeper) GetAllFaultsByStatus(ctx sdk.Context, status uint32) (list []types.Fault) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Fault
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if status&n.Status == status {
			list = append(list, n)
		}
	}

	return
}

func generateFaultId(fault *types.Fault) string {
	seed := fault.Provider + fault.Reporter + fault.CommitId + fault.ShardId
	return uuid.NewV5(uuid.FromStringOrNil(NS_URL), seed).String()
}

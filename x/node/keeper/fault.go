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
	), []byte(fault.FaultId))

	store.Set([]byte(fault.FaultId), b)
}

// GetFault returns a fault from its index
func (k Keeper) GetFault(
	ctx sdk.Context,
	faultId string,
) (val *types.Fault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))

	b := store.Get([]byte(faultId))
	if b == nil {
		return nil, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) GetFaultBySpAndShardId(
	ctx sdk.Context,
	provider string,
	shardId string,
) (val *types.Fault, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))

	faultIdBytes := store.Get(types.FaultKey(
		provider,
		shardId,
	))
	if faultIdBytes == nil {
		return nil, false
	}

	return k.GetFault(ctx, string(faultIdBytes))
}

// RemoveFault remove a fault from stre
func (k Keeper) RemoveFault(
	ctx sdk.Context,
	faultId string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	store.Delete([]byte(faultId))
}

// GetFaultsByStatus returns all faults with the expected status
func (k Keeper) GetFaultsByStatus(ctx sdk.Context, status uint32) (list []types.Fault) {
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

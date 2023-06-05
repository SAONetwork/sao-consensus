package keeper

import (
	"strconv"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	uuid "github.com/satori/go.uuid"
)

const NS_URL = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"

// SetFault set a specific fault in the store from its index
func (k Keeper) SetFault(ctx sdk.Context, fault *types.Fault) {
	IdStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultIdKeyPrefix))
	faultStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))

	if fault.Status == types.FaultStatusConfirming && fault.FaultId == "" {
		fault.FaultId = generateFaultId(fault)
	}
	b := k.cdc.MustMarshal(fault)
	faultStore.Set(types.FaultKey(
		fault.Provider,
		fault.ShardId,
	), []byte(fault.FaultId))

	IdStore.Set([]byte(fault.FaultId), b)
}

// GetFault returns a fault from its index
func (k Keeper) GetFault(
	ctx sdk.Context,
	faultId string,
) (val *types.Fault, found bool) {
	IdStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultIdKeyPrefix))

	faultBytes := IdStore.Get([]byte(faultId))
	if faultBytes == nil {
		k.Logger(ctx).Error("fault not found")
		return nil, false
	}

	var fault types.Fault
	err := k.cdc.Unmarshal(faultBytes, &fault)
	if err != nil {
		k.Logger(ctx).Error("unmarshal failed," + err.Error())
		return nil, false
	}

	return &fault, true
}

func (k Keeper) GetFaultBySpAndShardId(
	ctx sdk.Context,
	provider string,
	shardId uint64,
) (val *types.Fault, found bool) {
	faultStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	idStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultIdKeyPrefix))

	faultIdBytes := faultStore.Get(types.FaultKey(
		provider,
		shardId,
	))
	if faultIdBytes == nil {
		return nil, false
	}

	faultBytes := idStore.Get(faultIdBytes)
	if faultBytes == nil {
		k.Logger(ctx).Error("fault not found")
		faultStore.Delete(types.FaultKey(
			provider,
			shardId,
		))

		return nil, false
	}

	var fault types.Fault
	err := k.cdc.Unmarshal(faultBytes, &fault)
	if err != nil {
		k.Logger(ctx).Error("unmarshal failed," + err.Error())
		return nil, false
	}

	return &fault, true
}

// RemoveFault remove a fault from stre
func (k Keeper) RemoveFault(
	ctx sdk.Context,
	fault *types.Fault,
) {
	faultStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	IdStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultIdKeyPrefix))

	IdStore.Delete([]byte(fault.FaultId))
	faultStore.Delete(types.FaultKey(
		fault.Provider,
		fault.ShardId,
	))
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
	seed := fault.Provider + fault.Reporter + fault.CommitId + strconv.FormatUint(fault.ShardId, 10)
	return uuid.NewV5(uuid.FromStringOrNil(NS_URL), seed).String()
}

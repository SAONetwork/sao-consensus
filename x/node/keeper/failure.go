package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFailure set a specific node in the store from its index
func (k Keeper) SetFailure(ctx sdk.Context, node types.Failure) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))
	b := k.cdc.MustMarshal(&node)
	store.Set(types.FailureKey(
		node.Creator,
	), b)
}

// GetFailure returns a node from its index
func (k Keeper) GetFailure(
	ctx sdk.Context,
	creator string,

) (val types.Failure, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))

	b := store.Get(types.FailureKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFailure removes a node from the store
func (k Keeper) RemoveFailure(
	ctx sdk.Context,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))
	store.Delete(types.FailureKey(
		creator,
	))
}

// GetAllFailures returns all nodes
func (k Keeper) GetAllFailure(ctx sdk.Context) (list []types.Failure) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Failure
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllFailuresByStatus returns all nodes with the expected status
func (k Keeper) GetAllFailuresByStatus(ctx sdk.Context, status uint32) (list []types.Failure) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Failure
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if status&n.Status == status {
			list = append(list, n)
		}
	}

	return
}

// GetAllFailuresByStatus returns all nodes with the expected status and reputation
func (k Keeper) GetAllFailuresByStatusAndReputation(ctx sdk.Context, status uint32, reputation float32) (list []types.Failure) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Failure
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if status&n.Status == status && n.Reputation >= reputation {
			list = append(list, n)
		}
	}

	return
}

func (k Keeper) EndBlock(ctx sdk.Context) {
	if ctx.BlockHeight()%1800 == 0 {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FailureKeyPrefix))
		iterator := sdk.KVStorePrefixIterator(store, []byte{})

		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			var n types.Failure
			k.cdc.MustUnmarshal(iterator.Value(), &n)
			if n.Status&types.NODE_STATUS_ONLINE > 0 && ctx.BlockHeight()-n.LastAliveHeight > 3600 {
				n.Status = n.Status & (types.NODE_STATUS_NA ^ types.NODE_STATUS_ONLINE)
				b := k.cdc.MustMarshal(&n)
				store.Set(types.FailureKey(
					n.Creator,
				), b)
			}

			if n.Status&types.NODE_STATUS_ONLINE == 0 || ctx.BlockHeight()-n.LastAliveHeight > 10800 {
				store.Delete(types.FailureKey(n.Creator))
			}
		}
	}
}

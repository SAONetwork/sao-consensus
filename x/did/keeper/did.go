package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDid set a specific did in the store from its index
func (k Keeper) SetDid(ctx sdk.Context, did types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	b := k.cdc.MustMarshal(&did)
	store.Set(types.DidKey(
		did.AccountId,
	), b)
}

// GetDid returns a did from its index
func (k Keeper) GetDid(
	ctx sdk.Context,
	accountId string,

) (val types.Did, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))

	b := store.Get(types.DidKey(
		accountId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDid removes a did from the store
func (k Keeper) RemoveDid(
	ctx sdk.Context,
	accountId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	store.Delete(types.DidKey(
		accountId,
	))
}

// GetAllDid returns all did
func (k Keeper) GetAllDid(ctx sdk.Context) (list []types.Did) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Did
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

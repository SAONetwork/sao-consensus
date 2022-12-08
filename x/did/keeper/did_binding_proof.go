package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDidBingingProof set a specific DidBingingProof in the store from its index
func (k Keeper) SetDidBingingProof(ctx sdk.Context, DidBingingProof types.DidBingingProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBingingProofKeyPrefix))
	b := k.cdc.MustMarshal(&DidBingingProof)
	store.Set(types.DidBingingProofKey(
		DidBingingProof.AccountId,
	), b)
}

// GetDidBingingProof returns a DidBingingProof from its index
func (k Keeper) GetDidBingingProof(
	ctx sdk.Context,
	accountId string,

) (val types.DidBingingProof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBingingProofKeyPrefix))

	b := store.Get(types.DidBingingProofKey(
		accountId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDidBingingProof removes a DidBingingProof from the store
func (k Keeper) RemoveDidBingingProof(
	ctx sdk.Context,
	accountId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBingingProofKeyPrefix))
	store.Delete(types.DidBingingProofKey(
		accountId,
	))
}

// GetAllDidBingingProof returns all DidBingingProof
func (k Keeper) GetAllDidBingingProof(ctx sdk.Context) (list []types.DidBingingProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBingingProofKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidBingingProof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

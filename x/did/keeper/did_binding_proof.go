package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDidBindingProof set a specific DidBindingProof in the store from its index
func (k Keeper) SetDidBindingProof(ctx sdk.Context, DidBindingProof types.DidBindingProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofKeyPrefix))
	b := k.cdc.MustMarshal(&DidBindingProof)
	store.Set(types.DidBindingProofKey(
		DidBindingProof.AccountId,
	), b)
}

// GetDidBindingProof returns a DidBindingProof from its index
func (k Keeper) GetDidBindingProof(
	ctx sdk.Context,
	accountId string,

) (val types.DidBindingProof, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofKeyPrefix))

	b := store.Get(types.DidBindingProofKey(
		accountId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDidBindingProof removes a DidBindingProof from the store
func (k Keeper) RemoveDidBindingProof(
	ctx sdk.Context,
	accountId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofKeyPrefix))
	store.Delete(types.DidBindingProofKey(
		accountId,
	))
}

// GetAllDidBindingProof returns all DidBindingProof
func (k Keeper) GetAllDidBindingProof(ctx sdk.Context) (list []types.DidBindingProof) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidBindingProof
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

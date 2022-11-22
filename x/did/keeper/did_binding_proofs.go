package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDidBindingProofs set a specific didBindingProofs in the store from its index
func (k Keeper) SetDidBindingProofs(ctx sdk.Context, didBindingProofs types.DidBindingProofs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofsKeyPrefix))
	b := k.cdc.MustMarshal(&didBindingProofs)
	store.Set(types.DidBindingProofsKey(
		didBindingProofs.AccountId,
	), b)
}

// GetDidBindingProofs returns a didBindingProofs from its index
func (k Keeper) GetDidBindingProofs(
	ctx sdk.Context,
	accountId string,

) (val types.DidBindingProofs, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofsKeyPrefix))

	b := store.Get(types.DidBindingProofsKey(
		accountId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDidBindingProofs removes a didBindingProofs from the store
func (k Keeper) RemoveDidBindingProofs(
	ctx sdk.Context,
	accountId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofsKeyPrefix))
	store.Delete(types.DidBindingProofsKey(
		accountId,
	))
}

// GetAllDidBindingProofs returns all didBindingProofs
func (k Keeper) GetAllDidBindingProofs(ctx sdk.Context) (list []types.DidBindingProofs) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBindingProofsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidBindingProofs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

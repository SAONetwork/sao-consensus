package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKid set a specific kid in the store from its index
func (k Keeper) SetKid(ctx sdk.Context, kid types.Kid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KidKeyPrefix))
	b := k.cdc.MustMarshal(&kid)
	store.Set(types.KidKey(
		kid.Address,
	), b)
}

// GetKid returns a kid from its index
func (k Keeper) GetKid(
	ctx sdk.Context,
	address string,

) (val types.Kid, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KidKeyPrefix))

	b := store.Get(types.KidKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKid removes a kid from the store
func (k Keeper) RemoveKid(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KidKeyPrefix))
	store.Delete(types.KidKey(
		address,
	))
}

// GetAllKid returns all kid
func (k Keeper) GetAllKid(ctx sdk.Context) (list []types.Kid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KidKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Kid
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

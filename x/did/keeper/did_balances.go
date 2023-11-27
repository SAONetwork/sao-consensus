package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetDidBalances set a specific didBalances in the store from its index
func (k Keeper) SetDidBalances(ctx sdk.Context, didBalances types.DidBalances) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBalancesKeyPrefix))
	b := k.cdc.MustMarshal(&didBalances)
	store.Set(types.DidBalancesKey(
		didBalances.Did,
	), b)
}

// GetDidBalances returns a didBalances from its index
func (k Keeper) GetDidBalances(
	ctx sdk.Context,
	did string,

) (val types.DidBalances, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBalancesKeyPrefix))

	b := store.Get(types.DidBalancesKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDidBalances removes a didBalances from the store
func (k Keeper) RemoveDidBalances(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBalancesKeyPrefix))
	store.Delete(types.DidBalancesKey(
		did,
	))
}

// GetAllDidBalances returns all didBalances
func (k Keeper) GetAllDidBalances(ctx sdk.Context) (list []types.DidBalances) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidBalancesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DidBalances
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

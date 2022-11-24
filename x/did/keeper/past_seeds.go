package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPastSeeds set a specific pastSeeds in the store from its index
func (k Keeper) SetPastSeeds(ctx sdk.Context, pastSeeds types.PastSeeds) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastSeedsKeyPrefix))
	b := k.cdc.MustMarshal(&pastSeeds)
	store.Set(types.PastSeedsKey(
		pastSeeds.Did,
	), b)
}

// GetPastSeeds returns a pastSeeds from its index
func (k Keeper) GetPastSeeds(
	ctx sdk.Context,
	did string,

) (val types.PastSeeds, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastSeedsKeyPrefix))

	b := store.Get(types.PastSeedsKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePastSeeds removes a pastSeeds from the store
func (k Keeper) RemovePastSeeds(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastSeedsKeyPrefix))
	store.Delete(types.PastSeedsKey(
		did,
	))
}

// GetAllPastSeeds returns all pastSeeds
func (k Keeper) GetAllPastSeeds(ctx sdk.Context) (list []types.PastSeeds) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PastSeedsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PastSeeds
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

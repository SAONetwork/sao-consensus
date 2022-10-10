package keeper

import (
	"github.com/SaoNetwork/sao/x/earn/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPledge set a specific pledge in the store from its index
func (k Keeper) SetPledge(ctx sdk.Context, pledge types.Pledge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	b := k.cdc.MustMarshal(&pledge)
	store.Set(types.PledgeKey(
		pledge.Creator,
	), b)
}

// GetPledge returns a pledge from its index
func (k Keeper) GetPledge(
	ctx sdk.Context,
	creator string,

) (val types.Pledge, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeKeyPrefix))

	b := store.Get(types.PledgeKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePledge removes a pledge from the store
func (k Keeper) RemovePledge(
	ctx sdk.Context,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	store.Delete(types.PledgeKey(
		creator,
	))
}

// GetAllPledge returns all pledge
func (k Keeper) GetAllPledge(ctx sdk.Context) (list []types.Pledge) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pledge
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

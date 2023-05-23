package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPledgeDebt set a specific pledgeDebt in the store from its index
func (k Keeper) SetPledgeDebt(ctx sdk.Context, pledgeDebt types.PledgeDebt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeDebtKeyPrefix))
	b := k.cdc.MustMarshal(&pledgeDebt)
	store.Set(types.PledgeDebtKey(
		pledgeDebt.Sp,
	), b)
}

// GetPledgeDebt returns a pledgeDebt from its index
func (k Keeper) GetPledgeDebt(
	ctx sdk.Context,
	sp string,

) (val types.PledgeDebt, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeDebtKeyPrefix))

	b := store.Get(types.PledgeDebtKey(
		sp,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePledgeDebt removes a pledgeDebt from the store
func (k Keeper) RemovePledgeDebt(
	ctx sdk.Context,
	sp string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeDebtKeyPrefix))
	store.Delete(types.PledgeDebtKey(
		sp,
	))
}

// GetAllPledgeDebt returns all pledgeDebt
func (k Keeper) GetAllPledgeDebt(ctx sdk.Context) (list []types.PledgeDebt) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PledgeDebtKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PledgeDebt
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCredit set a specific credit in the store from its index
func (k Keeper) SetCredit(ctx sdk.Context, credit types.Credit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreditKeyPrefix))
	b := k.cdc.MustMarshal(&credit)
	store.Set(types.CreditKey(
		credit.Account,
	), b)
}

// GetCredit returns a credit from its index
func (k Keeper) GetCredit(
	ctx sdk.Context,
	account string,

) (val types.Credit, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreditKeyPrefix))

	b := store.Get(types.CreditKey(
		account,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCredit removes a credit from the store
func (k Keeper) RemoveCredit(
	ctx sdk.Context,
	account string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreditKeyPrefix))
	store.Delete(types.CreditKey(
		account,
	))
}

// GetAllCredit returns all credit
func (k Keeper) GetAllCredit(ctx sdk.Context) (list []types.Credit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CreditKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Credit
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

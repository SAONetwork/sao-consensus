package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAccountList set a specific accountList in the store from its index
func (k Keeper) SetAccountList(ctx sdk.Context, accountList types.AccountList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountListKeyPrefix))
	b := k.cdc.MustMarshal(&accountList)
	store.Set(types.AccountListKey(
		accountList.Did,
	), b)
}

// GetAccountList returns a accountList from its index
func (k Keeper) GetAccountList(
	ctx sdk.Context,
	did string,

) (val types.AccountList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountListKeyPrefix))

	b := store.Get(types.AccountListKey(
		did,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountList removes a accountList from the store
func (k Keeper) RemoveAccountList(
	ctx sdk.Context,
	did string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountListKeyPrefix))
	store.Delete(types.AccountListKey(
		did,
	))
}

// GetAllAccountList returns all accountList
func (k Keeper) GetAllAccountList(ctx sdk.Context) (list []types.AccountList) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountListKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountList
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

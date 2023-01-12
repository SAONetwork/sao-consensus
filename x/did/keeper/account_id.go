package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAccountId set a specific accountId in the store from its index
func (k Keeper) SetAccountId(ctx sdk.Context, accountId types.AccountId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountIdKeyPrefix))
	b := k.cdc.MustMarshal(&accountId)
	store.Set(types.AccountIdKey(
		accountId.AccountDid,
	), b)
}

// GetAccountId returns a accountId from its index
func (k Keeper) GetAccountId(
	ctx sdk.Context,
	accountDid string,

) (val types.AccountId, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountIdKeyPrefix))

	b := store.Get(types.AccountIdKey(
		accountDid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountId removes a accountId from the store
func (k Keeper) RemoveAccountId(
	ctx sdk.Context,
	accountDid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountIdKeyPrefix))
	store.Delete(types.AccountIdKey(
		accountDid,
	))
}

// GetAllAccountId returns all accountId
func (k Keeper) GetAllAccountId(ctx sdk.Context) (list []types.AccountId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountIdKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

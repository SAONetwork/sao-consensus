package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetAccountAuth set a specific accountAuth in the store from its index
func (k Keeper) SetAccountAuth(ctx sdk.Context, accountAuth types.AccountAuth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAuthKeyPrefix))
	b := k.cdc.MustMarshal(&accountAuth)
	store.Set(types.AccountAuthKey(
		accountAuth.AccountDid,
	), b)
}

// GetAccountAuth returns a accountAuth from its index
func (k Keeper) GetAccountAuth(
	ctx sdk.Context,
	accountDid string,

) (val types.AccountAuth, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAuthKeyPrefix))

	b := store.Get(types.AccountAuthKey(
		accountDid,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountAuth removes a accountAuth from the store
func (k Keeper) RemoveAccountAuth(
	ctx sdk.Context,
	accountDid string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAuthKeyPrefix))
	store.Delete(types.AccountAuthKey(
		accountDid,
	))
}

// GetAllAccountAuth returns all accountAuth
func (k Keeper) GetAllAccountAuth(ctx sdk.Context) (list []types.AccountAuth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAuthKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountAuth
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

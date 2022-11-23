package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSidDocument set a specific sidDocument in the store from its index
func (k Keeper) SetSidDocument(ctx sdk.Context, sidDocument types.SidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentKeyPrefix))
	b := k.cdc.MustMarshal(&sidDocument)
	store.Set(types.SidDocumentKey(
		sidDocument.VersionId,
	), b)
}

// GetSidDocument returns a sidDocument from its index
func (k Keeper) GetSidDocument(
	ctx sdk.Context,
	versionId string,

) (val types.SidDocument, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentKeyPrefix))

	b := store.Get(types.SidDocumentKey(
		versionId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSidDocument removes a sidDocument from the store
func (k Keeper) RemoveSidDocument(
	ctx sdk.Context,
	versionId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentKeyPrefix))
	store.Delete(types.SidDocumentKey(
		versionId,
	))
}

// GetAllSidDocument returns all sidDocument
func (k Keeper) GetAllSidDocument(ctx sdk.Context) (list []types.SidDocument) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SidDocument
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

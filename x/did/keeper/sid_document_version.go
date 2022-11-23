package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSidDocumentVersion set a specific sidDocumentVersion in the store from its index
func (k Keeper) SetSidDocumentVersion(ctx sdk.Context, sidDocumentVersion types.SidDocumentVersion) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentVersionKeyPrefix))
	b := k.cdc.MustMarshal(&sidDocumentVersion)
	store.Set(types.SidDocumentVersionKey(
		sidDocumentVersion.DocId,
	), b)
}

// GetSidDocumentVersion returns a sidDocumentVersion from its index
func (k Keeper) GetSidDocumentVersion(
	ctx sdk.Context,
	docId string,

) (val types.SidDocumentVersion, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentVersionKeyPrefix))

	b := store.Get(types.SidDocumentVersionKey(
		docId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSidDocumentVersion removes a sidDocumentVersion from the store
func (k Keeper) RemoveSidDocumentVersion(
	ctx sdk.Context,
	docId string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentVersionKeyPrefix))
	store.Delete(types.SidDocumentVersionKey(
		docId,
	))
}

// GetAllSidDocumentVersion returns all sidDocumentVersion
func (k Keeper) GetAllSidDocumentVersion(ctx sdk.Context) (list []types.SidDocumentVersion) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SidDocumentVersionKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SidDocumentVersion
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

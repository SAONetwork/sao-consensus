package v2

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	return migrateOrders(ctx, store, cdc)
}

func migrateOrders(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec) error {
	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))

	iterator := sdk.KVStorePrefixIterator(orderStore, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		oldKey := iterator.Key()
		oldVal := orderStore.Get(oldKey)

		fmt.Println(oldKey, oldVal)
	}
	return nil
}

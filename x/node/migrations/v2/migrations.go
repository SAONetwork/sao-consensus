package v2

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	pledgeStore := prefix.NewStore(store, types.KeyPrefix(string(types.PledgeKeyPrefix)))

	fmt.Println(pledgeStore)

	/*
		iterator := sdk.KVStorePrefixIterator(orderStore, []byte{})

		for ; iterator.Valid(); iterator.Next() {
			orderKey := iterator.Key()
			oldVal := orderStore.Get(orderKey)
			var order types.Order
			cdc.MustUnmarshal(oldVal, &order)

			newVal := cdc.MustMarshal(&order)
			orderStore.Set(orderKey, newVal)
		}
	*/

	return nil
}

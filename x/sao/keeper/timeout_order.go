package keeper

import (
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTimeoutOrder set a specific timeoutOrder in the store from its index
func (k Keeper) SetTimeoutOrder(ctx sdk.Context, timeoutOrder types.TimeoutOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutOrderKeyPrefix))
	b := k.cdc.MustMarshal(&timeoutOrder)
	store.Set(types.TimeoutOrderKey(
		timeoutOrder.Height,
	), b)
}

// GetTimeoutOrder returns a timeoutOrder from its index
func (k Keeper) GetTimeoutOrder(
	ctx sdk.Context,
	height uint64,

) (val types.TimeoutOrder, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutOrderKeyPrefix))

	b := store.Get(types.TimeoutOrderKey(
		height,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimeoutOrder removes a timeoutOrder from the store
func (k Keeper) RemoveTimeoutOrder(
	ctx sdk.Context,
	height uint64,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutOrderKeyPrefix))
	store.Delete(types.TimeoutOrderKey(
		height,
	))
}

// GetAllTimeoutOrder returns all timeoutOrder
func (k Keeper) GetAllTimeoutOrder(ctx sdk.Context) (list []types.TimeoutOrder) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutOrderKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimeoutOrder
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) SetTimeoutOrderBlock(ctx sdk.Context, order ordertypes.Order, timeoutHeight uint64) {

	TimeoutOrder, found := k.GetTimeoutOrder(ctx, timeoutHeight)
	if found {
		TimeoutOrder.OrderList = append(TimeoutOrder.OrderList, order.Id)
	} else {
		TimeoutOrder = types.TimeoutOrder{
			Height:    timeoutHeight,
			OrderList: []uint64{order.Id},
		}
	}

	k.SetTimeoutOrder(ctx, TimeoutOrder)
}

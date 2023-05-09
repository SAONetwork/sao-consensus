package v3

import (
	v2order "github.com/SaoNetwork/sao/x/order/migrations/v2/types"
	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))

	iterator := sdk.KVStorePrefixIterator(orderStore, []byte{})

	logger := ctx.Logger()
	logger.Debug("v3 order migrations, reward per byte")
	for ; iterator.Valid(); iterator.Next() {
		orderKey := iterator.Key()
		oldVal := orderStore.Get(orderKey)
		var order v2order.Order
		cdc.MustUnmarshal(oldVal, &order)
		amount := sdk.NewDecCoinFromCoin(order.Amount)
		rewardPerByte := amount.Amount.QuoInt64(int64(order.Replica)).QuoInt64(int64(order.Duration)).QuoInt64(int64(order.Size_))
		newOrder := types.Order{
			Creator:       order.Creator,
			Owner:         order.Owner,
			Id:            order.Id,
			Provider:      order.Provider,
			Cid:           order.Cid,
			Duration:      order.Duration,
			Status:        order.Status,
			Replica:       order.Replica,
			Shards:        order.Shards,
			Amount:        order.Amount,
			Size_:         order.Size_,
			Operation:     order.Operation,
			CreatedAt:     order.CreatedAt,
			Timeout:       order.Timeout,
			DataId:        order.DataId,
			Commit:        order.Commit,
			RewardPerByte: sdk.NewDecCoinFromDec(amount.Denom, rewardPerByte),
		}

		newVal := cdc.MustMarshal(&newOrder)
		orderStore.Set(orderKey, newVal)

	}

	return nil
}

package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/order/keeper"
	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNExpiredOrder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExpiredOrder {
	items := make([]types.ExpiredOrder, n)
	for i := range items {
		items[i].Height = uint64(i)

		keeper.SetExpiredOrder(ctx, items[i])
	}
	return items
}

func TestExpiredOrderGet(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	items := createNExpiredOrder(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExpiredOrder(ctx,
			item.Height,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExpiredOrderRemove(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	items := createNExpiredOrder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExpiredOrder(ctx,
			item.Height,
		)
		_, found := keeper.GetExpiredOrder(ctx,
			item.Height,
		)
		require.False(t, found)
	}
}

func TestExpiredOrderGetAll(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	items := createNExpiredOrder(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExpiredOrder(ctx)),
	)
}

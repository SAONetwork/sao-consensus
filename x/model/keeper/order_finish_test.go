package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/model/keeper"
	"github.com/SaoNetwork/sao/x/model/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOrderFinish(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OrderFinish {
	items := make([]types.OrderFinish, n)
	for i := range items {
		items[i].Height = uint64(i)

		keeper.SetOrderFinish(ctx, items[i])
	}
	return items
}

func TestOrderFinishGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNOrderFinish(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetOrderFinish(ctx,
			item.Height,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestOrderFinishRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNOrderFinish(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOrderFinish(ctx,
			item.Height,
		)
		_, found := keeper.GetOrderFinish(ctx,
			item.Height,
		)
		require.False(t, found)
	}
}

func TestOrderFinishGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNOrderFinish(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOrderFinish(ctx)),
	)
}

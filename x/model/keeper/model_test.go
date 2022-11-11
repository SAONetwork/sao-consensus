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

func createNModel(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Model {
	items := make([]types.Model, n)
	for i := range items {
		items[i].Key = strconv.Itoa(i)

		keeper.SetModel(ctx, items[i])
	}
	return items
}

func TestModelGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModel(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetModel(ctx,
			item.Key,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestModelRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModel(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveModel(ctx,
			item.Key,
		)
		_, found := keeper.GetModel(ctx,
			item.Key,
		)
		require.False(t, found)
	}
}

func TestModelGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNModel(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllModel(ctx)),
	)
}

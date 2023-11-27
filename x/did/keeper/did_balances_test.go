package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/did/keeper"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDidBalances(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DidBalances {
	items := make([]types.DidBalances, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)

		keeper.SetDidBalances(ctx, items[i])
	}
	return items
}

func TestDidBalancesGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBalances(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDidBalances(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDidBalancesRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBalances(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidBalances(ctx,
			item.Did,
		)
		_, found := keeper.GetDidBalances(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestDidBalancesGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBalances(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDidBalances(ctx)),
	)
}

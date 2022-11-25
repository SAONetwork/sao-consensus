package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/market/keeper"
	"github.com/SaoNetwork/sao/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWorker(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Worker {
	items := make([]types.Worker, n)
	for i := range items {
		items[i].Workername = strconv.Itoa(i)

		keeper.SetWorker(ctx, items[i])
	}
	return items
}

func TestWorkerGet(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNWorker(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWorker(ctx,
			item.Workername,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWorkerRemove(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNWorker(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWorker(ctx,
			item.Workername,
		)
		_, found := keeper.GetWorker(ctx,
			item.Workername,
		)
		require.False(t, found)
	}
}

func TestWorkerGetAll(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	items := createNWorker(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWorker(ctx)),
	)
}

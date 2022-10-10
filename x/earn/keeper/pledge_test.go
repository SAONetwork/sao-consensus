package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/earn/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPledge(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Pledge {
	items := make([]types.Pledge, n)
	for i := range items {
		items[i].Creator = strconv.Itoa(i)

		keeper.SetPledge(ctx, items[i])
	}
	return items
}

func TestPledgeGet(t *testing.T) {
	keeper, ctx := keepertest.EarnKeeper(t)
	items := createNPledge(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPledge(ctx,
			item.Creator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPledgeRemove(t *testing.T) {
	keeper, ctx := keepertest.EarnKeeper(t)
	items := createNPledge(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePledge(ctx,
			item.Creator,
		)
		_, found := keeper.GetPledge(ctx,
			item.Creator,
		)
		require.False(t, found)
	}
}

func TestPledgeGetAll(t *testing.T) {
	keeper, ctx := keepertest.EarnKeeper(t)
	items := createNPledge(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPledge(ctx)),
	)
}

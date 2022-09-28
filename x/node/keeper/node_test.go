package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNode(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Node {
	items := make([]types.Node, n)
	for i := range items {
		items[i].Creator = strconv.Itoa(i)

		keeper.SetNode(ctx, items[i])
	}
	return items
}

func TestNodeGet(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNNode(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNode(ctx,
			item.Creator,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNodeRemove(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNNode(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNode(ctx,
			item.Creator,
		)
		_, found := keeper.GetNode(ctx,
			item.Creator,
		)
		require.False(t, found)
	}
}

func TestNodeGetAll(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNNode(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNode(ctx)),
	)
}

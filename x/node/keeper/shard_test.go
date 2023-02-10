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

func createNShard(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Shard {
	items := make([]types.Shard, n)
	for i := range items {
		items[i].Idx = strconv.Itoa(i)

		keeper.SetShard(ctx, items[i])
	}
	return items
}

func TestShardGet(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNShard(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetShard(ctx,
			item.Idx,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestShardRemove(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNShard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShard(ctx,
			item.Idx,
		)
		_, found := keeper.GetShard(ctx,
			item.Idx,
		)
		require.False(t, found)
	}
}

func TestShardGetAll(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNShard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShard(ctx)),
	)
}

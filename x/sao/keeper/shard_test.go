package keeper_test

import (
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/sao/keeper"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNShard(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Shard {
	items := make([]types.Shard, n)
	for i := range items {
		items[i].Id = keeper.AppendShard(ctx, items[i])
	}
	return items
}

func TestShardGet(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNShard(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetShard(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestShardRemove(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNShard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShard(ctx, item.Id)
		_, found := keeper.GetShard(ctx, item.Id)
		require.False(t, found)
	}
}

func TestShardGetAll(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNShard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShard(ctx)),
	)
}

func TestShardCount(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNShard(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetShardCount(ctx))
}

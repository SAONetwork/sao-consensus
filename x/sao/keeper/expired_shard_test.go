package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/sao/keeper"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNExpiredShard(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExpiredShard {
	items := make([]types.ExpiredShard, n)
	for i := range items {
		items[i].Height = uint64(i)

		keeper.SetExpiredShard(ctx, items[i])
	}
	return items
}

func TestExpiredShardGet(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNExpiredShard(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExpiredShard(ctx,
			item.Height,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExpiredShardRemove(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNExpiredShard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExpiredShard(ctx,
			item.Height,
		)
		_, found := keeper.GetExpiredShard(ctx,
			item.Height,
		)
		require.False(t, found)
	}
}

func TestExpiredShardGetAll(t *testing.T) {
	keeper, ctx := keepertest.SaoKeeper(t)
	items := createNExpiredShard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExpiredShard(ctx)),
	)
}

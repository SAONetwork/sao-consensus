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

func createNExpiredData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExpiredData {
	items := make([]types.ExpiredData, n)
	for i := range items {
		items[i].Height = uint64(i)

		keeper.SetExpiredData(ctx, items[i])
	}
	return items
}

func TestExpiredDataGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNExpiredData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExpiredData(ctx,
			item.Height,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExpiredDataRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNExpiredData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExpiredData(ctx,
			item.Height,
		)
		_, found := keeper.GetExpiredData(ctx,
			item.Height,
		)
		require.False(t, found)
	}
}

func TestExpiredDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNExpiredData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExpiredData(ctx)),
	)
}

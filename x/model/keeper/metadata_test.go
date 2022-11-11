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

func createNMetadata(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Metadata {
	items := make([]types.Metadata, n)
	for i := range items {
		items[i].DataId = strconv.Itoa(i)

		keeper.SetMetadata(ctx, items[i])
	}
	return items
}

func TestMetadataGet(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNMetadata(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetMetadata(ctx,
			item.DataId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestMetadataRemove(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNMetadata(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMetadata(ctx,
			item.DataId,
		)
		_, found := keeper.GetMetadata(ctx,
			item.DataId,
		)
		require.False(t, found)
	}
}

func TestMetadataGetAll(t *testing.T) {
	keeper, ctx := keepertest.ModelKeeper(t)
	items := createNMetadata(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMetadata(ctx)),
	)
}

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

func createNSidDocument(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SidDocument {
	items := make([]types.SidDocument, n)
	for i := range items {
		items[i].VersionId = strconv.Itoa(i)

		keeper.SetSidDocument(ctx, items[i])
	}
	return items
}

func TestSidDocumentGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNSidDocument(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSidDocument(ctx,
			item.VersionId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSidDocumentRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNSidDocument(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSidDocument(ctx,
			item.VersionId,
		)
		_, found := keeper.GetSidDocument(ctx,
			item.VersionId,
		)
		require.False(t, found)
	}
}

func TestSidDocumentGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNSidDocument(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSidDocument(ctx)),
	)
}

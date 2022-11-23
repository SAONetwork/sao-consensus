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

func createNSidDocumentVersion(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SidDocumentVersion {
	items := make([]types.SidDocumentVersion, n)
	for i := range items {
		items[i].DocId = strconv.Itoa(i)

		keeper.SetSidDocumentVersion(ctx, items[i])
	}
	return items
}

func TestSidDocumentVersionGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNSidDocumentVersion(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSidDocumentVersion(ctx,
			item.DocId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSidDocumentVersionRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNSidDocumentVersion(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSidDocumentVersion(ctx,
			item.DocId,
		)
		_, found := keeper.GetSidDocumentVersion(ctx,
			item.DocId,
		)
		require.False(t, found)
	}
}

func TestSidDocumentVersionGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNSidDocumentVersion(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSidDocumentVersion(ctx)),
	)
}

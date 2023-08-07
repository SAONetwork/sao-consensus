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

func createNKid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Kid {
	items := make([]types.Kid, n)
	for i := range items {
		items[i].AccountId = strconv.Itoa(i)

		keeper.SetKid(ctx, items[i])
	}
	return items
}

func TestKidGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNKid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKid(ctx,
			item.AccountId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKidRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNKid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKid(ctx,
			item.AccountId,
		)
		_, found := keeper.GetKid(ctx,
			item.AccountId,
		)
		require.False(t, found)
	}
}

func TestKidGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNKid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKid(ctx)),
	)
}

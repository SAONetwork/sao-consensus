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

func createNDid(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Did {
	items := make([]types.Did, n)
	for i := range items {
		items[i].AccountId = strconv.Itoa(i)

		keeper.SetDid(ctx, items[i])
	}
	return items
}

func TestDidGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDid(ctx,
			item.AccountId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDidRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDid(ctx,
			item.AccountId,
		)
		_, found := keeper.GetDid(ctx,
			item.AccountId,
		)
		require.False(t, found)
	}
}

func TestDidGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDid(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDid(ctx)),
	)
}

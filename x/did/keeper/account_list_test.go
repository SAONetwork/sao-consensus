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

func createNAccountList(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccountList {
	items := make([]types.AccountList, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)

		keeper.SetAccountList(ctx, items[i])
	}
	return items
}

func TestAccountListGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountList(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccountList(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountListRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountList(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccountList(ctx,
			item.Did,
		)
		_, found := keeper.GetAccountList(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestAccountListGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountList(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccountList(ctx)),
	)
}

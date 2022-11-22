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

func createNAccountAuth(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccountAuth {
	items := make([]types.AccountAuth, n)
	for i := range items {
		items[i].AccountDid = strconv.Itoa(i)

		keeper.SetAccountAuth(ctx, items[i])
	}
	return items
}

func TestAccountAuthGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountAuth(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccountAuth(ctx,
			item.AccountDid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountAuthRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountAuth(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccountAuth(ctx,
			item.AccountDid,
		)
		_, found := keeper.GetAccountAuth(ctx,
			item.AccountDid,
		)
		require.False(t, found)
	}
}

func TestAccountAuthGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountAuth(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccountAuth(ctx)),
	)
}

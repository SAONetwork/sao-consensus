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

func createNAccountId(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.AccountId {
	items := make([]types.AccountId, n)
	for i := range items {
		items[i].AccountDid = strconv.Itoa(i)

		keeper.SetAccountId(ctx, items[i])
	}
	return items
}

func TestAccountIdGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccountId(ctx,
			item.AccountDid,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountIdRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountId(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccountId(ctx,
			item.AccountDid,
		)
		_, found := keeper.GetAccountId(ctx,
			item.AccountDid,
		)
		require.False(t, found)
	}
}

func TestAccountIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNAccountId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccountId(ctx)),
	)
}

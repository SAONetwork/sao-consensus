package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/loan/keeper"
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNCredit(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Credit {
	items := make([]types.Credit, n)
	for i := range items {
		items[i].Account = strconv.Itoa(i)

		keeper.SetCredit(ctx, items[i])
	}
	return items
}

func TestCreditGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNCredit(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetCredit(ctx,
			item.Account,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestCreditRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNCredit(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCredit(ctx,
			item.Account,
		)
		_, found := keeper.GetCredit(ctx,
			item.Account,
		)
		require.False(t, found)
	}
}

func TestCreditGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNCredit(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCredit(ctx)),
	)
}

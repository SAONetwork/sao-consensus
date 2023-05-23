package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPledgeDebt(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PledgeDebt {
	items := make([]types.PledgeDebt, n)
	for i := range items {
		items[i].Sp = strconv.Itoa(i)

		keeper.SetPledgeDebt(ctx, items[i])
	}
	return items
}

func TestPledgeDebtGet(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNPledgeDebt(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPledgeDebt(ctx,
			item.Sp,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPledgeDebtRemove(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNPledgeDebt(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePledgeDebt(ctx,
			item.Sp,
		)
		_, found := keeper.GetPledgeDebt(ctx,
			item.Sp,
		)
		require.False(t, found)
	}
}

func TestPledgeDebtGetAll(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	items := createNPledgeDebt(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPledgeDebt(ctx)),
	)
}

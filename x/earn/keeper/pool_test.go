package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/earn/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
)

func createTestPool(keeper *keeper.Keeper, ctx sdk.Context) types.Pool {
	item := types.Pool{}
	keeper.SetPool(ctx, item)
	return item
}

func TestPoolGet(t *testing.T) {
	keeper, ctx := keepertest.EarnKeeper(t)
	item := createTestPool(keeper, ctx)
	rst, found := keeper.GetPool(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.EarnKeeper(t)
	createTestPool(keeper, ctx)
	keeper.RemovePool(ctx)
	_, found := keeper.GetPool(ctx)
	require.False(t, found)
}

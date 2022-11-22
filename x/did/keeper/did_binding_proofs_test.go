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

func createNDidBindingProofs(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DidBindingProofs {
	items := make([]types.DidBindingProofs, n)
	for i := range items {
		items[i].AccountId = strconv.Itoa(i)

		keeper.SetDidBindingProofs(ctx, items[i])
	}
	return items
}

func TestDidBindingProofsGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBindingProofs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDidBindingProofs(ctx,
			item.AccountId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDidBindingProofsRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBindingProofs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidBindingProofs(ctx,
			item.AccountId,
		)
		_, found := keeper.GetDidBindingProofs(ctx,
			item.AccountId,
		)
		require.False(t, found)
	}
}

func TestDidBindingProofsGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBindingProofs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDidBindingProofs(ctx)),
	)
}

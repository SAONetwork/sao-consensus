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

func createNDidBindingProof(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DidBindingProof {
	items := make([]types.DidBindingProof, n)
	for i := range items {
		items[i].AccountId = strconv.Itoa(i)

		keeper.SetDidBindingProof(ctx, items[i])
	}
	return items
}

func TestDidBindingProofGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBindingProof(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDidBindingProof(ctx,
			item.AccountId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDidBindingProofRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBindingProof(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidBindingProof(ctx,
			item.AccountId,
		)
		_, found := keeper.GetDidBindingProof(ctx,
			item.AccountId,
		)
		require.False(t, found)
	}
}

func TestDidBindingProofGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBindingProof(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDidBindingProof(ctx)),
	)
}

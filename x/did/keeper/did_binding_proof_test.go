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

func createNDidBingingProof(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.DidBingingProof {
	items := make([]types.DidBingingProof, n)
	for i := range items {
		items[i].AccountId = strconv.Itoa(i)

		keeper.SetDidBingingProof(ctx, items[i])
	}
	return items
}

func TestDidBingingProofGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBingingProof(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDidBingingProof(ctx,
			item.AccountId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDidBingingProofRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBingingProof(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDidBingingProof(ctx,
			item.AccountId,
		)
		_, found := keeper.GetDidBingingProof(ctx,
			item.AccountId,
		)
		require.False(t, found)
	}
}

func TestDidBingingProofGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNDidBingingProof(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDidBingingProof(ctx)),
	)
}

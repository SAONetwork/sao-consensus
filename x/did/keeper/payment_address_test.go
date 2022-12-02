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

func createNPaymentAddress(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PaymentAddress {
	items := make([]types.PaymentAddress, n)
	for i := range items {
		items[i].Did = strconv.Itoa(i)

		keeper.SetPaymentAddress(ctx, items[i])
	}
	return items
}

func TestPaymentAddressGet(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNPaymentAddress(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPaymentAddress(ctx,
			item.Did,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPaymentAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNPaymentAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePaymentAddress(ctx,
			item.Did,
		)
		_, found := keeper.GetPaymentAddress(ctx,
			item.Did,
		)
		require.False(t, found)
	}
}

func TestPaymentAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.DidKeeper(t)
	items := createNPaymentAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPaymentAddress(ctx)),
	)
}

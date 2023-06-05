package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/loan/keeper"
	"github.com/SaoNetwork/sao/x/loan/types"
)

func createTestLoanPool(keeper *keeper.Keeper, ctx sdk.Context) types.LoanPool {
	item := types.LoanPool{
		Total:              sdk.DecCoin{"", sdk.NewDec(0)},
		LoanedOut:          sdk.Coin{"", sdk.NewInt(0)},
		TotalBonds:         sdk.NewDec(0),
		InterestDebt:       sdk.DecCoin{"", sdk.NewDec(0)},
		AccInterestPerCoin: sdk.DecCoin{"", sdk.NewDec(0)},
	}
	keeper.SetLoanPool(ctx, item)
	return item
}

func TestLoanPoolGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	item := createTestLoanPool(keeper, ctx)
	rst, found := keeper.GetLoanPool(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLoanPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	createTestLoanPool(keeper, ctx)
	keeper.RemoveLoanPool(ctx)
	_, found := keeper.GetLoanPool(ctx)
	require.False(t, found)
}

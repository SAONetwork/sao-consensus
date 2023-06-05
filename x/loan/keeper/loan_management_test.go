package keeper_test

import (
	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/loan"
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeposit(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	denom := "sao"

	loanPool, found := keeper.GetLoanPool(ctx)
	require.False(t, found)
	amount := sdk.NewCoin(denom, sdk.NewInt(100))
	decAmount := sdk.NewDecCoinFromCoin(amount)
	account1 := "account1"
	err := keeper.Deposit(ctx, account1, decAmount)
	require.NoError(t, err)

	loanPool, found = keeper.GetLoanPool(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&loanPool),
		nullify.Fill(&types.LoanPool{
			Total:              decAmount,
			LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(0)),
			TotalBonds:         decAmount.Amount,
			InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(0)),
			AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(0)),
		}),
	)
	credit1, found := keeper.GetCredit(ctx, account1)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&credit1),
		nullify.Fill(&types.Credit{
			Account: account1,
			Bonds:   decAmount.Amount,
		}),
	)
	loanPool.LoanedOut = sdk.NewCoin(denom, sdk.NewInt(50))
	keeper.SetLoanPool(ctx, loanPool)
	keeper.SetParams(ctx, types.Params{
		InterestRatePerBlock: "1",
		MinLiquidityRatio:    "0.3",
	})

	ctx = ctx.WithBlockHeight(2)
	keeper.SetLoanPool(ctx, loanPool)
	loan.BeginBlocker(ctx, *keeper)
	loan.BeginBlocker(ctx, *keeper)

	err = keeper.Deposit(ctx, account1, decAmount)
	require.NoError(t, err)

	loanPool, found = keeper.GetLoanPool(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&loanPool),
		nullify.Fill(&types.LoanPool{
			Total:              sdk.NewDecCoin(denom, sdk.NewInt(300)),
			LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(50)),
			TotalBonds:         sdk.NewDec(150),
			InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(100)),
			AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(2)),
		}),
	)
	credit1, found = keeper.GetCredit(ctx, account1)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&credit1),
		nullify.Fill(&types.Credit{
			Account: account1,
			Bonds:   sdk.NewDec(150),
		}),
	)
}

func TestWithdraw(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	denom := "sao"

	account1 := "account1"

	keeper.SetLoanPool(ctx, types.LoanPool{
		Total:              sdk.DecCoin{denom, sdk.NewDec(200)},
		LoanedOut:          sdk.Coin{denom, sdk.NewInt(50)},
		TotalBonds:         sdk.NewDec(100),
		InterestDebt:       sdk.DecCoin{denom, sdk.NewDec(0)},
		AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(0)),
	})

	keeper.SetCredit(ctx, types.Credit{
		Account: account1,
		Bonds:   sdk.NewDec(100),
	})

	ctx = ctx.WithBlockHeight(2)
	keeper.SetParams(ctx, types.Params{
		InterestRatePerBlock: "1",
		MinLiquidityRatio:    "0.3",
	})
	loan.BeginBlocker(ctx, *keeper)
	loan.BeginBlocker(ctx, *keeper)

	err := keeper.Withdraw(ctx, account1, sdk.DecCoin{denom, sdk.NewDec(150)})
	require.NoError(t, err)

	loanPool, found := keeper.GetLoanPool(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&loanPool),
		nullify.Fill(&types.LoanPool{
			Total:              sdk.NewDecCoin(denom, sdk.NewInt(150)),
			LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(50)),
			TotalBonds:         sdk.NewDec(50),
			InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(100)),
			AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(2)),
		}),
	)
	credit1, found := keeper.GetCredit(ctx, account1)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&credit1),
		nullify.Fill(&types.Credit{
			Account: account1,
			Bonds:   sdk.NewDec(50),
		}),
	)
}

func TestLoanOut(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	denom := "sao"

	loanPool := types.LoanPool{
		Total:              sdk.NewDecCoin(denom, sdk.NewInt(500)),
		LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(200)),
		TotalBonds:         sdk.NewDec(250),
		InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(0)),
		AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(0)),
	}
	keeper.SetLoanPool(ctx, loanPool)

	loanedOut, err := keeper.LoanOut(ctx, sdk.NewCoin(denom, sdk.NewInt(200)))
	require.NoError(t, err)
	require.Equal(t, loanedOut, sdk.NewCoin(denom, sdk.NewInt(150)))

	loanPool, found := keeper.GetLoanPool(ctx)
	require.True(t, found)
	require.Equal(t, loanPool, types.LoanPool{
		Total:              sdk.NewDecCoin(denom, sdk.NewInt(500)),
		LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(350)),
		TotalBonds:         sdk.NewDec(250),
		InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(0)),
		AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(0)),
	})
}

func TestRepay(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	denom := "sao"

	loanPool := types.LoanPool{
		Total:              sdk.NewDecCoin(denom, sdk.NewInt(500)),
		LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(200)),
		TotalBonds:         sdk.NewDec(250),
		InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(0)),
		AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(0)),
	}
	keeper.SetLoanPool(ctx, loanPool)

	err := keeper.Repay(ctx, sdk.NewCoin(denom, sdk.NewInt(200)))
	require.NoError(t, err)

	loanPool, found := keeper.GetLoanPool(ctx)
	require.True(t, found)
	require.Equal(t, loanPool, types.LoanPool{
		Total:              sdk.NewDecCoin(denom, sdk.NewInt(500)),
		LoanedOut:          sdk.NewCoin(denom, sdk.NewInt(0)),
		TotalBonds:         sdk.NewDec(250),
		InterestDebt:       sdk.NewDecCoin(denom, sdk.NewInt(0)),
		AccInterestPerCoin: sdk.NewDecCoin(denom, sdk.NewInt(0)),
	})
}

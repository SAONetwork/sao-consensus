package loan

import (
	"github.com/SaoNetwork/sao/x/loan/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	loanPool, found := k.GetLoanPool(ctx)
	if found && !loanPool.LoanedOut.IsZero() {
		params := k.GetParams(ctx)
		interestRate, _ := sdk.NewDecFromStr(params.InterestRatePerBlock)
		loanPool.AccInterestPerCoin.Amount = loanPool.AccInterestPerCoin.Amount.Add(interestRate)
		k.SetLoanPool(ctx, loanPool)
	}
}

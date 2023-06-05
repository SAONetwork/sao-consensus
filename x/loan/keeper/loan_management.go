package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Deposit(ctx sdk.Context, account string, amount sdk.DecCoin) error {

	var credit types.Credit
	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		loanPool = types.LoanPool{
			Total:              amount,
			LoanedOut:          sdk.NewCoin(amount.Denom, sdk.NewInt(0)),
			TotalBonds:         amount.Amount,
			InterestDebt:       sdk.NewDecCoin(amount.Denom, sdk.NewInt(0)),
			AccInterestPerCoin: sdk.NewDecCoin(amount.Denom, sdk.NewInt(0)),
		}

		credit = types.Credit{
			Account: account,
			Bonds:   amount.Amount,
		}

	} else {
		err := k.ChargeInterest(ctx, &loanPool)
		bonds := amount.Amount.Mul(loanPool.TotalBonds).Quo(loanPool.Total.Amount)
		if err != nil {
			return err
		}

		loanPool.TotalBonds = loanPool.TotalBonds.Add(bonds)
		loanPool.Total = loanPool.Total.Add(amount)

		credit, found = k.GetCredit(ctx, account)
		if !found {
			credit = types.Credit{
				Account: account,
				Bonds:   amount.Amount,
			}
		} else {
			credit.Bonds = credit.Bonds.Add(bonds)
		}
	}

	if !loanPool.LoanedOut.IsZero() {
		loanPool.InterestDebt.Amount = loanPool.AccInterestPerCoin.Amount.MulInt(loanPool.LoanedOut.Amount)
	}
	k.SetLoanPool(ctx, loanPool)
	k.SetCredit(ctx, credit)

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, account string, amount sdk.DecCoin) error {

	credit, found := k.GetCredit(ctx, account)
	if !found {
		return types.ErrCreditNotFound
	}

	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		return types.ErrLoanPoolNotFound
	}

	err := k.ChargeInterest(ctx, &loanPool)
	if err != nil {
		return err
	}

	bonds := amount.Amount.Mul(loanPool.TotalBonds).Quo(loanPool.Total.Amount)

	if bonds.GT(credit.Bonds) {
		return types.ErrInvalidAmount
	}

	loanPool.TotalBonds = loanPool.TotalBonds.Sub(bonds)
	loanPool.Total = loanPool.Total.Sub(amount)

	credit.Bonds = credit.Bonds.Sub(bonds)

	if !loanPool.LoanedOut.IsZero() {
		loanPool.InterestDebt.Amount = loanPool.AccInterestPerCoin.Amount.MulInt(loanPool.LoanedOut.Amount)
	}
	k.SetLoanPool(ctx, loanPool)
	k.SetCredit(ctx, credit)

	return nil
}

func (k Keeper) ChargeInterest(ctx sdk.Context, loanPool *types.LoanPool) error {
	if !loanPool.LoanedOut.IsZero() {
		loanedOut := sdk.NewDecCoinFromCoin(loanPool.LoanedOut)
		interest := loanPool.AccInterestPerCoin.Amount.Mul(loanedOut.Amount)
		if interest.GT(loanPool.InterestDebt.Amount) {
			loanPool.Total.Amount = loanPool.Total.Amount.Add(interest.Sub(loanPool.InterestDebt.Amount))
		}
	}

	return nil
}

func (k Keeper) LoanOut(ctx sdk.Context, amount sdk.Coin) (sdk.Coin, error) {

	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		return sdk.NewCoin(amount.Denom, sdk.NewInt(0)), types.ErrLoanPoolNotFound
	}

	loanable, err := k.GetLoanable(ctx, loanPool)
	if err != nil {
		return sdk.Coin{}, err
	}

	err = k.ChargeInterest(ctx, &loanPool)
	if err != nil {
		return sdk.Coin{}, err
	}

	var loanedOut sdk.Coin
	if loanable.IsZero() {
		return loanable, err
	} else if loanable.IsLT(amount) {
		loanPool.LoanedOut = loanPool.LoanedOut.Add(loanable)
		loanedOut = loanable
	} else {
		loanPool.LoanedOut = loanPool.LoanedOut.Add(amount)
		loanedOut = amount
	}
	loanPool.InterestDebt.Amount = loanPool.AccInterestPerCoin.Amount.MulInt(loanPool.LoanedOut.Amount)
	k.SetLoanPool(ctx, loanPool)
	return loanedOut, nil
}

func (k Keeper) Repay(ctx sdk.Context, amount sdk.Coin) error {

	if amount.IsZero() {
		return nil
	}

	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		return types.ErrLoanPoolNotFound
	}

	if loanPool.LoanedOut.IsLT(amount) {
		return types.ErrInvalidAmount
	}

	err := k.ChargeInterest(ctx, &loanPool)
	if err != nil {
		return err
	}

	loanPool.LoanedOut = loanPool.LoanedOut.Sub(amount)
	loanPool.InterestDebt.Amount = loanPool.AccInterestPerCoin.Amount.MulInt(loanPool.LoanedOut.Amount)
	k.SetLoanPool(ctx, loanPool)
	return nil
}

func (k Keeper) GetLoanable(ctx sdk.Context, loanPool types.LoanPool) (sdk.Coin, error) {

	params := k.GetParams(ctx)
	minLiquidityRatio, err := sdk.NewDecFromStr(params.MinLiquidityRatio)
	if err != nil {
		return sdk.Coin{}, err
	}

	// loanable = total - total * minLiquidityRatio - loanedOut
	totalLoanable := loanPool.Total.Amount.Sub(loanPool.Total.Amount.Mul(minLiquidityRatio))
	decLoanedOut := sdk.NewDecCoinFromCoin(loanPool.LoanedOut)
	if totalLoanable.LT(decLoanedOut.Amount) {
		return sdk.NewCoin(loanPool.Total.Denom, sdk.NewInt(0)), nil
	} else {
		loanable, _ := sdk.NewDecCoinFromDec(loanPool.Total.Denom, totalLoanable.Sub(decLoanedOut.Amount)).TruncateDecimal()
		return loanable, nil
	}
}

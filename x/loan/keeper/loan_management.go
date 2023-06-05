package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Deposit(ctx sdk.Context, account string, amount sdk.DecCoin) error {

	currentHeight := uint64(ctx.BlockHeight())

	var credit types.Credit
	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		loanPool = types.LoanPool{
			Total:         amount,
			LoanedOut:     sdk.NewCoin(amount.Denom, sdk.NewInt(0)),
			TotalBonds:    amount.Amount,
			LastChargedAt: currentHeight,
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

	k.SetLoanPool(ctx, loanPool)
	k.SetCredit(ctx, credit)

	return nil
}

func (k Keeper) ChargeInterest(ctx sdk.Context, loanPool *types.LoanPool) error {

	currentHeight := uint64(ctx.BlockHeight())

	params := k.GetParams(ctx)
	interestRatePerBlock, err := sdk.NewDecFromStr(params.InterestRatePerBlock)
	if err != nil {
		return err
	}

	if !loanPool.LoanedOut.IsZero() && currentHeight > loanPool.LastChargedAt {
		loanedOut := sdk.NewDecCoinFromCoin(loanPool.LoanedOut)
		interest := loanedOut.Amount.MulInt64(int64(currentHeight - loanPool.LastChargedAt)).Mul(interestRatePerBlock)
		loanPool.Total.Amount = loanPool.Total.Amount.Add(interest)
	}
	loanPool.LastChargedAt = currentHeight

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

	if loanable.IsZero() {
		return loanable, err
	} else if loanable.IsLT(amount) {
		loanPool.LoanedOut = loanPool.LoanedOut.Add(loanable)
		k.SetLoanPool(ctx, loanPool)
		return loanable, err
	} else {
		loanPool.LoanedOut = loanPool.LoanedOut.Add(amount)
		k.SetLoanPool(ctx, loanPool)
		return amount, err
	}
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

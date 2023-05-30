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
		bonds := amount.Amount.Mul(loanPool.TotalBonds).Quo(loanPool.Total.Amount)

		err := k.ChargeInterest(ctx, &loanPool)
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

	if bonds.LT(credit.Bonds) {
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

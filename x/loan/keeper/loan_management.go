package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const SaoLoanTokenDenom = "slt"
const InitialSaoLoanTokenConversionRatio int64 = 1000000000

func (k Keeper) Deposit(ctx sdk.Context, account string, amount sdk.DecCoin) (sdk.Coin, error) {
	var loanToken sdk.Coin
	loanPool, found := k.GetLoanPool(ctx)
	if !found || loanPool.Total.IsZero() {

		loanToken = sdk.NewCoin(SaoLoanTokenDenom, amount.Amount.MulInt64(InitialSaoLoanTokenConversionRatio).TruncateInt())

		err := k.bank.MintCoins(ctx, types.ModuleName, sdk.NewCoins(loanToken))
		if err != nil {
			return sdk.Coin{}, err
		}

		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(account), sdk.Coins{loanToken})
		if err != nil {
			return sdk.Coin{}, err
		}

		loanPool = types.LoanPool{
			Total:              amount,
			LoanedOut:          sdk.NewCoin(amount.Denom, sdk.NewInt(0)),
			AccInterestPerCoin: sdk.NewDecCoin(amount.Denom, sdk.NewInt(0)),
		}

	} else {
		totalSlt := k.bank.GetSupply(ctx, SaoLoanTokenDenom)
		loanToken = sdk.NewCoin(SaoLoanTokenDenom, amount.Amount.MulInt(totalSlt.Amount).Quo(loanPool.Total.Amount).TruncateInt())

		err := k.bank.MintCoins(ctx, types.ModuleName, sdk.NewCoins(loanToken))
		if err != nil {
			return sdk.Coin{}, err
		}

		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(account), sdk.Coins{loanToken})
		if err != nil {
			return sdk.Coin{}, err
		}

		loanPool.Total = loanPool.Total.Add(amount)
	}

	k.SetLoanPool(ctx, loanPool)

	return loanToken, nil
}

func (k Keeper) Withdraw(ctx sdk.Context, account string, amount sdk.DecCoin) (sdk.Coin, error) {

	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		return sdk.Coin{}, types.ErrLoanPoolNotFound
	}

	available := loanPool.Total.Amount.TruncateInt().Sub(loanPool.LoanedOut.Amount)
	if available.LT(amount.Amount.TruncateInt()) {
		return sdk.Coin{}, types.ErrNoEnoughAvailable
	}

	totalSlt := k.bank.GetSupply(ctx, SaoLoanTokenDenom)
	loanToken := sdk.NewCoin(SaoLoanTokenDenom, amount.Amount.MulInt(totalSlt.Amount).Quo(loanPool.Total.Amount).Ceil().TruncateInt())

	accToken := k.bank.GetBalance(ctx, sdk.MustAccAddressFromBech32(account), SaoLoanTokenDenom)

	if accToken.IsLT(loanToken) {
		return sdk.Coin{}, types.ErrInvalidAmount
	}

	loanPool.Total = loanPool.Total.Sub(amount)

	k.SetLoanPool(ctx, loanPool)

	return loanToken, nil
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

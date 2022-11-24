package keeper

import (
	"github.com/SaoNetwork/sao/x/market/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/status"
)

func (k Keeper) Deposit(ctx sdk.Context, order ordertypes.Order) error {

	amount := order.Amount
	duration := int64(order.Duration)

	if amount.IsZero() {
		return sdkerrors.Wrap(types.ErrInvalidAmount, "")
	}

	pool, found := k.GetPool(ctx, amount.Denom)
	if !found {
		pool = types.Pool{
			Index:           amount.Denom,
			TotalBalance:    sdk.NewInt64Coin(amount.Denom, 0),
			TotalPaid:       sdk.NewInt64Coin(amount.Denom, 0),
			IncomePerSecond: sdk.NewInt64DecCoin(amount.Denom, 0),
		}
	}

	err := k.bank.SendCoinsFromModuleToModule(ctx, ordertypes.ModuleName, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}

	income := sdk.NewDecCoinFromCoin(amount)

	pool.IncomePerSecond.Amount = pool.IncomePerSecond.Amount.Add(income.Amount.QuoInt64(duration))

	pool.TotalBalance = pool.TotalBalance.Add(amount)

	k.SetPool(ctx, pool)

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, order ordertypes.Order) error {

	amount := order.Amount
	duration := int64(order.Duration)

	if amount.IsZero() {
		return sdkerrors.Wrap(types.ErrInvalidAmount, "")
	}

	pool, found := k.GetPool(ctx, amount.Denom)
	if !found {
		return status.Error(codes.NotFound, "pool %d not found", amount.Denom)
	}

	income := sdk.NewDecCoinFromCoin(amount)

	incomePerSecond := income.Amount.QuoInt64(duration)

	incomePerSecond.MulInt64(ctx.BlockTime().Unix() - int64(order.CreatedAt)).TruncateInt()

	err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, ordertypes.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}

	pool.IncomePerSecond.Amount = pool.IncomePerSecond.Amount.Sub(incomePerSecond)

	pool.TotalBalance = pool.TotalBalance.Sub(amount)

	k.SetPool(ctx, pool)

	return nil
}

func (k Keeper) AddSpace(ctx sdk.Context, size uint64) {

}

func (k Keeper) RemoveSpace(ctx sdk.Context, size uint64) {

}

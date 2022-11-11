package keeper

import (
	"math/big"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error {

	logger := k.Logger(ctx)

	pledge, found_pledge := k.GetPledge(ctx, sp.String())

	if !found_pledge {
		pledge = types.Pledge{
			Creator: sp.String(),
			Pledged: sdk.NewInt64Coin(amount.Denom, 0),
			Reward:  sdk.NewInt64Coin(amount.Denom, 0),
		}
	}

	pool, found_pool := k.GetPool(ctx)

	if !found_pool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	if amount.Denom != pool.Denom.Denom {
		return sdkerrors.Wrapf(types.ErrDenom, "want %s but %s", pool.Denom.Denom, amount.Denom)
	}

	pool.Denom = pool.Denom.Add(amount)

	logger.Debug("pool denom", "amount", pool.Denom)

	coinPerShare, _ := new(big.Int).SetString(pool.CoinPerShare, 10)

	if found_pledge && pledge.Pledged.Amount.Int64() > 0 {
		pending := new(big.Int).Sub(new(big.Int).Div(new(big.Int).Mul(pledge.Pledged.Amount.BigInt(), coinPerShare), big.NewInt(1e12)), pledge.RewardDebt.Amount.BigInt())
		pledge.Reward = pledge.Reward.AddAmount(sdk.NewInt(pending.Int64()))
	}

	err := k.bank.SendCoinsFromAccountToModule(ctx, sp, types.ModuleName, sdk.NewCoins(amount))

	if err != nil {
		return err
	}

	pledge.Pledged = pledge.Pledged.Add(amount)

	rewardDebt := new(big.Int).Div(new(big.Int).Mul(pledge.Pledged.Amount.BigInt(), coinPerShare), big.NewInt(1e12))

	pledge.RewardDebt = sdk.NewInt64Coin(pool.Denom.Denom, rewardDebt.Int64())

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return nil
}

func (k Keeper) OrderRelease(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error {
	pledge, found := k.GetPledge(ctx, sp.String())

	if !found {
		return sdkerrors.Wrapf(types.ErrPledgeNotFound, "")
	}

	pool, found := k.GetPool(ctx)

	if !found {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	if amount.Denom != pool.Denom.Denom {
		return sdkerrors.Wrapf(types.ErrDenom, "want %s but %s", pool.Denom.Denom, amount.Denom)
	}

	if pool.Denom.IsLT(amount) {
		return sdkerrors.Wrap(types.ErrInsufficientCoin, "")
	}

	pool.Denom.Sub(amount)

	coinPerShare, _ := new(big.Int).SetString(pool.CoinPerShare, 10)

	pending := new(big.Int).Sub(new(big.Int).Div(new(big.Int).Mul(pledge.Pledged.Amount.BigInt(), coinPerShare), big.NewInt(1e12)), pledge.RewardDebt.Amount.BigInt())
	pledge.Reward.AddAmount(sdk.NewInt(pending.Int64()))

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sp, sdk.NewCoins(amount))

	if err != nil {
		return err
	}

	pledge.Pledged.Sub(amount)

	rewardDebt := new(big.Int).Div(new(big.Int).Mul(pledge.Pledged.Amount.BigInt(), coinPerShare), big.NewInt(1e12))

	pledge.RewardDebt = sdk.NewInt64Coin(pool.Denom.Denom, rewardDebt.Int64())

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return nil
}

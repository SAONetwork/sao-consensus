package keeper

import (
	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error {

	logger := k.Logger(ctx)

	pledge, found_pledge := k.GetPledge(ctx, sp.String())

	if !found_pledge {
		pledge = types.Pledge{
			Creator: sp.String(),
			Pledged: amount,
		}
	}

	pool, found_pool := k.GetPool(ctx)

	if !found_pool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	if amount.Denom != pool.Denom.Denom {
		return sdkerrors.Wrapf(types.ErrDenom, "want %s but %s", pool.Denom.Denom, amount.Denom)
	}

	logger.Debug("pool denom", "amount", pool.Denom)

	pool.Denom = pool.Denom.Add(amount)

	logger.Debug("pool denom", "amount", pool.Denom)

	if found_pledge {
		reward := uint64(amount.Amount.Int64())*(pool.CoinPerShare/1e12) - uint64(pledge.RewardDebt.Amount.Int64())
		pledge.Reward = pledge.Reward.AddAmount(sdk.NewInt(int64(reward)))
	}

	err := k.bank.SendCoinsFromAccountToModule(ctx, sp, types.ModuleName, sdk.NewCoins(amount))

	if err != nil {
		return err
	}

	pledge.Pledged = pledge.Pledged.Add(amount)

	rewardDebt := pledge.Pledged.Amount.Int64() * (int64(pool.CoinPerShare) / 1e12)

	pledge.RewardDebt = sdk.NewInt64Coin(pool.Denom.Denom, rewardDebt)

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

	reward := uint64(amount.Amount.Int64())*pool.CoinPerShare/1e12 - uint64(pledge.RewardDebt.Amount.Int64())

	pledge.Reward.AddAmount(sdk.NewInt(int64(reward)))

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sp, sdk.NewCoins(amount))

	if err != nil {
		return err
	}

	pledge.Pledged.Sub(amount)

	rewardDebt := pledge.Pledged.Amount.Int64() * int64(pool.CoinPerShare) / 1e12

	pledge.RewardDebt = sdk.NewInt64Coin(pool.Denom.Denom, rewardDebt)

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return nil
}

package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrderPledge(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {

	pledge, foundPledge := k.GetPledge(ctx, sp.String())

	if !foundPledge {
		pledge = types.Pledge{
			Creator:             sp.String(),
			TotalOrderPledged:   sdk.NewInt64Coin(sdk.DefaultBondDenom, 0),
			TotalStoragePledged: sdk.NewInt64Coin(sdk.DefaultBondDenom, 0),
			Reward:              sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0),
			RewardDebt:          sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0),
			TotalStorage:        0,
			LastRewardAt:        ctx.BlockTime().Unix(),
		}
	}

	pool, foundPool := k.GetPool(ctx)

	if !foundPool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	pledge.TotalOrderPledged = pledge.TotalOrderPledged.Add(order.Amount)

	params := k.GetParams(ctx)

	coins := sdk.Coins{order.Amount}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.Reward.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.TotalStorage += int64(order.Size_)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	var shardPledge sdk.Coin

	if !params.BlockReward.Amount.IsZero() {
		rewardPerByte := sdk.NewDecCoinFromCoin(params.BlockReward).Amount.QuoInt64(pool.TotalStorage)

		storageDecPledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
		storageDecPledge.Amount = rewardPerByte.MulInt64(int64(order.Shards[sp.String()].Size_ * order.Duration))
		shardPledge, _ = storageDecPledge.TruncateDecimal()

		coins = coins.Add(shardPledge)

		pledge.TotalStoragePledged.Amount = pledge.TotalStoragePledged.Amount.Add(shardPledge.Amount)

	}

	order_pledge_err := k.bank.SendCoinsFromAccountToModule(ctx, sp, types.ModuleName, coins)

	if order_pledge_err != nil {
		return order_pledge_err
	}

	pool.TotalStorage += int64(order.Shards[sp.String()].Size_)

	if !shardPledge.IsZero() {
		order.Shards[sp.String()].Pledge = shardPledge
	}

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)
	return nil
}

func (k Keeper) OrderRelease(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {

	pledge, foundPledge := k.GetPledge(ctx, sp.String())
	if !foundPledge {
		return sdkerrors.Wrap(types.ErrPledgeNotFound, "")
	}

	pool, foundPool := k.GetPool(ctx)

	if !foundPool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.Reward.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	var coins sdk.Coins

	if order != nil {
		pledge.TotalStorage -= int64(order.Size_)

		shardPledge := order.Shards[sp.String()].Pledge

		if !shardPledge.IsZero() {
			coins = coins.Add(shardPledge)
		}

		err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sp, coins)

		if err != nil {
			return err
		}

		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(shardPledge)

		pledge.TotalOrderPledged = pledge.TotalOrderPledged.Sub(order.Amount)

		pool.TotalStorage -= int64(order.Shards[sp.String()].Size_)
	}

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return nil
}

func (k Keeper) OrderSlash(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {
	if order == nil {
		return status.Errorf(codes.NotFound, "order %d not found", order.Id)
	}

	pledge, foundPledge := k.GetPledge(ctx, sp.String())

	if !foundPledge {
		return sdkerrors.Wrap(types.ErrPledgeNotFound, "")
	}

	pool, foundPool := k.GetPool(ctx)

	if !foundPool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.Reward.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	shardPledge := order.Shards[sp.String()].Pledge

	pledge.TotalOrderPledged = pledge.TotalOrderPledged.Sub(order.Amount)

	pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(shardPledge)

	pledge.TotalStorage -= int64(order.Shards[sp.String()].Size_)

	k.SetPledge(ctx, pledge)

	return nil
}

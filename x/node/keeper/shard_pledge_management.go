package keeper

import (
	"math/big"

	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ProjectionPeriodNumerator   = 1
	ProjectionPeriodDenominator = 10
	OrderAmountNumerator        = 1
	OrderAmountDenominator      = 10
	CirculatingNumerator        = 1
	CirculatingDenominator      = 10
)

func (k Keeper) OrderPledge(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {

	pledge, foundPledge := k.GetPledge(ctx, sp.String())

	denom := k.staking.BondDenom(ctx)
	if !foundPledge {
		pledge = types.Pledge{
			Creator:             sp.String(),
			TotalOrderPledged:   sdk.NewInt64Coin(denom, 0),
			TotalStoragePledged: sdk.NewInt64Coin(denom, 0),
			Reward:              sdk.NewInt64DecCoin(denom, 0),
			RewardDebt:          sdk.NewInt64DecCoin(denom, 0),
			TotalStorage:        0,
			// TODO: remove
			LastRewardAt: ctx.BlockHeight(),
		}
	}

	pool, foundPool := k.GetPool(ctx)

	if !foundPool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	pledge.TotalOrderPledged = pledge.TotalOrderPledged.Add(order.Amount)

	params := k.GetParams(ctx)

	coins := sdk.NewCoins()

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.RewardDebt.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	shardPledge := sdk.NewInt64Coin(order.Amount.Denom, 0)
	logger := k.Logger(ctx)

	pool.TotalStorage += int64(order.Shards[sp.String()].Size_)
	pledge.TotalStorage += int64(order.Shards[sp.String()].Size_)

	if !params.BlockReward.Amount.IsZero() {
		//rewardPerByte := sdk.NewDecFromInt(params.BlockReward.Amount).QuoInt64(pool.TotalStorage)
		// rewardPerByte := sdk.NewDecFromBigInt(big.NewInt(1))
		rewardPerByte := sdk.NewDecWithPrec(1, 6)

		storageDecPledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
		// 1. first N% rewards
		projectionPeriod := order.Duration * ProjectionPeriodNumerator / ProjectionPeriodDenominator
		projectionPeriodPledge := rewardPerByte.MulInt64(int64(order.Shards[sp.String()].Size_) * int64(projectionPeriod))
		logger.Error("pledge ", "part1", projectionPeriodPledge)
		storageDecPledge.Amount.AddMut(projectionPeriodPledge)

		// 2. order price N%. collateral amount can be negotiated between client and SP in the future.
		orderAmountPledge := order.Amount.Amount.BigInt()
		orderAmountPledge.Div(orderAmountPledge, big.NewInt(int64(order.Replica))).Mul(orderAmountPledge, big.NewInt(OrderAmountNumerator)).Div(orderAmountPledge, big.NewInt(OrderAmountDenominator))
		logger.Error("pledge ", "part2", orderAmountPledge)
		storageDecPledge.Amount.AddMut(sdk.NewDecFromBigInt(orderAmountPledge))

		// 3. circulating_supply_sp * shard size / network power * ratio
		// pool, found := k.GetPool(ctx)
		// if found {
		// 	concensusPledge := sdk.NewDecFromInt(
		// 		pool.TotalReward.Amount.MulRaw(int64(order.Shards[sp.String()].Size_ * CirculatingNumerator)).
		// 			QuoRaw(CirculatingDenominator * pool.TotalStorage),
		// 	)
		// 	logger.Debug("pledge part3: ", concensusPledge)
		// 	storageDecPledge.Amount.AddMut(concensusPledge)
		// }

		logger.Debug("order pledge ", "amount", storageDecPledge, "pool", pool.TotalStorage, "reward_per_byte", rewardPerByte, "size", order.Shards[sp.String()].Size_, "duration", order.Duration)
		shardPledge, _ = storageDecPledge.TruncateDecimal()

		// set shard pledge to min price if zero
		if shardPledge.IsZero() {
			shardPledge = sdk.NewInt64Coin(params.BlockReward.Denom, 1)
		}

		coins = coins.Add(shardPledge)

		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Add(shardPledge)
		pool.TotalPledged = pool.TotalPledged.Add(shardPledge)
	}

	order_pledge_err := k.bank.SendCoinsFromAccountToModule(ctx, sp, types.ModuleName, coins)

	if order_pledge_err != nil {
		return order_pledge_err
	}

	order.Shards[sp.String()].Pledge = shardPledge

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
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.RewardDebt.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	var coins sdk.Coins

	if order != nil {
		pledge.TotalStorage -= int64(order.Shards[sp.String()].Size_)

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

		pool.TotalPledged = pool.TotalPledged.Sub(shardPledge)
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
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.RewardDebt.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	shardPledge := order.Shards[sp.String()].Pledge

	pledge.TotalOrderPledged = pledge.TotalOrderPledged.Sub(order.Amount)

	pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(shardPledge)

	pledge.TotalStorage -= int64(order.Shards[sp.String()].Size_)

	pool.TotalPledged = pool.TotalPledged.Sub(shardPledge)

	k.SetPledge(ctx, pledge)

	return nil
}

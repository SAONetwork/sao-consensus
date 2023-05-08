package keeper

import (
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
			TotalStoragePledged: sdk.NewInt64Coin(denom, 0),
			Reward:              sdk.NewInt64DecCoin(denom, 0),
			RewardDebt:          sdk.NewInt64DecCoin(denom, 0),
			TotalStorage:        0,
		}
	}

	pool, foundPool := k.GetPool(ctx)

	if !foundPool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	params := k.GetParams(ctx)

	coins := sdk.NewCoins()

	logger := k.Logger(ctx)
	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.RewardDebt.Amount)
		logger.Debug("PledgeTrace: order pledge 1",
			"sp", sp.String(),
			"orderId", order.Id,
			"reward", pledge.Reward.String(),
			"accRewardPerByte", pool.AccRewardPerByte.String(),
			"totalStorage", pledge.TotalStorage,
			"RewardDebt", pledge.RewardDebt.String(),
			"RewardToAdd", pending.String())
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
	}

	shardPledge := sdk.NewInt64Coin(order.Amount.Denom, 0)

	shard := k.order.GetOrderShardBySP(ctx, order, sp.String())
	if shard == nil {
		return status.Errorf(codes.NotFound, "shard of %s not found", sp)
	}
	logger.Debug("PoolTrace: order pledge",
		"totalStorage", pool.TotalStorage,
		"shardSize", shard.Size_)
	pool.TotalStorage += int64(shard.Size_)

	logger.Debug("PledgeTrace: order pledge 2",
		"sp", sp.String(),
		"orderId", order.Id,
		"totalStorage", pledge.TotalStorage,
		"shardSizeToAdd", shard.Size_)
	pledge.TotalStorage += int64(shard.Size_)

	if !params.BlockReward.Amount.IsZero() {
		//rewardPerByte := sdk.NewDecFromInt(params.BlockReward.Amount).QuoInt64(pool.TotalStorage)
		// rewardPerByte := sdk.NewDecFromBigInt(big.NewInt(1))
		rewardPerByte := sdk.NewDecWithPrec(1, 6)

		storageDecPledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
		// 1. first N% rewards
		//projectionPeriod := order.Duration * ProjectionPeriodNumerator / ProjectionPeriodDenominator
		var remainingDuration uint64

		if order.Status == ordertypes.OrderCompleted || order.Status == ordertypes.OrderMigrating {
			remainingDuration = order.Duration + order.CreatedAt - uint64(ctx.BlockHeight())
		} else {
			remainingDuration = order.Duration
		}

		projectionPeriodPledge := k.BlockRewardPledge(remainingDuration, shard.Size_, sdk.NewDecCoinFromDec(denom, rewardPerByte))
		logger.Debug("pledge ", "part1", projectionPeriodPledge)
		storageDecPledge.Amount.AddMut(projectionPeriodPledge)

		// 2. order price N%. collateral amount can be negotiated between client and SP in the future.
		orderAmountPledge := k.StoreRewardPledge(remainingDuration, shard.Size_, order.RewardPerByte)
		logger.Debug("pledge ", "part2", orderAmountPledge)
		storageDecPledge.Amount.AddMut(orderAmountPledge)

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

		logger.Debug("order pledge ", "amount", storageDecPledge, "pool", pool.TotalStorage, "reward_per_byte", rewardPerByte, "size", shard.Size_, "duration", order.Duration)
		var dec sdk.DecCoin
		shardPledge, dec = storageDecPledge.TruncateDecimal()
		if !dec.IsZero() {
			shardPledge = shardPledge.AddAmount(sdk.NewInt(1))
		}

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

	logger.Debug("CoinTrace: order pledge",
		"from", sp.String(),
		"to", types.ModuleName,
		"amount", coins.String())

	shard.Pledge = shardPledge

	newRewardDebt := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage)
	logger.Debug("PledgeTrace: order pledge 3",
		"sp", sp.String(),
		"orderId", order.Id,
		"rewardDebt", pledge.RewardDebt.String(),
		"accRewardPerByte", pool.AccPledgePerByte.String(),
		"totalStorage", pledge.TotalStorage,
		"newRewardDebt", newRewardDebt.String())

	pledge.RewardDebt.Amount = newRewardDebt
	k.SetPledge(ctx, pledge)

	k.order.SetShard(ctx, *shard)

	k.SetPool(ctx, pool)
	return nil
}

func (k Keeper) OrderRelease(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {

	logger := k.Logger(ctx)
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
		logger.Debug("PledgeTrace: order release 1",
			"sp", sp.String(),
			"reward", pledge.Reward.String(),
			"accRewardPerByte", pool.AccRewardPerByte.String(),
			"totalStorage", pledge.TotalStorage,
			"RewardDebt", pledge.RewardDebt.String(),
			"RewardToAdd", pending.String())
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
	}

	var coins sdk.Coins

	if order != nil {
		shard := k.order.GetOrderShardBySP(ctx, order, sp.String())
		if shard == nil {
			return status.Errorf(codes.NotFound, "shard of %s in order %d not found", sp, order.Id)
		}
		logger.Debug("PledgeTrace: order release 2",
			"sp", sp.String(),
			"orderId", order.Id,
			"totalStorage", pledge.TotalStorage,
			"shardSize", shard.Size_)

		shardPledge := shard.Pledge

		if shardPledge.IsZero() {
			return nil
		}
		coins = coins.Add(shardPledge)

		pledge.TotalStorage -= int64(shard.Size_)

		logger.Debug("CoinTrace: order release",
			"from", types.ModuleName,
			"to", sp.String(),
			"amount", coins.String())

		err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sp, coins)

		if err != nil {
			return err
		}

		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(shardPledge)

		logger.Debug("PoolTrace: order release",
			"totalStorage", pool.TotalStorage,
			"shardSizeToSub", shard.Size_)

		pool.TotalStorage -= int64(shard.Size_)

		pool.TotalPledged = pool.TotalPledged.Sub(shardPledge)
	}

	newRewardDebt := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage)
	logger.Debug("PledgeTrace: order release 3",
		"sp", sp.String(),
		"rewardDebt", pledge.RewardDebt.String(),
		"accRewardPerByte", pool.AccPledgePerByte.String(),
		"totalStorage", pledge.TotalStorage,
		"newRewardDebt", newRewardDebt.String())

	pledge.RewardDebt.Amount = newRewardDebt
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

	shard := k.order.GetOrderShardBySP(ctx, order, sp.String())
	if shard == nil {
		return status.Errorf(codes.NotFound, "shard of %s not found", sp)
	}

	shardPledge := shard.Pledge

	pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(shardPledge)

	pledge.TotalStorage -= int64(shard.Size_)

	pool.TotalPledged = pool.TotalPledged.Sub(shardPledge)

	k.SetPledge(ctx, pledge)

	return nil
}

//
//func (Keeper) OrderPricePledge(remainingDuration, duration uint64, amount sdk.Int, replica int32) sdk.Dec {
//	// 2. order price N%. collateral amount can be negotiated between client and SP in the future.
//	orderAmountPledge := amount.BigInt()
//	if duration != remainingDuration {
//		orderAmountPledge.Mul(orderAmountPledge, big.NewInt(int64(remainingDuration))).Div(orderAmountPledge, big.NewInt(int64(duration)))
//	}
//	orderAmountPledge.Div(orderAmountPledge, big.NewInt(int64(replica))).Mul(orderAmountPledge, big.NewInt(OrderAmountNumerator)).Div(orderAmountPledge, big.NewInt(OrderAmountDenominator))
//
//	return sdk.NewDecFromBigInt(orderAmountPledge)
//}

func (Keeper) BlockRewardPledge(duration uint64, size uint64, rewardPerByte sdk.DecCoin) sdk.Dec {

	//rewardPerByte := sdk.NewDecFromInt(params.BlockReward.Amount).QuoInt64(pool.TotalStorage)
	// 1. first N% rewards
	projectionPeriod := duration * ProjectionPeriodNumerator / ProjectionPeriodDenominator
	projectionPeriodPledge := rewardPerByte.Amount.MulInt64(int64(size) * int64(projectionPeriod))

	return projectionPeriodPledge

	//return rewardPerByte.
	//	MulInt64(int64(size)).
	//	MulInt64(int64(duration)).
	//	MulInt64(CirculatingNumerator).
	//	QuoInt64(CirculatingDenominator)
}

func (Keeper) StoreRewardPledge(duration uint64, size uint64, rewardPerByte sdk.DecCoin) sdk.Dec {

	//rewardPerByte := sdk.NewDecFromInt(params.BlockReward.Amount).QuoInt64(pool.TotalStorage)
	// 1. first N% rewards
	projectionPeriod := duration * OrderAmountNumerator / OrderAmountDenominator
	projectionPeriodPledge := rewardPerByte.Amount.MulInt64(int64(size) * int64(projectionPeriod))

	return projectionPeriodPledge

	//return rewardPerByte.
	//	MulInt64(int64(size)).
	//	MulInt64(int64(duration)).
	//	MulInt64(CirculatingNumerator).
	//	QuoInt64(CirculatingDenominator)
}

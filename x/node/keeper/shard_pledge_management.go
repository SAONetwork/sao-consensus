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

func (k Keeper) ShardPledge(ctx sdk.Context, shard *ordertypes.Shard, unitPrice sdk.DecCoin) error {

	pledge, foundPledge := k.GetPledge(ctx, shard.Sp)

	if !foundPledge {
		return status.Error(codes.NotFound, "not plegded yet")
	}

	denom := k.staking.BondDenom(ctx)

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
			"sp", shard.Sp,
			"orderId", shard.OrderId,
			"reward", pledge.Reward.String(),
			"accRewardPerByte", pool.AccRewardPerByte.String(),
			"totalStorage", pledge.TotalStorage,
			"RewardDebt", pledge.RewardDebt.String(),
			"RewardToAdd", pending.String())
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
	}

	shardPledge := sdk.NewInt64Coin(denom, 0)

	logger.Debug("PoolTrace: order pledge",
		"totalStorage", pool.TotalStorage,
		"shardSize", shard.Size_)

	logger.Debug("PledgeTrace: order pledge 2",
		"sp", shard.Sp,
		"orderId", shard.OrderId,
		"totalStorage", pledge.TotalStorage,
		"shardSizeToAdd", shard.Size_)

	if uint64(pledge.TotalStorage-pledge.UsedStorage) < shard.Size_ {
		return sdkerrors.Wrap(types.ErrAvailableVstorage, "no enough available vstorage")
	}

	pledge.UsedStorage += int64(shard.Size_)

	storageDecPledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)

	// 1. order price N%. collateral amount can be negotiated between client and SP in the future.
	orderAmountPledge := k.StoreRewardPledge(shard.Duration, shard.Size_, unitPrice)
	logger.Debug("pledge ", "part2", orderAmountPledge)
	storageDecPledge.Amount.AddMut(orderAmountPledge)

	// 2. circulating_supply_sp * shard size / network power * ratio
	// pool, found := k.GetPool(ctx)
	// if found {
	// 	concensusPledge := sdk.NewDecFromInt(
	// 		pool.TotalReward.Amount.MulRaw(int64(order.Shards[shard.Sp].Size_ * CirculatingNumerator)).
	// 			QuoRaw(CirculatingDenominator * pool.TotalStorage),
	// 	)
	// 	logger.Debug("pledge part3: ", concensusPledge)
	// 	storageDecPledge.Amount.AddMut(concensusPledge)
	// }

	logger.Debug("order pledge ", "amount", storageDecPledge, "pool", pool.TotalStorage, "size", shard.Size_, "duration", shard.Duration)
	var dec sdk.DecCoin
	shardPledge, dec = storageDecPledge.TruncateDecimal()
	if !dec.IsZero() {
		shardPledge = shardPledge.AddAmount(sdk.NewInt(1))
	}

	for _, renewInfo := range shard.RenewInfos {
		if shardPledge.IsLT(renewInfo.Pledge) {
			shardPledge = renewInfo.Pledge
		}
	}

	coins = coins.Add(shardPledge)

	var err error
	if len(shard.RenewInfos) != 0 {
		balance := k.bank.GetBalance(ctx, sdk.MustAccAddressFromBech32(shard.Sp), denom)
		if balance.IsGTE(shardPledge) {
			err = k.bank.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(shard.Sp), types.ModuleName, coins)
		} else {
			err = k.bank.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(shard.Sp), types.ModuleName, sdk.Coins{balance})
			pledgeDebt, found := k.GetPledgeDebt(ctx, shard.Sp)
			if found {
				pledgeDebt.Debt = pledgeDebt.Debt.Add(shardPledge.Sub(balance))
			} else {
				pledgeDebt = types.PledgeDebt{
					Sp:   shard.Sp,
					Debt: shardPledge.Sub(balance),
				}
			}
			k.SetPledgeDebt(ctx, pledgeDebt)
		}
	} else {
		err = k.bank.SendCoinsFromAccountToModule(ctx, sdk.MustAccAddressFromBech32(shard.Sp), types.ModuleName, coins)
	}

	if err != nil {
		return err
	}

	logger.Debug("CoinTrace: order pledge",
		"from", shard.Sp,
		"to", types.ModuleName,
		"amount", coins.String())

	shard.Pledge = shardPledge

	newRewardDebt := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage)
	logger.Debug("PledgeTrace: order pledge 3",
		"sp", shard.Sp,
		"orderId", shard.OrderId,
		"rewardDebt", pledge.RewardDebt.String(),
		"accRewardPerByte", pool.AccPledgePerByte.String(),
		"totalStorage", pledge.TotalStorage,
		"newRewardDebt", newRewardDebt.String())

	pledge.RewardDebt.Amount = newRewardDebt

	pledge.UsedStorage += int64(shard.Size_)

	k.SetPledge(ctx, pledge)

	k.order.SetShard(ctx, *shard)

	return nil
}

func (k Keeper) ShardRelease(ctx sdk.Context, sp sdk.AccAddress, shard *ordertypes.Shard) error {

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

	if shard != nil {
		logger.Debug("PledgeTrace: order release 2",
			"sp", sp.String(),
			"orderId", shard.OrderId,
			"totalStorage", pledge.TotalStorage,
			"shardSize", shard.Size_)

		shardPledge := shard.Pledge

		k.RepayPledgeDebt(ctx, shard.Sp, []*sdk.Coin{&shardPledge})

		if !shardPledge.IsZero() {
			logger.Debug("CoinTrace: order release",
				"from", types.ModuleName,
				"to", sp.String(),
				"amount", shardPledge.String())

			err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sp, sdk.Coins{shardPledge})
			if err != nil {
				return err
			}
		}

		pledge.UsedStorage -= int64(shard.Size_)

		logger.Debug("PoolTrace: order release",
			"totalStorage", pool.TotalStorage,
			"shardSizeToSub", shard.Size_)

		pledge.UsedStorage -= int64(shard.Size_)

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

	// 1. first N% block rewards

	return rewardPerByte.
		Amount.
		MulInt64(int64(size)).
		MulInt64(int64(duration)).
		MulInt64(ProjectionPeriodNumerator).
		QuoInt64(ProjectionPeriodDenominator)
}

func (Keeper) StoreRewardPledge(duration uint64, size uint64, rewardPerByte sdk.DecCoin) sdk.Dec {

	// 2. first N% store rewards

	return rewardPerByte.
		Amount.
		MulInt64(int64(size)).
		MulInt64(int64(duration)).
		MulInt64(OrderAmountNumerator).
		QuoInt64(OrderAmountDenominator)
}

func (k Keeper) RepayPledgeDebt(ctx sdk.Context, sp string, rewards []*sdk.Coin) {
	pledgeDebt, found := k.GetPledgeDebt(ctx, sp)
	logger := k.Logger(ctx)
	if found {
		logger.Debug("CoinTrace: repay pledge debt", "rewards", rewards, "debt", pledgeDebt.Debt.String())

		for _, reward := range rewards {
			if reward.IsGTE(pledgeDebt.Debt) {
				*reward = reward.Sub(pledgeDebt.Debt)
				k.RemovePledgeDebt(ctx, sp)
				return
			} else {
				pledgeDebt.Debt = pledgeDebt.Debt.Sub(*reward)
				*reward = sdk.NewCoin(reward.Denom, sdk.NewInt(0))
			}
		}
		k.SetPledgeDebt(ctx, pledgeDebt)
	}
}

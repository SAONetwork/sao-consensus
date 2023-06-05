package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	denom := k.staking.BondDenom(ctx)
	if !foundPledge {
		pledge = types.Pledge{
			Creator:             shard.Sp,
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

	pool.PendingStorage += int64(shard.Size_)

	logger.Debug("PledgeTrace: order pledge 2",
		"sp", shard.Sp,
		"orderId", shard.OrderId,
		"totalStorage", pledge.TotalStorage,
		"shardSizeToAdd", shard.Size_)
	pledge.TotalStorage += int64(shard.Size_)

	if !pool.RewardPerBlock.Amount.IsZero() {
		rewardPerByte := pool.RewardPerBlock.Amount.Quo(sdk.NewDec(pool.TotalStorage))
		// rewardPerByte := sdk.NewDecFromBigInt(big.NewInt(1))
		//rewardPerByte := sdk.NewDecWithPrec(1, 6)

		storageDecPledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
		// 1. first N% rewards
		//projectionPeriod := order.Duration * ProjectionPeriodNumerator / ProjectionPeriodDenominator

		projectionPeriodPledge := k.BlockRewardPledge(shard.Duration, shard.Size_, sdk.NewDecCoinFromDec(denom, rewardPerByte))
		logger.Debug("pledge ", "part1", projectionPeriodPledge)
		storageDecPledge.Amount.AddMut(projectionPeriodPledge)

		// 2. order price N%. collateral amount can be negotiated between client and SP in the future.
		orderAmountPledge := k.StoreRewardPledge(shard.Duration, shard.Size_, unitPrice)
		logger.Debug("pledge ", "part2", orderAmountPledge)
		storageDecPledge.Amount.AddMut(orderAmountPledge)

		// 3. circulating_supply_sp * shard size / network power * ratio
		// pool, found := k.GetPool(ctx)
		// if found {
		// 	concensusPledge := sdk.NewDecFromInt(
		// 		pool.TotalReward.Amount.MulRaw(int64(order.Shards[shard.Sp].Size_ * CirculatingNumerator)).
		// 			QuoRaw(CirculatingDenominator * pool.TotalStorage),
		// 	)
		// 	logger.Debug("pledge part3: ", concensusPledge)
		// 	storageDecPledge.Amount.AddMut(concensusPledge)
		// }

		logger.Debug("order pledge ", "amount", storageDecPledge, "pool", pool.TotalStorage, "reward_per_byte", rewardPerByte, "size", shard.Size_, "duration", shard.Duration)
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

		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Add(shardPledge)
		pool.TotalPledged = pool.TotalPledged.Add(shardPledge)
	}

	err := k.DoPledge(ctx, &pledge, shardPledge)
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
	k.SetPledge(ctx, pledge)

	k.order.SetShard(ctx, *shard)

	k.SetPool(ctx, pool)
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

		k.RepayDebt(ctx, shard.Sp, []*sdk.Coin{&shardPledge})

		if pledge.LoanStrategy != types.LoanStrategyLoanFirst {
			err := k.RepayLoan(ctx, &pledge, []*sdk.Coin{&shardPledge})
			if err != nil {
				return err
			}
		}

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

		pledge.TotalStorage -= int64(shard.Size_)

		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(shard.Pledge)

		pool.PendingStorage -= int64(shard.Size_)

		pool.TotalPledged = pool.TotalPledged.Sub(shard.Pledge)
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

func (k Keeper) RepayDebt(ctx sdk.Context, sp string, rewards []*sdk.Coin) {
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

func (k Keeper) RepayLoan(ctx sdk.Context, pledge *types.Pledge, rewards []*sdk.Coin) error {
	repay := sdk.NewCoin(pledge.LoanPledged.Denom, sdk.NewInt(0))
	for _, reward := range rewards {
		if reward.IsGTE(pledge.LoanPledged) {
			*reward = reward.Sub(pledge.LoanPledged)
			pledge.LoanPledged = sdk.NewCoin(reward.Denom, sdk.NewInt(0))
			repay = repay.Add(pledge.LoanPledged)
			break
		} else {
			pledge.LoanPledged = pledge.LoanPledged.Sub(*reward)
			*reward = sdk.NewCoin(reward.Denom, sdk.NewInt(0))
			repay = repay.Add(*reward)
		}
	}
	err := k.loan.Repay(ctx, repay)
	return err
}

func (k Keeper) DoPledge(ctx sdk.Context, pledge *types.Pledge, shardPledge sdk.Coin) error {
	switch pledge.LoanStrategy {
	case types.LoanStrategyDisable:
		left, err := k.BalancePledge(ctx, pledge.Creator, shardPledge)
		if err != nil {
			return err
		}
		if !left.IsZero() {
			err = k.DebtPledge(ctx, pledge.Creator, left)
			return err
		}
	case types.LoanStrategyBalanceFirst:
		left, err := k.BalancePledge(ctx, pledge.Creator, shardPledge)
		if err != nil {
			return err
		}
		if !left.IsZero() {
			left, err = k.LoanPledge(ctx, pledge, left)
			if err != nil {
				return err
			}
			if !left.IsZero() {
				err = k.DebtPledge(ctx, pledge.Creator, left)
				return err
			}
		}
	case types.LoanStrategyLoanFirst:
		left, err := k.LoanPledge(ctx, pledge, shardPledge)
		if err != nil {
			return err
		}
		if !left.IsZero() {
			left, err = k.BalancePledge(ctx, pledge.Creator, left)
			if err != nil {
				return err
			}
			if !left.IsZero() {
				err = k.DebtPledge(ctx, pledge.Creator, left)
				return err
			}
		}
	}
	return nil
}

func (k Keeper) BalancePledge(ctx sdk.Context, sp string, amount sdk.Coin) (sdk.Coin, error) {
	accAddr := sdk.MustAccAddressFromBech32(sp)
	balance := k.bank.GetBalance(ctx, accAddr, amount.Denom)
	if !balance.IsZero() {
		if balance.IsGTE(amount) {
			err := k.bank.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.Coins{amount})
			return sdk.NewCoin(amount.Denom, sdk.NewInt(0)), err

		} else {
			err := k.bank.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.Coins{balance})
			return amount.Sub(balance), err
		}
	} else {
		return amount, nil
	}
}

func (k Keeper) LoanPledge(ctx sdk.Context, pledge *types.Pledge, amount sdk.Coin) (sdk.Coin, error) {
	loanedOut, err := k.loan.LoanOut(ctx, amount)
	if err != nil {
		return sdk.Coin{}, err
	}
	if !loanedOut.IsZero() {
		pledge.LoanPledged = pledge.LoanPledged.Add(loanedOut)
		if loanedOut.IsGTE(amount) {
			return sdk.NewCoin(amount.Denom, sdk.NewInt(0)), nil
		} else {
			return amount.Sub(loanedOut), nil
		}
	} else {
		return amount, nil
	}
}

func (k Keeper) DebtPledge(ctx sdk.Context, sp string, debt sdk.Coin) error {
	pledgeDebt, found := k.GetPledgeDebt(ctx, sp)
	if found {
		pledgeDebt.Debt = pledgeDebt.Debt.Add(debt)
	} else {
		pledgeDebt = types.PledgeDebt{
			Sp:   sp,
			Debt: debt,
		}
	}
	k.SetPledgeDebt(ctx, pledgeDebt)
	return nil
}

//func (k Keeper) DoRelease(ctx sdk.Context, sp string)

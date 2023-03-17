package keeper

import (
	"fmt"
	"github.com/SaoNetwork/sao/x/market/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Deposit(ctx sdk.Context, order ordertypes.Order) error {
	logger := k.Logger(ctx)
	amount := sdk.NewDecCoinFromCoin(order.Amount)

	if amount.IsZero() {
		return sdkerrors.Wrap(types.ErrInvalidAmount, "")
	}

	err := k.bank.SendCoinsFromModuleToModule(ctx, ordertypes.ModuleName, types.ModuleName, sdk.Coins{order.Amount})
	if err != nil {
		return err
	}
	logger.Debug("CoinTrace: deposit", "from", ordertypes.ModuleName, "to", types.ModuleName, "amount", order.Amount.String())

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			return status.Errorf(codes.NotFound, "shard %d not found", id)
		}
		sp := shard.Sp
		//workerName := fmt.Sprintf("%s-%s", amount.Denom, sp)
		//worker, found := k.GetWorker(ctx, workerName)
		//if !found {
		//	worker = types.Worker{
		//		Workername:      workerName,
		//		Storage:         0,
		//		Reward:          sdk.NewInt64DecCoin(amount.Denom, 0),
		//		IncomePerSecond: sdk.NewInt64DecCoin(amount.Denom, 0),
		//	}
		//}
		//incomePerSecond := amount.Amount.QuoInt64(int64(order.Replica)).QuoInt64(duration)
		//if worker.Storage > 0 {
		//	reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockTime().Unix() - worker.LastRewardAt)
		//	worker.LastRewardAt = ctx.BlockTime().Unix()
		//	worker.Reward.Amount = worker.Reward.Amount.Add(reward)
		//	worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Add(incomePerSecond)
		//}
		//worker.Storage += uint64(shard.Size_)

		err := k.WorkerAppend(ctx, &order, sp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, order ordertypes.Order) (sdk.Coin, error) {
	logger := k.Logger(ctx)

	amount := sdk.NewDecCoinFromCoin(order.Amount)
	duration := int64(order.Duration)

	if amount.IsZero() {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrInvalidAmount, "")
	}
	orderFinishHeight := int64(order.CreatedAt) + duration

	refundCoin := sdk.NewCoin(amount.Denom, sdk.NewInt(0))

	if orderFinishHeight < ctx.BlockHeight() {
		return sdk.Coin{}, status.Errorf(
			codes.Aborted,
			"invalid height to withdraw, order: %v, finishHeight: %v, currentHeight: %v", order.Id, orderFinishHeight, ctx.BlockHeight(),
		)
	} else if orderFinishHeight > ctx.BlockHeight() {
		incomePerBlock := amount.Amount.QuoInt64(duration)

		refund := incomePerBlock.MulInt64(orderFinishHeight - ctx.BlockHeight()).TruncateInt()

		refundCoin = sdk.NewCoin(amount.Denom, refund)

		err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, ordertypes.ModuleName, sdk.Coins{refundCoin})
		if err != nil {
			return sdk.Coin{}, err
		}
		logger.Debug("CoinTrace: withdraw", "from", types.ModuleName, "to", ordertypes.ModuleName, "amount", refundCoin.String())
	}

	for _, id := range order.Shards {

		shard, found := k.order.GetShard(ctx, id)
		if !found {
			return sdk.Coin{}, status.Errorf(codes.NotFound, "shard %d not found", id)
		}
		sp := shard.Sp
		//workerName := fmt.Sprintf("%s-%s", amount.Denom, sp)
		//worker, _ := k.GetWorker(ctx, workerName)
		//incomePerSecond := amount.Amount.QuoInt64(int64(order.Replica)).QuoInt64(duration)
		//reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockTime().Unix() - worker.LastRewardAt)
		//worker.Reward.Amount = worker.Reward.Amount.Add(reward)
		//worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Sub(incomePerSecond)
		//worker.Storage -= uint64(shard.Size_)
		//worker.LastRewardAt = ctx.BlockTime().Unix()
		//k.SetWorker(ctx, worker)

		err := k.WorkerRelease(ctx, &order, sp)
		if err != nil {
			return sdk.Coin{}, err
		}
	}

	return refundCoin, nil
}

func (k Keeper) Claim(ctx sdk.Context, denom string, sp string) (sdk.Coin, error) {

	logger := k.Logger(ctx)

	empty := sdk.NewCoin(denom, sdk.NewInt(0))

	workerName := fmt.Sprintf("%s-%s", denom, sp)
	worker, found := k.GetWorker(ctx, workerName)
	if !found {
		// return nil if not found worker
		logger.Error("denom worker not found", "denom", denom, "worker", sp)
		return empty, nil
	}

	reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
	logger.Debug("WorkerTrace: claim 1",
		"Worker", workerName,
		"reward", worker.Reward.String(),
		"rewardToAdd", reward.String())
	worker.Reward.Amount = worker.Reward.Amount.Add(reward)

	if worker.Reward.Amount.TruncateInt().IsZero() {
		// return nil , if  reward is 0
		logger.Error("no reward", "worker", workerName)
		return empty, nil
	}

	rewardCoin := sdk.NewCoin(denom, worker.Reward.Amount.TruncateInt())

	spAcc := sdk.MustAccAddressFromBech32(sp)

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, spAcc, sdk.Coins{rewardCoin})

	if err != nil {
		return empty, err
	}
	logger.Debug("CoinTrace: claim", "from", types.ModuleName, "to", spAcc.String(), "amount", rewardCoin.String())

	logger.Debug("WorkerTrace: claim 2",
		"Worker", workerName,
		"reward", worker.Reward.String(),
		"rewardToSub", rewardCoin.String(),
		"lastRewardAt", worker.LastRewardAt,
		"currentHeight", ctx.BlockHeight())
	worker.Reward.Amount = worker.Reward.Amount.Sub(sdk.NewDecFromInt(rewardCoin.Amount))
	worker.LastRewardAt = ctx.BlockHeight()

	k.SetWorker(ctx, worker)

	return rewardCoin, nil
}

func (k Keeper) Migrate(ctx sdk.Context, order ordertypes.Order, from string, to string) error {
	// from sp worker settlement
	err := k.WorkerRelease(ctx, &order, from)
	if err != nil {
		return err
	}

	// to sp worker begin work
	err = k.WorkerAppend(ctx, &order, to)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) Release(ctx sdk.Context, order ordertypes.Order, sp string) (sdk.Coin, error) {
	logger := k.Logger(ctx)
	amount := sdk.NewDecCoinFromCoin(order.Amount)
	empty := sdk.NewCoin(amount.Denom, sdk.NewInt(0))
	duration := int64(order.Duration)
	orderFinishHeight := int64(order.CreatedAt) + duration
	if orderFinishHeight < ctx.BlockHeight() {
		return empty, status.Errorf(
			codes.Aborted,
			"invalid height to withdraw, order: %v, finishHeight: %v, currentHeight: %v", order.Id, orderFinishHeight, ctx.BlockHeight(),
		)
	}

	incomePerBlock := amount.Amount.QuoInt64(duration * int64(order.Replica))

	refund := incomePerBlock.MulInt64(orderFinishHeight - ctx.BlockHeight()).TruncateInt()

	refundCoin := sdk.NewCoin(amount.Denom, refund)

	err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, ordertypes.ModuleName, sdk.Coins{refundCoin})
	if err != nil {
		return empty, err
	}
	logger.Debug("CoinTrace: release single worker", "from", types.ModuleName, "to", ordertypes.ModuleName, "amount", refundCoin.String())

	err = k.WorkerRelease(ctx, &order, sp)
	if err != nil {
		return empty, err
	}

	return refundCoin, nil
}

func (k *Keeper) WorkerRelease(ctx sdk.Context, order *ordertypes.Order, sp string) error {
	logger := k.Logger(ctx)

	if order == nil {
		return status.Errorf(codes.NotFound, "WorkerRelease order not found")
	}

	amount := sdk.NewDecCoinFromCoin(order.Amount)
	shard := order.Shards[sp]

	workerName := fmt.Sprintf("%s-%s", amount.Denom, sp)
	worker, foundWorker := k.GetWorker(ctx, workerName)
	if !foundWorker {
		return status.Errorf(codes.NotFound, "worker: %v not found", workerName)
	}
	incomePerSecond := amount.Amount.QuoInt64(int64(order.Replica) * int64(order.Duration))
	reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
	logger.Debug("WorkerTrace: worker release",
		"Worker", workerName,
		"orderId", order.Id,
		"reward", worker.Reward.String(),
		"rewardToAdd", reward.String(),
		"lastRewardAt", worker.LastRewardAt,
		"currentHeight", ctx.BlockHeight(),
		"incomePerBlock", worker.IncomePerSecond.String(),
		"incomePerBlockToSub", incomePerSecond.String())
	worker.Reward.Amount = worker.Reward.Amount.Add(reward)
	worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Sub(incomePerSecond)
	worker.Storage -= shard.Size_
	worker.LastRewardAt = ctx.BlockHeight()
	k.SetWorker(ctx, worker)
	return nil
}

func (k *Keeper) WorkerAppend(ctx sdk.Context, order *ordertypes.Order, sp string) error {
	logger := k.Logger(ctx)

	if order == nil {
		return status.Errorf(codes.NotFound, "WorkerRelease order not found")
	}

	amount := sdk.NewDecCoinFromCoin(order.Amount)
	shard := order.Shards[sp]
	duration := int64(order.Duration)

	workerName := fmt.Sprintf("%s-%s", amount.Denom, sp)
	worker, found := k.GetWorker(ctx, workerName)
	if !found {
		worker = types.Worker{
			Workername:      workerName,
			Storage:         0,
			Reward:          sdk.NewInt64DecCoin(amount.Denom, 0),
			IncomePerSecond: sdk.NewInt64DecCoin(amount.Denom, 0),
		}
	}

	incomePerSecond := amount.Amount.QuoInt64(int64(order.Replica) * duration)
	if worker.Storage > 0 {
		reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
		logger.Debug("WorkerTrace: deposit 1",
			"Worker", workerName,
			"orderId", order.Id,
			"reward", worker.Reward.String(),
			"rewardToAdd", reward.String(),
			"lastRewardAt", worker.LastRewardAt,
			"currentHeight", ctx.BlockHeight(),
			"incomePerBlock", worker.IncomePerSecond.String())
		worker.LastRewardAt = ctx.BlockHeight()
		worker.Reward.Amount = worker.Reward.Amount.Add(reward)
	}

	logger.Debug("WorkerTrace: deposit 2",
		"Worker", workerName,
		"orderId", order.Id,
		"incomePerBlock", worker.IncomePerSecond.String(),
		"incomePerBlockToAdd", incomePerSecond.String())
	worker.Storage += shard.Size_
	worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Add(incomePerSecond)

	k.SetWorker(ctx, worker)
	return nil
}

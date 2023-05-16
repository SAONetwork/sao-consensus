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
		if shard.Status == ordertypes.ShardCompleted {
			err := k.WorkerAppend(ctx, &order, &shard)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, order ordertypes.Order) (sdk.Coin, error) {
	logger := k.Logger(ctx)

	amount := sdk.NewDecCoinFromCoin(order.Amount)

	if amount.IsZero() {
		return sdk.Coin{}, sdkerrors.Wrap(types.ErrInvalidAmount, "")
	}

	// add all refund parts,
	// refundDec = amount - price * size * replica * duration
	refundDec := amount.Amount.Sub(order.UnitPrice.Amount.MulInt64(int64(order.Size_)).MulInt64(int64(order.Replica)).MulInt64(int64(order.Duration)))

	for _, id := range order.Shards {

		shard, found := k.order.GetShard(ctx, id)
		if !found {
			continue
		}

		shardIncomePerBlock := order.UnitPrice.Amount.MulInt64(int64(shard.Size_))

		if shard.Status == ordertypes.ShardCompleted {

			// refundDec += price * shardSize * (shardExpiredAt - currentHeight)
			refundDec = refundDec.Add(shardIncomePerBlock.MulInt64(int64(shard.CreatedAt+shard.Duration) - ctx.BlockHeight()))

			err := k.WorkerRelease(ctx, &order, &shard)
			if err != nil {
				return sdk.Coin{}, err
			}
		} else {
			// refundDec += price * shardSize * shardDuration
			refundDec = refundDec.Add(shardIncomePerBlock.MulInt64(int64(order.Duration)))
		}
	}

	refundCoin, _ := sdk.NewDecCoinFromDec(amount.Denom, refundDec).TruncateDecimal()

	err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, ordertypes.ModuleName, sdk.Coins{refundCoin})
	if err != nil {
		return sdk.Coin{}, err
	}
	logger.Debug("CoinTrace: withdraw", "from", types.ModuleName, "to", ordertypes.ModuleName, "amount", refundCoin.String())

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
	fromShard := k.order.GetOrderShardBySP(ctx, &order, from)
	// from sp worker settlement
	err := k.WorkerRelease(ctx, &order, fromShard)
	if err != nil {
		return err
	}

	toShard := k.order.GetOrderShardBySP(ctx, &order, to)
	// to sp worker begin work
	err = k.WorkerAppend(ctx, &order, toShard)
	if err != nil {
		return err
	}

	return nil
}

//func (k Keeper) Release(ctx sdk.Context, order ordertypes.Order, sp string) (sdk.Coin, error) {
//	logger := k.Logger(ctx)
//	amount := sdk.NewDecCoinFromCoin(order.Amount)
//	empty := sdk.NewCoin(amount.Denom, sdk.NewInt(0))
//	duration := int64(order.Duration)
//	orderFinishHeight := int64(order.CreatedAt) + duration
//	if orderFinishHeight < ctx.BlockHeight() {
//		return empty, status.Errorf(
//			codes.Aborted,
//			"invalid height to withdraw, order: %v, finishHeight: %v, currentHeight: %v", order.Id, orderFinishHeight, ctx.BlockHeight(),
//		)
//	}
//
//	incomePerBlock := amount.Amount.QuoInt64(duration * int64(order.Replica))
//
//	refund := incomePerBlock.MulInt64(orderFinishHeight - ctx.BlockHeight()).TruncateInt()
//
//	refundCoin := sdk.NewCoin(amount.Denom, refund)
//
//	err := k.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, ordertypes.ModuleName, sdk.Coins{refundCoin})
//	if err != nil {
//		return empty, err
//	}
//	logger.Debug("CoinTrace: release single worker", "from", types.ModuleName, "to", ordertypes.ModuleName, "amount", refundCoin.String())
//
//	shard := k.order.GetOrderShardBySP(ctx, &order, sp)
//	err = k.WorkerRelease(ctx, &order, shard)
//	if err != nil {
//		return empty, err
//	}
//
//	return refundCoin, nil
//}

func (k Keeper) WorkerRelease(ctx sdk.Context, order *ordertypes.Order, shard *ordertypes.Shard) error {
	logger := k.Logger(ctx)

	if order == nil {
		return status.Errorf(codes.NotFound, "WorkerRelease order not found")
	}
	if shard == nil {
		return status.Errorf(codes.NotFound, "WorkerRelease shard not found")
	}

	amount := sdk.NewDecCoinFromCoin(order.Amount)

	workerName := fmt.Sprintf("%s-%s", amount.Denom, shard.Sp)
	worker, foundWorker := k.GetWorker(ctx, workerName)
	if !foundWorker {
		return status.Errorf(codes.NotFound, "worker: %v not found", workerName)
	}
	IncomePerBlock := order.UnitPrice.Amount.MulInt64(int64(shard.Size_))
	reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
	logger.Debug("WorkerTrace: worker release",
		"Worker", workerName,
		"orderId", order.Id,
		"reward", worker.Reward.String(),
		"rewardToAdd", reward.String(),
		"lastRewardAt", worker.LastRewardAt,
		"currentHeight", ctx.BlockHeight(),
		"incomePerBlock", worker.IncomePerSecond.String(),
		"incomePerBlockToSub", IncomePerBlock.String())
	worker.Reward.Amount = worker.Reward.Amount.Add(reward)
	worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Sub(IncomePerBlock)
	worker.Storage -= shard.Size_
	worker.LastRewardAt = ctx.BlockHeight()
	k.SetWorker(ctx, worker)
	return nil
}

func (k Keeper) WorkerAppend(ctx sdk.Context, order *ordertypes.Order, shard *ordertypes.Shard) error {
	logger := k.Logger(ctx)

	if order == nil {
		return status.Errorf(codes.NotFound, "WorkerRelease order not found")
	}
	if shard == nil {
		return status.Errorf(codes.NotFound, "WorkerRelease shard not found")
	}

	amount := sdk.NewDecCoinFromCoin(order.Amount)

	workerName := fmt.Sprintf("%s-%s", amount.Denom, shard.Sp)
	worker, found := k.GetWorker(ctx, workerName)
	if !found {
		worker = types.Worker{
			Workername:      workerName,
			Storage:         0,
			Reward:          sdk.NewInt64DecCoin(amount.Denom, 0),
			IncomePerSecond: sdk.NewInt64DecCoin(amount.Denom, 0),
		}
	}

	IncomePerBlock := order.UnitPrice.Amount.MulInt64(int64(shard.Size_))
	if worker.Storage > 0 {
		reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
		reward = reward.Add(IncomePerBlock.MulInt64(ctx.BlockHeight() - int64(shard.CreatedAt)))
		logger.Debug("WorkerTrace: deposit 1",
			"Worker", workerName,
			"orderId", order.Id,
			"reward", worker.Reward.String(),
			"rewardToAdd", reward.String(),
			"lastRewardAt", worker.LastRewardAt,
			"currentHeight", ctx.BlockHeight(),
			"incomePerBlock", worker.IncomePerSecond.String())
		worker.Reward.Amount = worker.Reward.Amount.Add(reward)
	}
	worker.LastRewardAt = ctx.BlockHeight()

	logger.Debug("WorkerTrace: deposit 2",
		"Worker", workerName,
		"orderId", order.Id,
		"incomePerBlock", worker.IncomePerSecond.String(),
		"incomePerBlockToAdd", IncomePerBlock.String())
	worker.Storage += shard.Size_
	worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Add(IncomePerBlock)

	k.SetWorker(ctx, worker)

	return nil
}

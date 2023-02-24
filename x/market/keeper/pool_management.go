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

	amount := sdk.NewDecCoinFromCoin(order.Amount)
	duration := int64(order.Duration)

	if amount.IsZero() {
		return sdkerrors.Wrap(types.ErrInvalidAmount, "")
	}

	err := k.bank.SendCoinsFromModuleToModule(ctx, ordertypes.ModuleName, types.ModuleName, sdk.Coins{order.Amount})
	if err != nil {
		return err
	}

	for sp, shard := range order.Shards {
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
		incomePerSecond := amount.Amount.QuoInt64(int64(order.Replica)).QuoInt64(duration)
		if worker.Storage > 0 {
			reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
			worker.LastRewardAt = ctx.BlockHeight()
			worker.Reward.Amount = worker.Reward.Amount.Add(reward)
			worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Add(incomePerSecond)
		}
		worker.Storage += shard.Size_

		k.SetWorker(ctx, worker)
	}

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, order ordertypes.Order) (sdk.Coin, error) {

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
	}

	for sp, shard := range order.Shards {
		workerName := fmt.Sprintf("%s-%s", amount.Denom, sp)
		worker, _ := k.GetWorker(ctx, workerName)
		incomePerSecond := amount.Amount.QuoInt64(int64(order.Replica)).QuoInt64(duration)
		reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
		worker.Reward.Amount = worker.Reward.Amount.Add(reward)
		worker.IncomePerSecond.Amount = worker.IncomePerSecond.Amount.Sub(incomePerSecond)
		worker.Storage -= shard.Size_
		worker.LastRewardAt = ctx.BlockHeight()
		k.SetWorker(ctx, worker)
	}

	return refundCoin, nil
}

func (k Keeper) Claim(ctx sdk.Context, denom string, sp string) error {

	workername := fmt.Sprintf("%s-%s", denom, sp)
	worker, found := k.GetWorker(ctx, workername)
	if !found {
		return status.Errorf(codes.NotFound, "not %s payment for worker %s found", denom, sp)
	}

	reward := worker.IncomePerSecond.Amount.MulInt64(ctx.BlockHeight() - worker.LastRewardAt)
	worker.Reward.Amount = worker.Reward.Amount.Add(reward)

	if worker.Debt.Amount.TruncateInt().IsZero() {
		return sdkerrors.Wrap(types.ErrInvalidAmount, "no reward")
	}

	rewardCoin := sdk.NewCoin(denom, worker.Reward.Amount.TruncateInt())

	spAcc := sdk.MustAccAddressFromBech32(sp)

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, spAcc, sdk.Coins{rewardCoin})

	if err != nil {
		return err
	}

	worker.Reward.Amount = worker.Reward.Amount.Sub(sdk.NewDecFromInt(rewardCoin.Amount))
	worker.LastRewardAt = ctx.BlockHeight()

	k.SetWorker(ctx, worker)

	return nil
}

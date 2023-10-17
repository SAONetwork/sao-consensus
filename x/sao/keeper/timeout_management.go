package keeper

import (
	markettypes "github.com/SaoNetwork/sao/x/market/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const MaxTries uint64 = 10

func (k Keeper) HandleTimeoutOrder(ctx sdk.Context, orderId uint64) {
	log := k.Logger(ctx)
	order, found := k.order.GetOrder(ctx, orderId)
	if !found {
		return
	}

	if order.Status == ordertypes.OrderPending {
		k.model.CancelOrder(ctx, orderId)
		return
	}

	if uint64(ctx.BlockHeight())+order.Timeout >= order.CreatedAt+order.Duration {
		return
	}

	var timeoutShards []ordertypes.Shard
	var uncompletedShards []uint64
	var completedShards []uint64
	var sps []string
	timeoutCount := 0

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			continue
		}
		sps = append(sps, shard.Sp)
		// TODO: migrating timeout
		if shard.Status == ordertypes.ShardWaiting {
			timeoutShards = append(timeoutShards, shard)
			timeoutCount++
		}
		if shard.Status == ordertypes.ShardCompleted {
			completedShards = append(completedShards, id)
		} else {
			uncompletedShards = append(uncompletedShards, id)
		}
	}

	// all shard completes
	if timeoutCount == 0 {
		for _, shardId := range uncompletedShards {
			k.order.RemoveShard(ctx, shardId)
		}
		if len(uncompletedShards) != 0 {
			order.Shards = completedShards
			k.order.SetOrder(ctx, order)
		}
		return
	}

	log.Debug("order timeout", "orderId", order.Id, "sps", sps)

	// TODO: sp punishment?
	randSp := k.node.RandomSP(ctx, timeoutCount, sps, int64(order.Size_))

	if len(randSp) == 0 {
		if uint64(ctx.BlockHeight())-order.CreatedAt > MaxTries*order.Timeout {
			if order.Status != ordertypes.OrderCompleted {
				// order timeout , remove shard and cancel order
				for _, shardId := range order.Shards {
					k.order.RemoveShard(ctx, shardId)
				}
				k.model.CancelOrder(ctx, orderId)
			} else {
				// order complete, decrease replica to calculate refund correctly
				for _, shardId := range uncompletedShards {
					k.order.RemoveShard(ctx, shardId)
				}
				order.Replica -= int32(timeoutCount)
				order.Shards = completedShards

				amount := sdk.NewDecCoinFromCoin(order.Amount)
				// refundDec = amount - price * size * replica * duration
				refundDec := amount.Amount.Sub(order.UnitPrice.Amount.MulInt64(int64(order.Size_)).MulInt64(int64(order.Replica)).MulInt64(int64(order.Duration)))
				refundCoin := sdk.NewCoin(amount.Denom, refundDec.TruncateInt())
				if !refundCoin.IsZero() {
					var payAddr sdk.AccAddress
					var err error
					if order.PaymentDid != "" {
						payAddr, err = k.did.GetCosmosPaymentAddress(ctx, order.Owner)
					} else {
						payAddr, err = k.did.GetCosmosPaymentAddress(ctx, order.Owner)
					}
					if err == nil {
						err = k.bank.SendCoinsFromModuleToAccount(ctx, markettypes.ModuleName, payAddr, sdk.Coins{refundCoin})
						if err != nil {
							log.Error("failed to refund", "err", err)
						}
					} else {
						log.Error("failed to refund", "err", err)
					}
					order.Amount = order.Amount.Sub(refundCoin)
				}
				k.order.SetOrder(ctx, order)
			}
			return
		}
	} else {

		for i, node := range randSp {
			shard := timeoutShards[i]
			log.Debug("fix shard", "shardId", shard.Id, "oldSP", shard.Sp, "newSp", node.Creator)

			shard.Status = ordertypes.ShardTimeout
			k.order.SetShard(ctx, shard)

			newShard := k.order.NewShardTask(ctx, &order, node.Creator)
			order.Shards = append(order.Shards, newShard.Id)
		}

		k.order.SetOrder(ctx, order)
	}

	k.SetTimeoutOrderBlock(ctx, order, uint64(ctx.BlockHeight())+order.Timeout)

	return
}

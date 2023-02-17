package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NewShardTask(ctx sdk.Context, order *types.Order, provider string) *types.Shard {

	var size int32
	if order.Size_ == 0 {
		size = 1
	} else {
		size = int32(order.Size_)
	}

	shard := &types.Shard{
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     order.Cid,
		//TODO: use the same type as Order.Size_
		Size_: size,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.ShardEventProvider, provider),
			sdk.NewAttribute(types.EventCid, shard.Cid),
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventOperation, fmt.Sprintf("%d", order.Operation)),
		),
	)
	return shard
}

func (k Keeper) FulfillShard(ctx sdk.Context, order *types.Order, sp string, cid string, _ int32) error {

	shard := order.Shards[sp]

	shard.Status = types.ShardCompleted
	shard.Cid = cid
	/*
		amount := order.Amount.Amount.QuoRaw(int64(order.Replica))
			shard.Amount = sdk.NewCoin(order.Amount.Denom, amount)
			shard.CreatedAt = uint64(ctx.BlockTime().Unix())
			shard.Duration = uint64(order.Duration)
	*/

	order.Shards[sp] = shard

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
		),
	)

	k.SetOrder(ctx, *order)
	return nil
}

func (k Keeper) TerminateShard(ctx sdk.Context, shard *types.Shard, sp string, owner string, orderId uint64) error {

	/*
		totalCost := shard.Amount.Amount.QuoRaw(int64(shard.Duration)).MulRaw(ctx.BlockTime().Unix() - int64(shard.CreatedAt))

		pending := shard.Amount.SubAmount(totalCost)

		shard.Status = types.ShardTerminated

		err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(owner), sdk.Coins{pending})

		if err != nil {
			return err
		}
	*/

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TerminateShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", orderId)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
			sdk.NewAttribute(types.EventCid, shard.Cid),
		),
	)
	return nil
}

func (k Keeper) RenewShard(ctx sdk.Context, order *types.Order, sp string) error {

	shard := order.Shards[sp]

	/*
		amount := order.Amount.Amount.QuoRaw(int64(order.Replica))
			shard.Amount = shard.Amount.Add(sdk.NewCoin(order.Amount.Denom, amount))
			shard.Duration += uint64(order.Duration)
	*/

	order.Shards[sp] = shard

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, sp),
		),
	)

	k.SetOrder(ctx, *order)
	return nil
}

func (k Keeper) ShardsPayment(ctx sdk.Context, orders []*types.Order, sp string) error {

	/*
		totalPending := sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)
		shards := make([]*types.Shard, 0)

		for _, order := range orders {
			shard := order.Shards[sp]
			totalCost := shard.Amount.Amount.QuoRaw(int64(shard.Duration)).MulRaw(ctx.BlockTime().Unix() - int64(shard.CreatedAt))
			pending := shard.Amount.SubAmount(totalCost)
			shard.Paid = shard.Paid.Add(pending)
			shards = append(shards, shard)
			totalPending = totalPending.Add(pending)

		}
		err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(sp), sdk.Coins{totalPending})

		if err != nil {
			return err
		}

		for idx, order := range orders {
			order.Shards[sp] = shards[idx]
			k.SetOrder(ctx, *order)
		}
	*/

	return nil
}

func (k Keeper) MigrateShard(ctx sdk.Context, order *types.Order, from string, to string) *types.Shard {

	shard := &types.Shard{
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     order.Cid,
		From:    from,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.ShardEventProvider, from),
			sdk.NewAttribute(types.EventCid, shard.Cid),
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventOperation, fmt.Sprintf("%d", order.Operation)),
		),
	)

	return shard
}

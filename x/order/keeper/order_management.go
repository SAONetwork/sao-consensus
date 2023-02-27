package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewOrder(ctx sdk.Context, order *types.Order, sps []string) (uint64, error) {

	paymentAcc, err := k.did.GetCosmosPaymentAddress(ctx, order.Owner)
	if err != nil {
		return 0, err
	}

	logger := k.Logger(ctx)

	logger.Error("try payment", "payer", paymentAcc, "amount", order.Amount)

	err = k.bank.SendCoinsFromAccountToModule(ctx, paymentAcc, types.ModuleName, sdk.Coins{order.Amount})
	if err != nil {
		return 0, err
	}

	order.Id = k.AppendOrder(ctx, *order)

	k.GenerateShards(ctx, order, sps)

	order.CreatedAt = uint64(ctx.BlockHeight())
	order.Metadata.CreatedAt = order.CreatedAt

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	expiredOrder, found := k.GetExpiredOrder(ctx, uint64(order.Expire))
	if found {
		expiredOrder.Data = append(expiredOrder.Data, order.Id)
	} else {
		expiredOrder = types.ExpiredOrder{
			Height: uint64(order.Expire),
			Data:   []uint64{order.Id},
		}
	}

	k.SetExpiredOrder(ctx, expiredOrder)

	k.SetOrder(ctx, *order)

	return order.Id, nil
}

func (k Keeper) GenerateShards(ctx sdk.Context, order *types.Order, sps []string) {

	if len(sps) > 0 {
		shards := make(map[string]*types.Shard, 0)
		for _, sp := range sps {
			shards[sp] = k.NewShardTask(ctx, order, sp)
		}

		order.Shards = shards

		order.Status = types.OrderDataReady

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderDataReadyEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute(types.OrderEventCreator, order.Creator),
				sdk.NewAttribute(types.OrderEventProvider, order.Provider),
				sdk.NewAttribute(types.EventCid, order.Cid),
			),
		)
	}
}

func (k Keeper) TerminateOrder(ctx sdk.Context, orderId uint64, refundCoin sdk.Coin) error {

	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return status.Errorf(codes.NotFound, "order %d not found", orderId)
	}

	if order.Status != types.OrderCompleted {
		return sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "invalid order stauts, expect complete")
	}

	//order.Status = types.OrderTerminated

	paymentAcc, err := k.did.GetCosmosPaymentAddress(ctx, order.Owner)
	if err != nil {
		return err
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, paymentAcc, sdk.Coins{refundCoin})

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TerminateOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	k.RemoveOrder(ctx, orderId)

	return nil
}

func (k Keeper) CancelOrder(ctx sdk.Context, orderId uint64) error {

	order, _ := k.GetOrder(ctx, orderId)

	if k.refundOrder(ctx, orderId) != nil {
		return sdkerrors.Wrapf(types.ErrorRefundOrder, "refund order failed")
	}

	//order.Status = types.OrderCanceled

	k.RemoveOrder(ctx, orderId)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CancelOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	return nil
}

func (k Keeper) refundOrder(ctx sdk.Context, orderId uint64) error {
	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return status.Errorf(codes.NotFound, "order %d not found", orderId)
	}
	paymentAcc, err := k.did.GetCosmosPaymentAddress(ctx, order.Owner)
	if err != nil {
		return err
	}
	return k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, paymentAcc, sdk.Coins{order.Amount})
}

func (k Keeper) RefundExpiredOrder(ctx sdk.Context, orderId uint64) error {

	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return status.Errorf(codes.NotFound, "order %d not found", orderId)
	}

	if order.Status == types.OrderCompleted {
		return sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "invalid order stauts")
	}

	//order.Status = types.OrderTerminated
	for sp, shard := range order.Shards {
		if shard.Status == types.ShardCompleted {
			err := k.node.OrderRelease(ctx, sdk.MustAccAddressFromBech32(sp), &order)
			if err != nil {
				return err
			}
		}
	}

	if k.refundOrder(ctx, orderId) != nil {
		return sdkerrors.Wrapf(types.ErrorRefundOrder, "refund order failed")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TerminateOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	k.RemoveOrder(ctx, orderId)

	return nil
}

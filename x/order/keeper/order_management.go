package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewOrder(ctx sdk.Context, order *types.Order) (uint64, error) {

	paymentAcc, err := k.did.GetCosmosPaymentAddress(ctx, order.Owner)
	if err != nil {
		return 0, err
	}

	logger := k.Logger(ctx)

	logger.Error("order payment", "payer", paymentAcc, "amount", order.Amount)

	err = k.bank.SendCoinsFromAccountToModule(ctx, paymentAcc, types.ModuleName, sdk.Coins{order.Amount})
	if err != nil {
		return 0, err
	}

	order.Id = k.AppendOrder(ctx, *order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	k.SetOrder(ctx, *order)

	return order.Id, nil
}

/*
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
}*/

func (k Keeper) TerminateOrder(ctx sdk.Context, orderId uint64) error {

	order, found := k.GetOrder(ctx, orderId)
	if !found {
		return status.Errorf(codes.NotFound, "order %d not found", orderId)
	}

	if order.Status != types.OrderCompleted {
		return sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "invalid order stauts, expect complete")
	}

	order.Status = types.OrderTerminated

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TerminateOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	k.SetOrder(ctx, order)

	return nil
}

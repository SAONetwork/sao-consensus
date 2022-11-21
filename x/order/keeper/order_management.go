package keeper

import (
	"fmt"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewOrder(ctx sdk.Context, order types.Order, sps []nodetypes.Node) (uint64, error) {

	// pay for order
	err := k.bank.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(order.Owner), types.ModuleName, sdk.Coins{order.Amount})
	if err != nil {
		return 0, err
	}

	order.Id = k.AppendOrder(ctx, order)

	k.GenerateShards(ctx, order, sps)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	k.SetOrder(ctx, order)
	return order.Id, nil
}

func (k Keeper) GenerateShards(ctx sdk.Context, order types.Order, sps []nodetypes.Node) {

	if len(sps) > 0 {
		shards := make(map[string]*types.Shard, 0)
		for _, sp := range sps {
			shards[sp.Creator] = k.NewShardTask(ctx, order, sp.Creator)
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
	k.SetOrder(ctx, order)
}

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

	for provider, shard := range order.Shards {
		err := k.TerminateShard(ctx, shard, provider, order.Owner, order.Id)
		if err != nil {
			return err
		}
		err = k.node.OrderRelease(ctx, sdk.AccAddress(provider), shard.Pledge)
		if err != nil {
			return err
		}
	}

	k.SetOrder(ctx, order)

	return nil
}

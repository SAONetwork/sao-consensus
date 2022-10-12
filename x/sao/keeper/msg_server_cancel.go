package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	order, found := k.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	signers := msg.GetSigners()

	if len(signers) != 1 || signers[0].String() != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrSignerAndCreator, "signer shoud equal to creator")
	}

	if order.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotCreator, "only order creator allowed")
	}

	if order.Status != types.OrderCompleted {
		return nil, sdkerrors.Wrapf(types.ErrOrderCompleted, "order %d already completed", msg.OrderId)
	}

	if order.Status == types.OrderCanceled {
		return nil, sdkerrors.Wrapf(types.ErrOrderCanceled, "order %d already canceld", msg.OrderId)
	}

	order.Status = types.OrderCanceled

	k.SetOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CancelOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	return &types.MsgCancelResponse{}, nil
}
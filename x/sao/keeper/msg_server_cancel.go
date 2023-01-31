package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Cancel(goCtx context.Context, msg *types.MsgCancel) (*types.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
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

	k.order.CancelOrder(ctx, msg.OrderId)

	return &types.MsgCancelResponse{}, nil
}

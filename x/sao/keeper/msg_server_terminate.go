package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Terminate(goCtx context.Context, msg *types.MsgTerminate) (*types.MsgTerminateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	if order.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotCreator, "only order creator allowed")
	}

	if order.Metadata != nil && order.Metadata.DataId != "" {
		err := k.model.DeleteMeta(ctx, order.Metadata.DataId)
		if err != nil {
			return nil, err
		}
	}

	err := k.order.TerminateOrder(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	k.market.Withdraw(ctx, order)

	return &types.MsgTerminateResponse{}, nil
}

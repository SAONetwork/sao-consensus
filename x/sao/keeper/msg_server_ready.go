package keeper

import (
	"context"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Ready(goCtx context.Context, msg *types.MsgReady) (*types.MsgReadyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	if msg.Creator != order.Provider {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "")
	}

	if order.Status != types.OrderPending {
		return nil, sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "expect pending order")
	}

	var sps []nodetypes.Node

	if order.Operation == 0 {
		sps = k.node.RandomSP(ctx, order)
	} else if order.Operation == 1 {
		sps = k.FindSPByDataId(ctx, order.Metadata.DataId)
	}

	k.order.GenerateShards(ctx, &order, sps)

	k.order.SetOrder(ctx, order)

	return &types.MsgReadyResponse{}, nil
}

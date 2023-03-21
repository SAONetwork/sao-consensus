package keeper

import (
	"context"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

	if order.Status == ordertypes.OrderCompleted || order.Status == ordertypes.OrderMigrating {
		return nil, sdkerrors.Wrapf(types.ErrOrderCompleted, "order %d already completed", msg.OrderId)
	}

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			return nil, status.Errorf(codes.NotFound, "shard %d not found", id)
		}
		if shard.Status == ordertypes.ShardCompleted {
			err := k.node.OrderRelease(ctx, sdk.MustAccAddressFromBech32(shard.Sp), &order)
			if err != nil {
				return nil, err
			}
		}
	}

	err := k.order.CancelOrder(ctx, msg.OrderId)
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelResponse{}, nil
}

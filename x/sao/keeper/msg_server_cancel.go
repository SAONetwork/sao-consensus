package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ordertypes "github.com/SaoNetwork/sao/x/order/types"
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

	isCreator := false
	if order.Creator == msg.Creator {
		isCreator = true
	} else {
		node, found := k.node.GetNode(ctx, msg.Provider)
		if found {
			for _, address := range node.TxAddresses {
				if order.Creator == address {
					isCreator = true
				}
			}
		}
	}

	if !isCreator {
		return nil, sdkerrors.Wrapf(types.ErrNotCreator, "only order creator allowed")
	}

	if order.Status == ordertypes.OrderCompleted || order.Status == ordertypes.OrderMigrating {
		return nil, sdkerrors.Wrapf(types.ErrOrderCompleted, "order %d already completed", msg.OrderId)
	}

	isProvider := false
	if msg.Provider == msg.Creator {
		isProvider = true
	} else {
		provider, found := k.node.GetNode(ctx, msg.Provider)
		if found {
			for _, address := range provider.TxAddresses {
				if address == msg.Creator {
					isProvider = true
				}
			}
		}
	}

	if !isProvider {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, msg.Provider: %s", msg.Creator, msg.Provider)
	}

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			return nil, status.Errorf(codes.NotFound, "shard %d not found", id)
		}
		if shard.Status == ordertypes.ShardCompleted {
			err := k.node.ShardRelease(ctx, sdk.MustAccAddressFromBech32(shard.Sp), &shard)
			if err != nil {
				return nil, err
			}
		}
		k.order.RemoveShard(ctx, id)
	}

	err := k.model.CancelOrder(ctx, msg.OrderId)
	if err != nil {
		return nil, err
	}

	return &types.MsgCancelResponse{}, nil
}

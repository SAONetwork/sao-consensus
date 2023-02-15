package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Reject(goCtx context.Context, msg *types.MsgReject) (*types.MsgRejectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	shards := k.node.GetMetadataShards(ctx, order.Metadata.DataId, int(order.Replica))
	if _, ok := shards[msg.Creator]; !ok {
		return nil, sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider")
	}

	shard := shards[msg.Creator]

	if shard.Status != types.ShardWaiting {
		return nil, sdkerrors.Wrapf(types.ErrShardUnexpectedStatus, "invalid shard status: expect pending")
	}

	shard.Status = types.ShardRejected

	order.Status = types.OrderUnexpected

	k.Keeper.order.SetOrder(ctx, order)
	k.Keeper.node.SetShard(ctx, *shard)

	k.node.DecreaseReputation(ctx, msg.Creator, 1000)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.RejectShardEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, msg.Creator),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.OrderUnexpectedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	return &types.MsgRejectResponse{}, nil
}

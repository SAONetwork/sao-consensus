package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
)

func (k msgServer) Complete(goCtx context.Context, msg *types.MsgComplete) (*types.MsgCompleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	if order.Status != types.OrderDataReady && order.Status != types.OrderInProgress {
		return nil, sdkerrors.Wrapf(types.ErrOrderComplete, "order not waiting completed")
	}

	if _, ok := order.Shards[msg.Creator]; !ok {
		return nil, sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider")
	}

	shard := order.Shards[msg.Creator]

	if shard.Status == types.ShardCompleted {
		return nil, sdkerrors.Wrapf(types.ErrShardCompleted, "%s already completed the shard task in order %d", msg.Creator, order.Id)
	}

	if shard.Status != types.ShardWaiting {
		return nil, sdkerrors.Wrapf(types.ErrShardUnexpectedStatus, "invalid shard status, expect: wating")
	}

	logger := k.Logger(ctx)

	// check cid
	_, err := cid.Decode(msg.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invali cid: %s", msg.Cid)
	}

	// active shard
	err = k.order.FulfillShard(ctx, &order, msg.Creator, msg.Cid, msg.Size_)

	order.Status = types.OrderCompleted

	// set order status
	for _, shard := range order.Shards {
		if shard.Status != types.ShardCompleted {
			order.Status = types.OrderInProgress
		}
	}

	shard = order.Shards[msg.Creator]

	k.node.IncreaseReputation(ctx, msg.Creator, float32(shard.Amount.Amount.Int64()))

	k.node.OrderPledge(ctx, sdk.AccAddress(msg.Creator), shard.Amount)

	order.Shards[msg.Creator].Pledge = shard.Amount

	if order.Status == types.OrderCompleted {

		if order.Metadata != nil {
			_, found_metadata := k.model.GetMetadata(ctx, order.Metadata.DataId)

			if found_metadata {
				err = k.Keeper.model.UpdateMeta(ctx, order)
				if err != nil {
					logger.Error("failed to update metadata", "err", err.Error())
					return &types.MsgCompleteResponse{}, err
				}
			} else {
				err = k.Keeper.model.NewMeta(ctx, order)
				if err != nil {
					logger.Error("failed to store metadata", "err", err.Error())
					return &types.MsgCompleteResponse{}, err
				}
			}
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderCompletedEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			),
		)
	}

	k.order.SetOrder(ctx, order)

	return &types.MsgCompleteResponse{}, nil
}

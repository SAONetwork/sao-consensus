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

	var err error = nil
	var orderId uint64 = 0
	defer emitEvent(ctx, &orderId, &err)

	if msg.Size_ == 0 {
		err = sdkerrors.Wrapf(types.ErrorInvalidShardSize, "order %d shard %s: invalid shard size %d", msg.OrderId, msg.Cid, msg.Size_)
		return &types.MsgCompleteResponse{}, err
	}

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		err = sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
		return &types.MsgCompleteResponse{}, err
	}
	orderId = order.Id

	if order.Status != types.OrderDataReady && order.Status != types.OrderInProgress {
		err = sdkerrors.Wrapf(types.ErrOrderComplete, "order not waiting completed")
		return &types.MsgCompleteResponse{}, err
	}

	if _, ok := order.Shards[msg.Creator]; !ok {
		err = sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider")
		return &types.MsgCompleteResponse{}, err
	}

	shard := order.Shards[msg.Creator]

	if shard.Status == types.ShardCompleted {
		err = sdkerrors.Wrapf(types.ErrShardCompleted, "%s already completed the shard task in order %d", msg.Creator, order.Id)
		return &types.MsgCompleteResponse{}, err
	}

	if shard.Status != types.ShardWaiting {
		err = sdkerrors.Wrapf(types.ErrShardUnexpectedStatus, "invalid shard status, expect: wating")
		return &types.MsgCompleteResponse{}, err
	}

	logger := k.Logger(ctx)

	// check cid
	_, err = cid.Decode(msg.Cid)
	if err != nil {
		err = sdkerrors.Wrapf(types.ErrInvalidCid, "invali cid: %s", msg.Cid)
		return &types.MsgCompleteResponse{}, err
	}

	// active shard
	k.order.FulfillShard(ctx, &order, msg.Creator, msg.Cid, msg.Size_)

	order.Status = types.OrderCompleted

	// set order status
	for _, shard := range order.Shards {
		if shard.Status != types.ShardCompleted {
			order.Status = types.OrderInProgress
		}
	}

	// shard = order.Shards[msg.Creator]

	err = k.node.OrderPledge(ctx, msg.GetSigners()[0], &order)
	if err != nil {
		err = sdkerrors.Wrap(types.ErrorOrderPledgeFailed, err.Error())
		return &types.MsgCompleteResponse{}, err
	}

	amount := sdk.NewCoin(order.Amount.Denom, order.Amount.Amount.QuoRaw(int64(order.Replica)))
	k.node.IncreaseReputation(ctx, msg.Creator, float32(amount.Amount.Int64()))

	if order.Status == types.OrderCompleted {

		if order.Metadata != nil {
			_, foundMeta := k.model.GetMetadata(ctx, order.Metadata.DataId)

			if foundMeta {
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

		k.market.Deposit(ctx, order)

	}

	k.order.SetOrder(ctx, order)

	return &types.MsgCompleteResponse{}, err
}

func emitEvent(ctx sdk.Context, orderId *uint64, err *error) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.OrderCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", *orderId)),
			sdk.NewAttribute(types.EventErrorInfo, (*err).Error()),
		),
	)
}

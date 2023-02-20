package keeper

import (
	"context"
	"fmt"
	"strings"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
)

func (k msgServer) CompleteShard(goCtx context.Context, msg *types.MsgCompleteShard) (*types.MsgCompleteShardResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	var err error = nil
	var orderId uint64 = 0
	defer emitEvent(ctx, &orderId, &err)
	resp := &types.MsgCompleteShardResponse{}

	if msg.Size_ == 0 {
		err = sdkerrors.Wrapf(types.ErrorInvalidShardSize, "shard %d: invalid shard size %d", msg.Idx, msg.Size_)
		return resp, err
	}

	shard, found := k.node.GetShard(ctx, msg.Idx)

	if !found {
		err = sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider")
		return resp, err
	}

	if shard.Node != msg.Creator {
		err = sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider")
		return resp, err
	}

	if shard.Status == types.ShardCompleted {
		err = sdkerrors.Wrapf(types.ErrShardCompleted, "shared %s already completed", msg.Idx)
		return resp, err
	}

	if shard.Status != types.ShardWaiting {
		err = sdkerrors.Wrapf(types.ErrShardUnexpectedStatus, "invalid shard status, expect: wating")
		return resp, err
	}

	// check shard size
	if msg.Size_ != int32(shard.Size_) {
		err = sdkerrors.Wrapf(types.ErrorInvalidShardSize, "shard %s: invalid shard size %d, expect %d", msg.Idx, msg.Cid, msg.Size_, shard.Size_)
		return resp, err
	}

	// check cid
	_, err = cid.Decode(msg.Cid)
	if err != nil {
		err = sdkerrors.Wrapf(types.ErrInvalidCid, "invali cid: %s", msg.Cid)
		return resp, err
	}

	metadata, found := k.model.GetMetadata(ctx, shard.DataId)
	if !found {
		err = sdkerrors.Wrapf(types.ErrorMetadataNotFound, "metadata %d not found", shard.DataId)
		return resp, err
	}

	order, found := k.order.GetOrder(ctx, metadata.OrderId)
	if !found {
		err = sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", metadata.OrderId)
		return resp, err
	}

	// active shard && pledge
	err = k.node.ActiveShard(ctx, &order, &shard, msg.Cid, uint64(msg.Size_))

	if err != nil {
		err = sdkerrors.Wrap(types.ErrorOrderPledgeFailed, err.Error())
		return resp, err
	}

	orderId = order.Id

	// check order status
	shards := k.node.GetMetadataShards(ctx, order.Metadata.DataId, int(order.Replica))

	completeShardSCount := 0
	for _, shard := range shards {
		if shard.Status == types.ShardCompleted {
			completeShardSCount += 1
		}
	}

	if completeShardSCount == int(order.Replica) {
		order.Status = types.OrderCompleted
	}

	// add node reputation
	amount := sdk.NewCoin(order.Amount.Denom, order.Amount.Amount.QuoRaw(int64(order.Replica)))
	k.node.IncreaseReputation(ctx, msg.Creator, float32(amount.Amount.Int64()))

	// avoid version conflicts
	meta, isFound := k.model.GetMetadata(ctx, order.Metadata.DataId)
	if isFound && order.Status == types.OrderCompleted {
		if meta.OrderId > orderId {
			// report error if order id is less than the latest version
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidCommitId, "invalid commitId: %s, detected version conficts with order: %d", order.Metadata.Commit, meta.OrderId)
		}

		lastOrder, isFound := k.order.GetOrder(ctx, meta.OrderId)
		if isFound {
			if lastOrder.Status == ordertypes.OrderPending || lastOrder.Status == ordertypes.OrderInProgress || lastOrder.Status == ordertypes.OrderDataReady {
				return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "unexpected last order: %s, status: %d", meta.OrderId, lastOrder.Status)
			}
		} else {
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "invalid last order: %s", meta.OrderId)
		}

		if strings.Contains(order.Metadata.Commit, "|") {
			lastCommitId := strings.Split(order.Metadata.Commit, "|")[0]
			commitId := strings.Split(order.Metadata.Commit, "|")[1]

			if !strings.Contains(meta.Commit, lastCommitId) {
				// report error if base version is not the latest version
				return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidCommitId, "invalid commitId: %s, detected version conficts, should be %s", lastCommitId, meta.Commit[:36])
			}
			order.Metadata.Commit = commitId
		}
	}

	if order.Status == types.OrderCompleted {

		if order.Metadata != nil {
			_, foundMeta := k.model.GetMetadata(ctx, order.Metadata.DataId)

			if foundMeta {
				err = k.Keeper.model.UpdateMeta(ctx, order)
				if err != nil {
					logger.Error("failed to update metadata", "err", err.Error())
					return resp, err
				}
			} else {
				err = k.Keeper.model.NewMeta(ctx, order)
				if err != nil {
					logger.Error("failed to store metadata", "err", err.Error())
					return resp, err
				}
			}
		}

		err := k.market.Deposit(ctx, order)
		if err != nil {
			return nil, sdkerrors.Wrap(ordertypes.ErrorOrderPayment, err.Error())
		}
	}

	if shard.From != "" {
		sp := sdk.MustAccAddressFromBech32(shard.From)
		err := k.node.OrderRelease(ctx, sp, &order)
		if err != nil {
			return nil, err
		}
		oldShard := k.node.GetMetadataShardByNode(ctx, shard.DataId, shard.From, int(order.Replica))
		oldIdx := fmt.Sprintf("%s-%d", shard.DataId, oldShard.Idx)
		shard.Idx = oldIdx
		k.node.SetShard(ctx, shard)
		k.node.RemoveShard(ctx, msg.Idx)
	}

	k.order.SetOrder(ctx, order)

	return &types.MsgCompleteShardResponse{}, nil
}

package keeper

import (
	"context"
	"fmt"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	shard := k.order.GetOrderShardBySP(ctx, &order, msg.Provider)

	if shard == nil {
		err = sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider", msg.Provider)
		return &types.MsgCompleteResponse{}, err
	}

	if shard.Status == ordertypes.ShardCompleted {
		err = sdkerrors.Wrapf(types.ErrShardCompleted, "%s already completed the shard task in order %d", msg.Provider, order.Id)
		return &types.MsgCompleteResponse{}, err
	}

	if shard.Status != ordertypes.ShardWaiting && shard.Status != ordertypes.ShardMigrating {
		err = sdkerrors.Wrapf(types.ErrShardUnexpectedStatus, "invalid shard status, expect: waiting/migrating")
		return &types.MsgCompleteResponse{}, err
	}

	// check shard size
	if msg.Size_ != shard.Size_ {
		err = sdkerrors.Wrapf(types.ErrorInvalidShardSize, "order %d shard %s: invalid shard size %d, expect %d", msg.OrderId, msg.Cid, msg.Size_, shard.Size_)
		return &types.MsgCompleteResponse{}, err
	}

	// avoid version conflicts
	meta, isFoundMeta := k.model.GetMetadata(ctx, order.DataId)
	if !isFoundMeta {
		return nil, status.Errorf(codes.NotFound, "metadata %s not found", order.DataId)
	}

	if meta.Status != modeltypes.MetaNew && meta.Status != modeltypes.MetaComplete && meta.Status != int32(order.Operation) {
		err = sdkerrors.Wrapf(types.ErrorInvalidOperation, "meta.Status %d is different to order.Operation %d", meta.Status, order.Operation)
		return nil, err
	}

	if len(meta.Orders) != 0 {
		lastOrder, isFound := k.order.GetOrder(ctx, meta.Orders[len(meta.Orders)-1])
		if isFound {
			if lastOrder.Status == ordertypes.OrderPending || lastOrder.Status == ordertypes.OrderInProgress || lastOrder.Status == ordertypes.OrderDataReady {
				return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "unexpected last order: %s, status: %d", meta.OrderId, lastOrder.Status)
			}
		} else {
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "invalid last order: %s", meta.OrderId)
		}
	}

	logger := k.Logger(ctx)

	// check cid
	_, err = cid.Decode(msg.Cid)
	if err != nil {
		err = sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", msg.Cid)
		return &types.MsgCompleteResponse{}, err
	}

	orderInProgress := order

	if shard.Status == ordertypes.ShardMigrating {
		// shard migrate
		if shard.From == "" {
			err = sdkerrors.Wrapf(types.ErrorEmptyShardFrom, "invalid cid: %s", msg.Cid)
			return &types.MsgCompleteResponse{}, err
		}
		sp := sdk.MustAccAddressFromBech32(shard.From)
		oldShard := k.order.GetOrderShardBySP(ctx, &order, shard.From)
		err = k.node.ShardRelease(ctx, sp, oldShard)
		if err != nil {
			return nil, err
		}
		orderList := []*ordertypes.Order{&order}
		if oldShard.OrderId != order.Id {
			// The order in progress is the one corresponding to the orderId field in oldShard,
			// which is used to correctly calculate the next Migrate and ShardPledge
			orderInProgress, _ = k.order.GetOrder(ctx, oldShard.OrderId)
			orderList = append(orderList, &orderInProgress)
		}
		shard.OrderId = oldShard.OrderId
		shard.RenewInfos = oldShard.RenewInfos
		shard.CreatedAt = uint64(ctx.BlockHeight())
		shard.Duration = oldShard.CreatedAt + oldShard.Duration - shard.CreatedAt
		err = k.market.Migrate(ctx, orderInProgress, *oldShard, *shard)
		if err != nil {
			return nil, err
		}
		k.order.RemoveShard(ctx, oldShard.Id)
		if len(oldShard.RenewInfos) > 1 {
			for i := 0; i < len(oldShard.RenewInfos)-1; i++ {
				order, _ := k.order.GetOrder(ctx, oldShard.RenewInfos[i].OrderId)
				orderList = append(orderList, &order)
			}
		}
		for i, order := range orderList {
			newShards := make([]uint64, 0)
			for _, id := range order.Shards {
				if id != oldShard.Id {
					newShards = append(newShards, id)
				}
			}
			// first order has set new shard in shards in migrate
			if i > 0 {
				newShards = append(newShards, shard.Id)
			}
			order.Shards = newShards
			k.order.SetOrder(ctx, *order)
		}
	} else {
		shard.CreatedAt = uint64(ctx.BlockHeight())
		shard.Duration = order.Duration
		k.market.WorkerAppend(ctx, &order, shard)
		if order.Status != ordertypes.OrderCompleted {
			// order complete
			err = k.Keeper.model.UpdateMeta(ctx, order)
			if err != nil {
				logger.Error("failed to update metadata", "err", err.Error())
				return nil, err
			}

			err = k.market.Deposit(ctx, order)
			if err != nil {
				return nil, err
			}

			order.Status = ordertypes.OrderCompleted
		}
	}

	// active shard
	k.order.FulfillShard(ctx, shard, msg.Provider, msg.Cid)
	k.SetExpiredShardBlock(ctx, shard.Id, shard.CreatedAt+shard.Duration)
	k.model.ExtendMetaDuration(ctx, meta.DataId, shard.CreatedAt+shard.Duration)

	// shard = order.Shards[msg.Provider]

	err = k.node.ShardPledge(ctx, shard, orderInProgress.UnitPrice)
	if err != nil {
		err = sdkerrors.Wrap(types.ErrorOrderPledgeFailed, err.Error())
		return &types.MsgCompleteResponse{}, err
	}

	amount := sdk.NewCoin(order.Amount.Denom, order.Amount.Amount.QuoRaw(int64(order.Replica)))
	k.node.IncreaseReputation(ctx, msg.Provider, float32(amount.Amount.Int64()))

	k.order.SetOrder(ctx, order)

	return &types.MsgCompleteResponse{}, err
}

func emitEvent(ctx sdk.Context, orderId *uint64, err *error) {
	if (*err) != nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderCompletedEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", *orderId)),
				sdk.NewAttribute(types.EventErrorInfo, (*err).Error()),
			),
		)
	} else {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderCompletedEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", *orderId)),
			),
		)
	}
}

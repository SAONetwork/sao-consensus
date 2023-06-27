package keeper

import (
	"context"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Ready(goCtx context.Context, msg *types.MsgReady) (*types.MsgReadyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	isProvider := false
	if order.Provider == msg.Creator && msg.Provider == msg.Creator {
		isProvider = true
	} else if order.Provider == msg.Provider {
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
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, msg.Provider: %s", msg.Creator, order.Provider)
	}

	if order.Status != ordertypes.OrderPending {
		return nil, sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "expect pending order")
	}

	var sps []nodetypes.Node
	var err error

	sps, err = k.GetSps(ctx, order, order.DataId)
	if err != nil {
		return nil, err
	}

	spAddresses := make([]string, 0)
	for _, sp := range sps {

		spAddresses = append(spAddresses, sp.Creator)
	}

	k.order.GenerateShards(ctx, &order, spAddresses)

	k.order.SetOrder(ctx, order)

	k.SetTimeoutOrderBlock(ctx, order, uint64(ctx.BlockHeight())+order.Timeout)

	shards := make([]*types.ShardMeta, 0)
	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			return nil, status.Errorf(codes.NotFound, "shard %d not found", id)
		}
		node, node_found := k.node.GetNode(ctx, shard.Sp)
		if !node_found {
			continue
		}
		meta := types.ShardMeta{
			ShardId:  shard.Id,
			Peer:     node.Peer,
			Cid:      shard.Cid,
			Provider: order.Provider,
			Sp:       shard.Sp,
		}
		shards = append(shards, &meta)
	}

	return &types.MsgReadyResponse{
		OrderId: order.Id,
		Shards:  shards,
	}, nil
}

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

	logger := k.Logger(ctx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	if msg.Creator != order.Provider {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, order.Provider: %s", msg.Creator, order.Provider)
	}

	if order.Status != types.OrderPending {
		return nil, sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "expect pending order")
	}

	var sps []nodetypes.Node

	if order.Operation == 1 {
		sps = k.node.RandomSP(ctx, int(order.Replica))
	} else if order.Operation == 2 {
		sps = k.FindSPByDataId(ctx, order.Metadata.DataId)
	}

	sps_addr := make([]string, 0)
	for _, sp := range sps {

		logger.Error("sp sp.String ###################", "sp.String", sp.String())

		sps_addr = append(sps_addr, sp.Creator)
	}
	k.order.GenerateShards(ctx, &order, sps_addr)

	k.order.SetOrder(ctx, order)

	shards := make(map[string]*types.ShardMeta, 0)
	for p, shard := range order.Shards {
		node, node_found := k.node.GetNode(ctx, p)
		if !node_found {
			continue
		}
		meta := types.ShardMeta{
			ShardId:  shard.Id,
			Peer:     node.Peer,
			Cid:      shard.Cid,
			Provider: order.Provider,
		}
		shards[p] = &meta
	}

	return &types.MsgReadyResponse{
		OrderId: order.Id,
		Shards:  shards,
	}, nil
}

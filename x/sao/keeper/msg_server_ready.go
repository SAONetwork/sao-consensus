package keeper

import (
	"context"

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
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, order.Provider: %s", msg.Creator, order.Provider)
	}

	if order.Status != types.OrderPending {
		return nil, sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "expect pending order")
	}

	shards := k.node.NewShards(ctx, &order)

	k.order.SetOrder(ctx, order)

	_shards := make(map[string]*types.ShardMeta, 0)
	for _, shard := range shards {
		node, node_found := k.node.GetNode(ctx, shard.Node)
		if !node_found {
			continue
		}
		meta := types.ShardMeta{
			Idx:      shard.Idx,
			Peer:     node.Peer,
			Cid:      shard.Cid,
			Provider: order.Provider,
		}
		_shards[shard.Node] = &meta
	}

	return &types.MsgReadyResponse{
		OrderId: order.Id,
		Shards:  _shards,
	}, nil
}

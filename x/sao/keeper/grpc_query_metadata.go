package keeper

import (
	"context"
	"fmt"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Metadata(goCtx context.Context, req *types.QueryMetadataRequest) (*types.QueryMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &req.Proposal

	var dataId string
	if proposal.KeywordType > 1 {
		model, isFound := k.model.GetModel(ctx, fmt.Sprintf("%s-%s-%s",
			proposal.DataOwner, proposal.Keyword, proposal.GroupId,
		))
		if !isFound {
			return nil, status.Errorf(codes.NotFound, "dataId not found by Alias: %s", proposal.Keyword)
		}
		dataId = model.Data
	} else {
		dataId = proposal.Keyword
	}

	meta, isFound := k.model.GetMetadata(ctx, dataId)
	if !isFound {
		return nil, status.Errorf(codes.NotFound, "dataId:%s not found", dataId)
	}

	order, found := k.order.GetOrder(ctx, meta.OrderId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "order:%d not found", meta.OrderId)
	}

	metadata := types.Metadata{
		DataId:     meta.DataId,
		Owner:      meta.Owner,
		Alias:      meta.Alias,
		GroupId:    meta.GroupId,
		OrderId:    meta.OrderId,
		Tags:       meta.Tags,
		Cid:        meta.Cid,
		Commits:    meta.Commits,
		ExtendInfo: meta.ExtendInfo,
		Update:     meta.Update,
		Commit:     meta.Commit,
		Rule:       meta.Rule,
		Duration:   meta.Duration,
		CreatedAt:  meta.CreatedAt,
		Provider:   order.Provider,
		Expire:     int32(order.CreatedAt + order.Timeout),
		Status:     order.Status,
		Replica:    order.Replica,
		Amount:     order.Amount,
		Size_:      order.Size_,
		Operation:  order.Operation,
	}

	shards := make(map[string]*types.ShardMeta, 0)
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
			Provider: shard.Sp,
		}
		shards[shard.Sp] = &meta
	}

	return &types.QueryMetadataResponse{
		Metadata: metadata,
		Shards:   shards,
	}, nil
}

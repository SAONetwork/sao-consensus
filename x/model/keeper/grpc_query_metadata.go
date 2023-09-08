package keeper

import (
	"context"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MetadataAll(c context.Context, req *types.QueryAllMetadataRequest) (*types.QueryAllMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var metadatas []types.Metadata
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	metadataStore := prefix.NewStore(store, types.KeyPrefix(types.MetadataKeyPrefix))

	pageRes, err := query.Paginate(metadataStore, req.Pagination, func(key []byte, value []byte) error {
		var metadata types.Metadata
		if err := k.cdc.Unmarshal(value, &metadata); err != nil {
			return err
		}

		metadatas = append(metadatas, metadata)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMetadataResponse{Metadata: metadatas, Pagination: pageRes}, nil
}

func (k Keeper) Metadata(c context.Context, req *types.QueryGetMetadataRequest) (*types.QueryGetMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMetadata(
		ctx,
		req.DataId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	orderId := val.OrderId

	if orderId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid orderId")
	}

	order, found := k.order.GetOrder(ctx, orderId)

	if !found {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	shard_metas := make(map[string]*types.ShardMeta, 0)

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			return nil, status.Errorf(codes.NotFound, "shard %d not found", id)
		}
		if shard.Status != ordertypes.ShardCompleted {
			continue
		}

		node, node_found := k.node.GetNode(ctx, shard.Sp)
		if !node_found {
			continue
		}

		meta := types.ShardMeta{
			ShardId: shard.Id,
			Peer:    node.Peer,
			Cid:     shard.Cid,
		}
		shard_metas[shard.Sp] = &meta
	}

	return &types.QueryGetMetadataResponse{
		Metadata: val,
		OrderId:  orderId,
		Shards:   shard_metas,
	}, nil
}

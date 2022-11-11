package keeper

import (
	"context"

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

	logger := k.Logger(ctx)

	orderId := val.OrderId

	if orderId < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid orderId")
	}

	logger.Info("#########", "orderId", orderId, "sao", k.sao)

	order, found := k.sao.GetOrder(ctx, orderId)

	logger.Debug("#########", "found", found)

	if !found {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	shard_metas := make(map[string]*types.ShardMeta, 0)

	for p, shard := range order.Shards {
		node, node_found := k.node.GetNode(ctx, p)
		if !node_found {
			continue
		}
		meta := types.ShardMeta{
			ShardId: shard.Id,
			Peer:    node.Peer,
			Cid:     shard.Cid,
		}
		shard_metas[p] = &meta
	}

	logger.Debug("#########", "order", order)

	return &types.QueryGetMetadataResponse{
		Metadata: val,
		OrderId:  orderId,
		Shards:   shard_metas,
	}, nil
}

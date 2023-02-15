package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShardAll(c context.Context, req *types.QueryAllShardRequest) (*types.QueryAllShardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var shards []types.Shard
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	shardStore := prefix.NewStore(store, types.KeyPrefix(types.ShardKeyPrefix))

	pageRes, err := query.Paginate(shardStore, req.Pagination, func(key []byte, value []byte) error {
		var shard types.Shard
		if err := k.cdc.Unmarshal(value, &shard); err != nil {
			return err
		}

		shards = append(shards, shard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllShardResponse{Shard: shards, Pagination: pageRes}, nil
}

func (k Keeper) Shard(c context.Context, req *types.QueryGetShardRequest) (*types.QueryGetShardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetShard(
		ctx,
		req.Idx,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetShardResponse{Shard: val}, nil
}

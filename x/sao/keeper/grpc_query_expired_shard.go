package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ExpiredShardAll(c context.Context, req *types.QueryAllExpiredShardRequest) (*types.QueryAllExpiredShardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var expiredShards []types.ExpiredShard
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	expiredShardStore := prefix.NewStore(store, types.KeyPrefix(types.ExpiredShardKeyPrefix))

	pageRes, err := query.Paginate(expiredShardStore, req.Pagination, func(key []byte, value []byte) error {
		var expiredShard types.ExpiredShard
		if err := k.cdc.Unmarshal(value, &expiredShard); err != nil {
			return err
		}

		expiredShards = append(expiredShards, expiredShard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExpiredShardResponse{ExpiredShard: expiredShards, Pagination: pageRes}, nil
}

func (k Keeper) ExpiredShard(c context.Context, req *types.QueryGetExpiredShardRequest) (*types.QueryGetExpiredShardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExpiredShard(
		ctx,
		req.Height,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExpiredShardResponse{ExpiredShard: val}, nil
}

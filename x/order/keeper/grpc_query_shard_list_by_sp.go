package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShardListBySp(goCtx context.Context, req *types.QueryShardListBySpRequest) (*types.QueryShardListBySpResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	nextShardId := k.GetShardCount(ctx)
	if nextShardId == req.ShardId {
		return &types.QueryShardListBySpResponse{Shard: nil, NextShardId: nextShardId}, nil
	}

	shards := k.GetAllShardWithIdAndSp(ctx, req.ShardId, req.Sp)

	return &types.QueryShardListBySpResponse{Shard: shards, NextShardId: nextShardId}, nil
}

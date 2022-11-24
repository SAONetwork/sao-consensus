package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PastSeedsAll(c context.Context, req *types.QueryAllPastSeedsRequest) (*types.QueryAllPastSeedsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pastSeedss []types.PastSeeds
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pastSeedsStore := prefix.NewStore(store, types.KeyPrefix(types.PastSeedsKeyPrefix))

	pageRes, err := query.Paginate(pastSeedsStore, req.Pagination, func(key []byte, value []byte) error {
		var pastSeeds types.PastSeeds
		if err := k.cdc.Unmarshal(value, &pastSeeds); err != nil {
			return err
		}

		pastSeedss = append(pastSeedss, pastSeeds)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPastSeedsResponse{PastSeeds: pastSeedss, Pagination: pageRes}, nil
}

func (k Keeper) PastSeeds(c context.Context, req *types.QueryGetPastSeedsRequest) (*types.QueryGetPastSeedsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPastSeeds(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPastSeedsResponse{PastSeeds: val}, nil
}

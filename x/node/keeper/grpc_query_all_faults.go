package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllFaults(goCtx context.Context, req *types.QueryAllFaultsRequest) (*types.QueryAllFaultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// _ := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryAllFaultsResponse{
		Faults: nil,
	}, nil
}

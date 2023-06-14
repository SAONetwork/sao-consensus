package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NetVersion(goCtx context.Context, req *types.QueryNetVersionRequest) (*types.QueryNetVersionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryNetVersionResponse{
		Version: "v1.5.0",
	}, nil
}

package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/model/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MetaStatus(goCtx context.Context, req *types.QueryMetaStatusRequest) (*types.QueryMetaStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	status := make([]int32, len(req.DataIds))
	for i, id := range req.DataIds {
		meta, found := k.GetMetadata(ctx, id)
		if !found {
			status[i] = -1
		} else {
			status[i] = meta.Status
		}
	}

	return &types.QueryMetaStatusResponse{}, nil
}

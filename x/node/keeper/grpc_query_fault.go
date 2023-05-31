package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Fault(goCtx context.Context, req *types.QueryFaultRequest) (*types.QueryFaultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	fault, found := k.GetFault(ctx, req.FaultId)
	if found {
		return nil, status.Error(codes.NotFound, "fault not found")
	}

	return &types.QueryFaultResponse{
		Fault: fault,
	}, nil
}

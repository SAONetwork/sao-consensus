package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Latesthight(goCtx context.Context, req *types.QueryLatesthightRequest) (*types.QueryLatesthightResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryLatesthightResponse{
		LatestBlockHeight: uint64(ctx.BlockHeight()),
		LatestBlockTime:   ctx.BlockTime().Format("2006-01-02 15:04:05"),
	}, nil
}

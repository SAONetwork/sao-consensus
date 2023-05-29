package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LoanPool(c context.Context, req *types.QueryGetLoanPoolRequest) (*types.QueryGetLoanPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetLoanPool(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLoanPoolResponse{LoanPool: val}, nil
}

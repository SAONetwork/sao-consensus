package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidateDid(goCtx context.Context, req *types.QueryValidateDidRequest) (*types.QueryValidateDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	did := req.Did
	if err := k.ValidDid(ctx, did); err != nil {
		return nil, err
	}

	return &types.QueryValidateDidResponse{}, nil
}

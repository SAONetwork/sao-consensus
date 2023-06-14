package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllFaults(goCtx context.Context, req *types.QueryAllFaultsRequest) (*types.QueryAllFaultsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	faultIds := make([]string, 0)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		faultIds = append(faultIds, string(value))

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFaultsResponse{FaultIds: faultIds, Pagination: pageRes}, nil
}

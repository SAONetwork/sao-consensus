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

	faults := make([]*types.Fault, 0)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
	nodeStore := prefix.NewStore(store, types.KeyPrefix(types.NodeKeyPrefix))

	pageRes, err := query.Paginate(nodeStore, req.Pagination, func(_ []byte, value []byte) error {
		var fault types.Fault
		if err := k.cdc.Unmarshal(value, &fault); err != nil {
			return err
		}

		faults = append(faults, &fault)

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFaultsResponse{Faults: faults, Pagination: pageRes}, nil
}

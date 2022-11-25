package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/market/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WorkerAll(c context.Context, req *types.QueryAllWorkerRequest) (*types.QueryAllWorkerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var workers []types.Worker
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	workerStore := prefix.NewStore(store, types.KeyPrefix(types.WorkerKeyPrefix))

	pageRes, err := query.Paginate(workerStore, req.Pagination, func(key []byte, value []byte) error {
		var worker types.Worker
		if err := k.cdc.Unmarshal(value, &worker); err != nil {
			return err
		}

		workers = append(workers, worker)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWorkerResponse{Worker: workers, Pagination: pageRes}, nil
}

func (k Keeper) Worker(c context.Context, req *types.QueryGetWorkerRequest) (*types.QueryGetWorkerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWorker(
		ctx,
		req.Workername,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWorkerResponse{Worker: val}, nil
}

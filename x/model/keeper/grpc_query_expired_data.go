package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ExpiredDataAll(c context.Context, req *types.QueryAllExpiredDataRequest) (*types.QueryAllExpiredDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var expiredDatas []types.ExpiredData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	expiredDataStore := prefix.NewStore(store, types.KeyPrefix(types.ExpiredDataKeyPrefix))

	pageRes, err := query.Paginate(expiredDataStore, req.Pagination, func(key []byte, value []byte) error {
		var expiredData types.ExpiredData
		if err := k.cdc.Unmarshal(value, &expiredData); err != nil {
			return err
		}

		expiredDatas = append(expiredDatas, expiredData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExpiredDataResponse{ExpiredData: expiredDatas, Pagination: pageRes}, nil
}

func (k Keeper) ExpiredData(c context.Context, req *types.QueryGetExpiredDataRequest) (*types.QueryGetExpiredDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExpiredData(
		ctx,
		req.Height,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExpiredDataResponse{ExpiredData: val}, nil
}

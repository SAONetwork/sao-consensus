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

func (k Keeper) ExpiredOrderAll(c context.Context, req *types.QueryAllExpiredOrderRequest) (*types.QueryAllExpiredOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var expiredOrders []types.ExpiredOrder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	expiredOrderStore := prefix.NewStore(store, types.KeyPrefix(types.ExpiredOrderKeyPrefix))

	pageRes, err := query.Paginate(expiredOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var expiredOrder types.ExpiredOrder
		if err := k.cdc.Unmarshal(value, &expiredOrder); err != nil {
			return err
		}

		expiredOrders = append(expiredOrders, expiredOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExpiredOrderResponse{ExpiredOrder: expiredOrders, Pagination: pageRes}, nil
}

func (k Keeper) ExpiredOrder(c context.Context, req *types.QueryGetExpiredOrderRequest) (*types.QueryGetExpiredOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExpiredOrder(
		ctx,
		req.Height,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExpiredOrderResponse{ExpiredOrder: val}, nil
}

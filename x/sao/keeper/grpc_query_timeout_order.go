package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TimeoutOrderAll(c context.Context, req *types.QueryAllTimeoutOrderRequest) (*types.QueryAllTimeoutOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timeoutOrders []types.TimeoutOrder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	timeoutOrderStore := prefix.NewStore(store, types.KeyPrefix(types.TimeoutOrderKeyPrefix))

	pageRes, err := query.Paginate(timeoutOrderStore, req.Pagination, func(key []byte, value []byte) error {
		var timeoutOrder types.TimeoutOrder
		if err := k.cdc.Unmarshal(value, &timeoutOrder); err != nil {
			return err
		}

		timeoutOrders = append(timeoutOrders, timeoutOrder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimeoutOrderResponse{TimeoutOrder: timeoutOrders, Pagination: pageRes}, nil
}

func (k Keeper) TimeoutOrder(c context.Context, req *types.QueryGetTimeoutOrderRequest) (*types.QueryGetTimeoutOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTimeoutOrder(
		ctx,
		req.Height,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTimeoutOrderResponse{TimeoutOrder: val}, nil
}

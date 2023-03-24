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

func (k Keeper) OrderFinishAll(c context.Context, req *types.QueryAllOrderFinishRequest) (*types.QueryAllOrderFinishResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orderFinishs []types.OrderFinish
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	orderFinishStore := prefix.NewStore(store, types.KeyPrefix(types.OrderFinishKeyPrefix))

	pageRes, err := query.Paginate(orderFinishStore, req.Pagination, func(key []byte, value []byte) error {
		var orderFinish types.OrderFinish
		if err := k.cdc.Unmarshal(value, &orderFinish); err != nil {
			return err
		}

		orderFinishs = append(orderFinishs, orderFinish)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrderFinishResponse{OrderFinish: orderFinishs, Pagination: pageRes}, nil
}

func (k Keeper) OrderFinish(c context.Context, req *types.QueryGetOrderFinishRequest) (*types.QueryGetOrderFinishResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOrderFinish(
		ctx,
		req.Timestamp,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOrderFinishResponse{OrderFinish: val}, nil
}

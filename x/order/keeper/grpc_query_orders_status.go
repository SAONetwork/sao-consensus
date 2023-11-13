package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrdersStatus(goCtx context.Context, req *types.QueryOrdersStatusRequest) (*types.QueryOrdersStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	status := make([]int32, len(req.OrderIds))
	for i, id := range req.OrderIds {
		order, found := k.GetOrder(ctx, id)
		if !found {
			status[i] = -1
		} else {
			status[i] = order.Status
		}
	}

	return &types.QueryOrdersStatusResponse{OrderStatus: status}, nil
}

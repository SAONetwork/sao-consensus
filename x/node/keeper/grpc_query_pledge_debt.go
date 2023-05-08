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

func (k Keeper) PledgeDebtAll(c context.Context, req *types.QueryAllPledgeDebtRequest) (*types.QueryAllPledgeDebtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pledgeDebts []types.PledgeDebt
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pledgeDebtStore := prefix.NewStore(store, types.KeyPrefix(types.PledgeDebtKeyPrefix))

	pageRes, err := query.Paginate(pledgeDebtStore, req.Pagination, func(key []byte, value []byte) error {
		var pledgeDebt types.PledgeDebt
		if err := k.cdc.Unmarshal(value, &pledgeDebt); err != nil {
			return err
		}

		pledgeDebts = append(pledgeDebts, pledgeDebt)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPledgeDebtResponse{PledgeDebt: pledgeDebts, Pagination: pageRes}, nil
}

func (k Keeper) PledgeDebt(c context.Context, req *types.QueryGetPledgeDebtRequest) (*types.QueryGetPledgeDebtResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPledgeDebt(
		ctx,
		req.Sp,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPledgeDebtResponse{PledgeDebt: val}, nil
}

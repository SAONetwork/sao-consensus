package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidBalancesAll(c context.Context, req *types.QueryAllDidBalancesRequest) (*types.QueryAllDidBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var didBalancess []types.DidBalances
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didBalancesStore := prefix.NewStore(store, types.KeyPrefix(types.DidBalancesKeyPrefix))

	pageRes, err := query.Paginate(didBalancesStore, req.Pagination, func(key []byte, value []byte) error {
		var didBalances types.DidBalances
		if err := k.cdc.Unmarshal(value, &didBalances); err != nil {
			return err
		}

		didBalancess = append(didBalancess, didBalances)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidBalancesResponse{DidBalances: didBalancess, Pagination: pageRes}, nil
}

func (k Keeper) DidBalances(c context.Context, req *types.QueryGetDidBalancesRequest) (*types.QueryGetDidBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDidBalances(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDidBalancesResponse{DidBalances: val}, nil
}

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

func (k Keeper) KidAll(c context.Context, req *types.QueryAllKidRequest) (*types.QueryAllKidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kids []types.Kid
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	kidStore := prefix.NewStore(store, types.KeyPrefix(types.KidKeyPrefix))

	pageRes, err := query.Paginate(kidStore, req.Pagination, func(key []byte, value []byte) error {
		var kid types.Kid
		if err := k.cdc.Unmarshal(value, &kid); err != nil {
			return err
		}

		kids = append(kids, kid)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKidResponse{Kid: kids, Pagination: pageRes}, nil
}

func (k Keeper) Kid(c context.Context, req *types.QueryGetKidRequest) (*types.QueryGetKidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetKid(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKidResponse{Kid: val}, nil
}

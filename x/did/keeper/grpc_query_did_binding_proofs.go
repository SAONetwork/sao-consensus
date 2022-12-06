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

func (k Keeper) DidBingingProofAll(c context.Context, req *types.QueryAllDidBingingProofRequest) (*types.QueryAllDidBingingProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DidBingingProofs []types.DidBingingProof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	DidBingingProofStore := prefix.NewStore(store, types.KeyPrefix(types.DidBingingProofKeyPrefix))

	pageRes, err := query.Paginate(DidBingingProofStore, req.Pagination, func(key []byte, value []byte) error {
		var DidBingingProof types.DidBingingProof
		if err := k.cdc.Unmarshal(value, &DidBingingProof); err != nil {
			return err
		}

		DidBingingProofs = append(DidBingingProofs, DidBingingProof)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidBingingProofResponse{DidBingingProof: DidBingingProofs, Pagination: pageRes}, nil
}

func (k Keeper) DidBingingProof(c context.Context, req *types.QueryGetDidBingingProofRequest) (*types.QueryGetDidBingingProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDidBingingProof(
		ctx,
		req.AccountId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDidBingingProofResponse{DidBingingProof: val}, nil
}

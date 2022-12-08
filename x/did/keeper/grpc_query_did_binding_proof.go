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

func (k Keeper) DidBindingProofAll(c context.Context, req *types.QueryAllDidBindingProofRequest) (*types.QueryAllDidBindingProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DidBindingProofs []types.DidBindingProof
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	DidBindingProofStore := prefix.NewStore(store, types.KeyPrefix(types.DidBindingProofKeyPrefix))

	pageRes, err := query.Paginate(DidBindingProofStore, req.Pagination, func(key []byte, value []byte) error {
		var DidBindingProof types.DidBindingProof
		if err := k.cdc.Unmarshal(value, &DidBindingProof); err != nil {
			return err
		}

		DidBindingProofs = append(DidBindingProofs, DidBindingProof)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidBindingProofResponse{DidBindingProof: DidBindingProofs, Pagination: pageRes}, nil
}

func (k Keeper) DidBindingProof(c context.Context, req *types.QueryGetDidBindingProofRequest) (*types.QueryGetDidBindingProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDidBindingProof(
		ctx,
		req.AccountId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDidBindingProofResponse{DidBindingProof: val}, nil
}

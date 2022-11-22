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

func (k Keeper) DidBindingProofsAll(c context.Context, req *types.QueryAllDidBindingProofsRequest) (*types.QueryAllDidBindingProofsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var didBindingProofss []types.DidBindingProofs
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	didBindingProofsStore := prefix.NewStore(store, types.KeyPrefix(types.DidBindingProofsKeyPrefix))

	pageRes, err := query.Paginate(didBindingProofsStore, req.Pagination, func(key []byte, value []byte) error {
		var didBindingProofs types.DidBindingProofs
		if err := k.cdc.Unmarshal(value, &didBindingProofs); err != nil {
			return err
		}

		didBindingProofss = append(didBindingProofss, didBindingProofs)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDidBindingProofsResponse{DidBindingProofs: didBindingProofss, Pagination: pageRes}, nil
}

func (k Keeper) DidBindingProofs(c context.Context, req *types.QueryGetDidBindingProofsRequest) (*types.QueryGetDidBindingProofsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDidBindingProofs(
		ctx,
		req.AccountId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDidBindingProofsResponse{DidBindingProofs: val}, nil
}

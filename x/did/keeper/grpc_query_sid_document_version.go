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

func (k Keeper) SidDocumentVersionAll(c context.Context, req *types.QueryAllSidDocumentVersionRequest) (*types.QueryAllSidDocumentVersionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sidDocumentVersions []types.SidDocumentVersion
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sidDocumentVersionStore := prefix.NewStore(store, types.KeyPrefix(types.SidDocumentVersionKeyPrefix))

	pageRes, err := query.Paginate(sidDocumentVersionStore, req.Pagination, func(key []byte, value []byte) error {
		var sidDocumentVersion types.SidDocumentVersion
		if err := k.cdc.Unmarshal(value, &sidDocumentVersion); err != nil {
			return err
		}

		sidDocumentVersions = append(sidDocumentVersions, sidDocumentVersion)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSidDocumentVersionResponse{SidDocumentVersion: sidDocumentVersions, Pagination: pageRes}, nil
}

func (k Keeper) SidDocumentVersion(c context.Context, req *types.QueryGetSidDocumentVersionRequest) (*types.QueryGetSidDocumentVersionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSidDocumentVersion(
		ctx,
		req.DocId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSidDocumentVersionResponse{SidDocumentVersion: val}, nil
}

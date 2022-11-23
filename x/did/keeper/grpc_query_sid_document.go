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

func (k Keeper) SidDocumentAll(c context.Context, req *types.QueryAllSidDocumentRequest) (*types.QueryAllSidDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sidDocuments []types.SidDocument
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sidDocumentStore := prefix.NewStore(store, types.KeyPrefix(types.SidDocumentKeyPrefix))

	pageRes, err := query.Paginate(sidDocumentStore, req.Pagination, func(key []byte, value []byte) error {
		var sidDocument types.SidDocument
		if err := k.cdc.Unmarshal(value, &sidDocument); err != nil {
			return err
		}

		sidDocuments = append(sidDocuments, sidDocument)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSidDocumentResponse{SidDocument: sidDocuments, Pagination: pageRes}, nil
}

func (k Keeper) SidDocument(c context.Context, req *types.QueryGetSidDocumentRequest) (*types.QueryGetSidDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSidDocument(
		ctx,
		req.VersionId,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSidDocumentResponse{SidDocument: val}, nil
}

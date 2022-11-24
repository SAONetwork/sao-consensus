package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CleanupSidDocuments(goCtx context.Context, msg *types.MsgCleanupSidDocuments) (*types.MsgCleanupSidDocumentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	rootDocId := msg.RootDocId
	docVersions, found := k.GetSidDocumentVersion(ctx, rootDocId)
	if !found {
		return nil, types.ErrVersionsNotFound
	}
	for _, version := range docVersions.VersionList {
		k.RemoveSidDocument(ctx, version)
	}

	k.RemoveSidDocumentVersion(ctx, rootDocId)

	return &types.MsgCleanupSidDocumentsResponse{}, nil
}

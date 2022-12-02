package keeper

import (
	"context"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO Remove
func (k msgServer) ResetStore(goCtx context.Context, msg *types.MsgResetStore) (*types.MsgResetStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sdvlist := k.GetAllSidDocumentVersion(ctx)
	for _, item := range sdvlist {
		k.RemoveSidDocumentVersion(ctx, item.DocId)
	}

	sdlist := k.GetAllSidDocument(ctx)
	for _, item := range sdlist {
		k.RemoveSidDocument(ctx, item.VersionId)
	}

	aalist := k.GetAllAccountAuth(ctx)
	for _, item := range aalist {
		k.RemoveAccountAuth(ctx, item.AccountDid)
	}

	allist := k.GetAllAccountList(ctx)
	for _, item := range allist {
		k.RemoveAccountList(ctx, item.Did)
	}

	dbplist := k.GetAllDidBindingProofs(ctx)
	for _, item := range dbplist {
		k.RemoveDidBindingProofs(ctx, item.AccountId)
	}

	pslist := k.GetAllPastSeeds(ctx)
	for _, item := range pslist {
		k.RemovePastSeeds(ctx, item.Did)
	}

	palist := k.GetAllPaymentAddress(ctx)
	for _, item := range palist {
		k.RemovePaymentAddress(ctx, item.Did)
	}

	return &types.MsgResetStoreResponse{}, nil
}

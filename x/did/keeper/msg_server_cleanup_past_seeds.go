package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CleanupPastSeeds(goCtx context.Context, msg *types.MsgCleanupPastSeeds) (*types.MsgCleanupPastSeedsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	did := msg.Did

	_, found := k.GetPastSeeds(ctx, did)
	if !found {
		return nil, types.ErrSeedsNotFound
	}

	k.RemovePastSeeds(ctx, did)

	return &types.MsgCleanupPastSeedsResponse{}, nil
}

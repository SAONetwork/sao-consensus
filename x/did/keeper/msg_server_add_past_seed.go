package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddPastSeed(goCtx context.Context, msg *types.MsgAddPastSeed) (*types.MsgAddPastSeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	did := msg.Did
	seed := msg.PastSeed
	pastSeeds, found := k.GetPastSeeds(ctx, did)
	if found {
		for _, ps := range pastSeeds.Seeds {
			if ps == seed {
				return nil, types.ErrSeedExists
			}
		}
		pastSeeds.Seeds = append(pastSeeds.Seeds, seed)
	} else {
		pastSeeds = types.PastSeeds{Did: did, Seeds: []string{seed}}
	}

	k.SetPastSeeds(ctx, pastSeeds)

	return &types.MsgAddPastSeedResponse{}, nil
}

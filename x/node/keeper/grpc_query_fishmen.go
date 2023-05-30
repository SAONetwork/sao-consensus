package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Fishmen(goCtx context.Context, req *types.QueryFishmenRequest) (*types.QueryFishmenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FishmenKeyPrefix))
	currentFishmenBytes := store.Get([]byte("Fishmen"))

	return &types.QueryFishmenResponse{
		Fishmen: string(currentFishmenBytes),
	}, nil
}

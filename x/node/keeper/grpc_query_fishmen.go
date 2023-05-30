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

	fishmenParamStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FishmenKeyPrefix))
	currentFishmenParamBytes := fishmenParamStore.Get([]byte("FishmenParam"))
	var fishmenParam types.FishmenParam
	err := fishmenParam.Unmarshal(currentFishmenParamBytes)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "invalid fishmen params")
	}

	return &types.QueryFishmenResponse{
		Fishmen: fishmenParam.Fishmen,
	}, nil
}

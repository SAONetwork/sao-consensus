package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllAccountAuths(goCtx context.Context, req *types.QueryGetAllAccountAuthsRequest) (*types.QueryGetAllAccountAuthsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	did := req.Did
	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		return nil, status.Error(codes.NotFound, "account list not found")
	}

	accAuths := make([]*types.AccountAuth, 0)
	for _, accDid := range accountList.AccountDids {
		accAuth, found := k.GetAccountAuth(ctx, accDid)
		if found {
			accAuths = append(accAuths, &accAuth)
		}
	}

	return &types.QueryGetAllAccountAuthsResponse{accAuths}, nil
}

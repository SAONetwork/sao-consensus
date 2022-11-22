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

func (k Keeper) AccountAuthAll(c context.Context, req *types.QueryAllAccountAuthRequest) (*types.QueryAllAccountAuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountAuths []types.AccountAuth
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountAuthStore := prefix.NewStore(store, types.KeyPrefix(types.AccountAuthKeyPrefix))

	pageRes, err := query.Paginate(accountAuthStore, req.Pagination, func(key []byte, value []byte) error {
		var accountAuth types.AccountAuth
		if err := k.cdc.Unmarshal(value, &accountAuth); err != nil {
			return err
		}

		accountAuths = append(accountAuths, accountAuth)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountAuthResponse{AccountAuth: accountAuths, Pagination: pageRes}, nil
}

func (k Keeper) AccountAuth(c context.Context, req *types.QueryGetAccountAuthRequest) (*types.QueryGetAccountAuthResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAccountAuth(
		ctx,
		req.AccountDid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAccountAuthResponse{AccountAuth: val}, nil
}

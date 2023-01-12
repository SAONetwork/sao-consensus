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

func (k Keeper) AccountIdAll(c context.Context, req *types.QueryAllAccountIdRequest) (*types.QueryAllAccountIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountIds []types.AccountId
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountIdStore := prefix.NewStore(store, types.KeyPrefix(types.AccountIdKeyPrefix))

	pageRes, err := query.Paginate(accountIdStore, req.Pagination, func(key []byte, value []byte) error {
		var accountId types.AccountId
		if err := k.cdc.Unmarshal(value, &accountId); err != nil {
			return err
		}

		accountIds = append(accountIds, accountId)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountIdResponse{AccountId: accountIds, Pagination: pageRes}, nil
}

func (k Keeper) AccountId(c context.Context, req *types.QueryGetAccountIdRequest) (*types.QueryGetAccountIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAccountId(
		ctx,
		req.AccountDid,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAccountIdResponse{AccountId: val}, nil
}

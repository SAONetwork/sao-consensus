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

func (k Keeper) AccountListAll(c context.Context, req *types.QueryAllAccountListRequest) (*types.QueryAllAccountListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountLists []types.AccountList
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	accountListStore := prefix.NewStore(store, types.KeyPrefix(types.AccountListKeyPrefix))

	pageRes, err := query.Paginate(accountListStore, req.Pagination, func(key []byte, value []byte) error {
		var accountList types.AccountList
		if err := k.cdc.Unmarshal(value, &accountList); err != nil {
			return err
		}

		accountLists = append(accountLists, accountList)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountListResponse{AccountList: accountLists, Pagination: pageRes}, nil
}

func (k Keeper) AccountList(c context.Context, req *types.QueryGetAccountListRequest) (*types.QueryGetAccountListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetAccountList(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAccountListResponse{AccountList: val}, nil
}

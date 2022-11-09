package keeper

import (
	"context"
	"math/big"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PledgeAll(c context.Context, req *types.QueryAllPledgeRequest) (*types.QueryAllPledgeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pledges []types.Pledge
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	pledgeStore := prefix.NewStore(store, types.KeyPrefix(types.PledgeKeyPrefix))

	pageRes, err := query.Paginate(pledgeStore, req.Pagination, func(key []byte, value []byte) error {
		var pledge types.Pledge
		if err := k.cdc.Unmarshal(value, &pledge); err != nil {
			return err
		}

		pledges = append(pledges, pledge)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPledgeResponse{Pledge: pledges, Pagination: pageRes}, nil
}

func (k Keeper) Pledge(c context.Context, req *types.QueryGetPledgeRequest) (*types.QueryGetPledgeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPledge(
		ctx,
		req.Creator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}
	pool, found_pool := k.GetPool(ctx)
	if !found_pool {
		return nil, sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	coinPerShare, _ := new(big.Int).SetString(pool.CoinPerShare, 10)

	pending := new(big.Int).Sub(new(big.Int).Div(new(big.Int).Mul(val.Pledged.Amount.BigInt(), coinPerShare), big.NewInt(1e12)), val.RewardDebt.Amount.BigInt())

	val.Reward = val.Reward.AddAmount(sdk.NewInt(pending.Int64()))

	return &types.QueryGetPledgeResponse{Pledge: val}, nil
}

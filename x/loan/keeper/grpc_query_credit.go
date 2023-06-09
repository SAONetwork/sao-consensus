package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/loan/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CreditAll(c context.Context, req *types.QueryAllCreditRequest) (*types.QueryAllCreditResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var credits []types.Credit
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	creditStore := prefix.NewStore(store, types.KeyPrefix(types.CreditKeyPrefix))

	pageRes, err := query.Paginate(creditStore, req.Pagination, func(key []byte, value []byte) error {
		var credit types.Credit
		if err := k.cdc.Unmarshal(value, &credit); err != nil {
			return err
		}

		credits = append(credits, credit)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCreditResponse{Credit: credits, Pagination: pageRes}, nil
}

func (k Keeper) Credit(c context.Context, req *types.QueryGetCreditRequest) (*types.QueryGetCreditResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetCredit(
		ctx,
		req.Account,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		return nil, types.ErrLoanPoolNotFound
	}

	var total sdk.Dec
	if !loanPool.LoanedOut.IsZero() {
		interest := loanPool.AccInterestPerCoin.Amount.MulInt(loanPool.LoanedOut.Amount).Sub(loanPool.InterestDebt.Amount)
		total = loanPool.Total.Amount.Add(interest)
	} else {
		total = loanPool.Total.Amount
	}
	exchange, _ := sdk.NewDecCoinFromDec(loanPool.Total.Denom, total.Mul(val.Bonds).Quo(loanPool.TotalBonds)).TruncateDecimal()

	return &types.QueryGetCreditResponse{Credit: val, Amount: exchange}, nil
}

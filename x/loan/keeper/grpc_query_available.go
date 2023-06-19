package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Available(goCtx context.Context, req *types.QueryAvailableRequest) (*types.QueryAvailableResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenBalance := k.bank.GetBalance(ctx, sdk.MustAccAddressFromBech32(req.Account), SaoLoanTokenDenom)
	if tokenBalance.IsZero() {
		return &types.QueryAvailableResponse{
			Total:     sdk.NewCoin(SaoDenom, sdk.NewInt(0)),
			Available: sdk.NewCoin(SaoDenom, sdk.NewInt(0)),
		}, nil
	}

	loanPool, found := k.GetLoanPool(ctx)
	if !found {
		return nil, types.ErrLoanPoolNotFound
	}

	totalSlt := k.bank.GetSupply(ctx, SaoLoanTokenDenom)

	convertableSao := sdk.NewCoin(loanPool.Total.Denom, loanPool.Total.Amount.MulInt(tokenBalance.Amount).QuoInt(totalSlt.Amount).TruncateInt())

	if convertableSao.IsZero() {
		return &types.QueryAvailableResponse{
			Total:     sdk.NewCoin(SaoDenom, sdk.NewInt(0)),
			Available: sdk.NewCoin(SaoDenom, sdk.NewInt(0)),
		}, nil
	}

	totalSao, _ := loanPool.Total.TruncateDecimal()
	availableSao := totalSao.Sub(loanPool.LoanedOut)

	if availableSao.IsLT(convertableSao) {
		return &types.QueryAvailableResponse{
			Total:     convertableSao,
			Available: availableSao,
		}, nil
	} else {
		return &types.QueryAvailableResponse{
			Total:     convertableSao,
			Available: convertableSao,
		}, nil
	}
}

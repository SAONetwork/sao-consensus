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

func (k Keeper) PaymentAddressAll(c context.Context, req *types.QueryAllPaymentAddressRequest) (*types.QueryAllPaymentAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var paymentAddresss []types.PaymentAddress
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	paymentAddressStore := prefix.NewStore(store, types.KeyPrefix(types.PaymentAddressKeyPrefix))

	pageRes, err := query.Paginate(paymentAddressStore, req.Pagination, func(key []byte, value []byte) error {
		var paymentAddress types.PaymentAddress
		if err := k.cdc.Unmarshal(value, &paymentAddress); err != nil {
			return err
		}

		paymentAddresss = append(paymentAddresss, paymentAddress)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPaymentAddressResponse{PaymentAddress: paymentAddresss, Pagination: pageRes}, nil
}

func (k Keeper) PaymentAddress(c context.Context, req *types.QueryGetPaymentAddressRequest) (*types.QueryGetPaymentAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPaymentAddress(
		ctx,
		req.Did,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPaymentAddressResponse{PaymentAddress: val}, nil
}

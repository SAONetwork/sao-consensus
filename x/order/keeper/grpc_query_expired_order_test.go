package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/order/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestExpiredOrderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNExpiredOrder(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetExpiredOrderRequest
		response *types.QueryGetExpiredOrderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetExpiredOrderRequest{
				Height: msgs[0].Height,
			},
			response: &types.QueryGetExpiredOrderResponse{ExpiredOrder: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetExpiredOrderRequest{
				Height: msgs[1].Height,
			},
			response: &types.QueryGetExpiredOrderResponse{ExpiredOrder: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetExpiredOrderRequest{
				Height: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ExpiredOrder(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestExpiredOrderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.OrderKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNExpiredOrder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllExpiredOrderRequest {
		return &types.QueryAllExpiredOrderRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ExpiredOrderAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ExpiredOrder), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ExpiredOrder),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ExpiredOrderAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ExpiredOrder), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ExpiredOrder),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ExpiredOrderAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ExpiredOrder),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ExpiredOrderAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

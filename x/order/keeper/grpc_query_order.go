package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	db "github.com/tendermint/tm-db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrderAll(c context.Context, req *types.QueryAllOrderRequest) (*types.QueryAllOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orders []types.Order
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))

	var err error
	var pageRes *query.PageResponse
	if req.Did != "" {
		pageRequest := req.Pagination
		// if the PageRequest is nil, use default PageRequest
		if pageRequest == nil {
			pageRequest = &query.PageRequest{}
		}

		offset := pageRequest.Offset
		key := pageRequest.Key
		limit := pageRequest.Limit
		countTotal := pageRequest.CountTotal
		reverse := pageRequest.Reverse

		if offset > 0 && key != nil {
			return nil, fmt.Errorf("invalid request, either offset or key is expected, got both")
		}

		if limit == 0 {
			limit = query.DefaultLimit

			// count total results when the limit is zero/not supplied
			countTotal = true
		}
		iterator := getIterator(orderStore, nil, reverse)
		defer iterator.Close()

		end := offset + limit

		var count uint64
		var nextKey []byte

		for ; iterator.Valid(); iterator.Next() {
			var order types.Order
			if err := k.cdc.Unmarshal(iterator.Value(), &order); err != nil {
				return nil, err
			}
			if order.Owner == req.Did {
				contains := false
				for _, i := range req.States {
					if i == order.Status {
						count++
						contains = true
					}
				}
				if !contains {
					continue
				}
			}

			if count <= offset {
				continue
			}
			if count <= end {
				orders = append(orders, order)
			} else if count == end+1 {
				nextKey = iterator.Key()

				if !countTotal {
					break
				}
			}
			if iterator.Error() != nil {
				return nil, iterator.Error()
			}
		}

		res := &query.PageResponse{NextKey: nextKey}
		if countTotal {
			res.Total = count
		}
	} else {
		pageRes, err = query.Paginate(orderStore, req.Pagination, func(key []byte, value []byte) error {
			var order types.Order
			if err := k.cdc.Unmarshal(value, &order); err != nil {
				return err
			}

			orders = append(orders, order)
			return nil
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &types.QueryAllOrderResponse{Order: orders, Pagination: pageRes}, nil
}

func getIterator(prefixStore storetypes.KVStore, start []byte, reverse bool) db.Iterator {
	if reverse {
		var end []byte
		if start != nil {
			itr := prefixStore.Iterator(start, nil)
			defer itr.Close()
			if itr.Valid() {
				itr.Next()
				end = itr.Key()
			}
		}
		return prefixStore.ReverseIterator(nil, end)
	}
	return prefixStore.Iterator(start, nil)
}

func (k Keeper) Order(c context.Context, req *types.QueryGetOrderRequest) (*types.QueryGetOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	order, found := k.GetOrder(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	shards := make(map[string]*types.Shard, 0)
	for _, id := range order.Shards {
		shard, found := k.GetShard(ctx, id)
		if found {
			shards[shard.Sp] = &shard
		}

	}

	fullOrder := types.FullOrder{
		Creator:       order.Creator,
		Owner:         order.Owner,
		Id:            order.Id,
		Provider:      order.Provider,
		Cid:           order.Cid,
		Duration:      order.Duration,
		Status:        order.Status,
		Replica:       order.Replica,
		ShardIds:      order.Shards,
		Shards:        shards,
		Amount:        order.Amount,
		Size_:         order.Size_,
		Operation:     order.Operation,
		CreatedAt:     order.CreatedAt,
		Timeout:       order.Timeout,
		DataId:        order.DataId,
		Commit:        order.Commit,
		RewardPerByte: order.RewardPerByte,
	}

	return &types.QueryGetOrderResponse{Order: fullOrder}, nil
}

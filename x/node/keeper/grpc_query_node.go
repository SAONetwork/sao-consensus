package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NodeAll(c context.Context, req *types.QueryAllNodeRequest) (*types.QueryAllNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nodes []types.Node
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nodeStore := prefix.NewStore(store, types.KeyPrefix(types.NodeKeyPrefix))

	pageRes, err := query.Paginate(nodeStore, req.Pagination, func(_ []byte, value []byte) error {
		var node types.Node
		if err := k.cdc.Unmarshal(value, &node); err != nil {
			return err
		}

		if req.Status == types.NODE_STATUS_NA || req.Status&node.Status > 0 {
			nodes = append(nodes, node)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNodeResponse{Node: nodes, Pagination: pageRes}, nil
}

func (k Keeper) Node(c context.Context, req *types.QueryGetNodeRequest) (*types.QueryGetNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNode(
		ctx,
		req.Creator,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNodeResponse{Node: val}, nil
}

package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Migrate(goCtx context.Context, msg *types.MsgMigrate) (*types.MsgMigrateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	fmt.Println(ctx)

	resp := types.MsgMigrateResponse{
		Result: make(map[string]string, 0),
	}

	for _, dataId := range msg.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			resp.Result[dataId] = status.Errorf(codes.NotFound, "FAILED: dataId %s not found", dataId).Error()
			continue
		}

		shard := k.node.GetMetadataShardByNode(ctx, dataId, msg.Creator, int(metadata.Replica))
		if shard == nil {
			resp.Result[dataId] = status.Errorf(codes.NotFound, "FAILED: %s shard not found", msg.Creator).Error()
			continue
		}

		sps := k.node.RandomSP(ctx, 1)

		newShard := k.node.MigrateShard(ctx, dataId, shard.Index, sps[0].Creator)

		if newShard != nil {
			resp.Result[dataId] = fmt.Sprintf("SUCCESS: new storage provider %s", sps[0].Creator)
		} else {
			resp.Result[dataId] = fmt.Sprintf("FAILED: %s is not valid data provider of %s", msg.Creator, dataId)
		}
	}

	return &resp, nil
}

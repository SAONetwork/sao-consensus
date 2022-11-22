package keeper

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Unbinding(goCtx context.Context, msg *types.MsgUnbinding) (*types.MsgUnbindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	accountId := msg.GetAccountId()
	_, found := k.GetDidBindingProofs(ctx, accountId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	k.RemoveDidBindingProofs(ctx, accountId)

	return &types.MsgUnbindingResponse{}, nil
}

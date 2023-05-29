package keeper

import (
	"context"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReportFailures(goCtx context.Context, msg *types.MsgReportFailures) (*types.MsgReportFailuresResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.node.GetNode(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s", msg.Creator)
	}

	if node.Status != nodetypes.NODE_STATUS_SERVE_FISHING {
		return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidStatus, "%s", msg.Creator)
	}

	provider, found := k.node.GetNode(ctx, msg.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s", msg.Provider)
	}

	if provider.Status != nodetypes.NODE_STATUS_SERVE_STORAGE {
		return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidStatus, "%s", msg.Provider)
	}

	for _, failure := range msg.Failures {
		if msg.ReportType == nodetypes.ReportTypeRaiseFailure {
			k.node.SetPledge()
		}
	}

	return &types.MsgReportFailuresResponse{}, nil
}

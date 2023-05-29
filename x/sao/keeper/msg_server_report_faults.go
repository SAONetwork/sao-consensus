package keeper

import (
	"context"
	"strings"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ReportFaults(goCtx context.Context, msg *types.MsgReportFaults) (*types.MsgReportFaultsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.node.GetNode(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s", msg.Creator)
	}

	if node.Status != nodetypes.NODE_STATUS_SERVE_FISHING {
		return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidStatus, "%s", msg.Creator)
	}

	reportedFaults := make([]*nodetypes.Fault, 0)
	confirmedFaults := make([]*nodetypes.Fault, 0)
	for _, fault := range msg.Faults {
		if msg.Provider != fault.Provider {
			continue
		}

		_, found := k.model.GetMetadata(ctx, fault.DataId)
		if !found {
			continue
		}

		orderMeta, found := k.order.GetOrder(ctx, fault.OrderId)
		if !found {
			continue
		}

		if orderMeta.DataId != fault.DataId || strings.Contains(orderMeta.Commit, fault.CommitId) {
			continue
		}

		isValidInfo := false
		for _, shardId := range orderMeta.Shards {
			shard, found := k.order.GetShard(ctx, shardId)
			if found && shard.Sp == fault.Provider {
				if shard.CreatedAt+shard.Duration > uint64(ctx.BlockHeight()) {
					isValidInfo = true
				}
				break
			}
		}
		if !isValidInfo {
			continue
		}

		faultOrg, found := k.node.GetFault(ctx, fault.Provider, fault.ShardId)
		faultMeta := &nodetypes.Fault{
			DataId:   fault.DataId,
			ShardId:  fault.ShardId,
			CommitId: fault.CommitId,
			Provider: fault.Provider,
			Reporter: fault.Reporter,
		}
		if found {
			if faultOrg.Reporter != fault.Reporter {
				faultMeta.Confirms = fault.Confirms + 1
			} else {
				faultMeta.Confirms = fault.Confirms
			}

			if faultMeta.Confirms > 4 {
				faultMeta.Status = nodetypes.FaultStatusConfirmed
				confirmedFaults = append(confirmedFaults, faultMeta)
			}
		} else {
			faultMeta.FaultId = ""
			faultMeta.Status = nodetypes.FaultStatusConfirming
		}

		reportedFaults = append(reportedFaults, faultMeta)
		k.node.SetFault(ctx, faultMeta)
	}

	if len(reportedFaults) > 0 {
		faultIds := ""
		for index, fault := range reportedFaults {
			if index > 0 {
				faultIds = faultIds + "," + fault.FaultId
			} else {
				faultIds = fault.FaultId
			}
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.FaultsReportedEventType,
				sdk.NewAttribute("provider", msg.Provider),
				sdk.NewAttribute("faults-ids", faultIds),
			),
		)
	}

	if len(confirmedFaults) > 0 {
		faultIds := ""
		for index, fault := range confirmedFaults {
			if index > 0 {
				faultIds = faultIds + "," + fault.FaultId
			} else {
				faultIds = fault.FaultId
			}
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.FaultsConfirmedEventType,
				sdk.NewAttribute("provider", msg.Provider),
				sdk.NewAttribute("faults-ids", faultIds),
			),
		)
	}

	return &types.MsgReportFaultsResponse{}, nil
}

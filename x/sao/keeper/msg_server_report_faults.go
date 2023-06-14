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

	fishmenInfo := k.node.FishmenInfo(ctx)
	if !strings.Contains(fishmenInfo, node.Creator) {
		return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidFinshmen, "%s is not a fishmen", msg.Creator)
	}

	// if node.Status != nodetypes.NODE_STATUS_SERVE_FISHING {
	// 	return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidStatus, "fishing is not enabled on %s", msg.Creator)
	// }

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
			if shardId != fault.ShardId {
				continue
			}

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

		faultOrg, found := k.node.GetFaultBySpAndShardId(ctx, fault.Provider, fault.ShardId)
		faultMeta := &nodetypes.Fault{
			DataId:   fault.DataId,
			OrderId:  fault.OrderId,
			ShardId:  fault.ShardId,
			CommitId: fault.CommitId,
			Provider: fault.Provider,
		}
		if found {
			faultMeta.Reporter = faultOrg.Reporter
			faultMeta.FaultId = faultOrg.FaultId
			faultMeta.Penalty = faultOrg.Penalty
			faultMeta.Confirms = faultOrg.Confirms
			if faultOrg.DataId != faultMeta.DataId || faultOrg.OrderId != faultMeta.OrderId || faultOrg.ShardId != faultMeta.ShardId {
				continue
			}

			if faultOrg.Reporter != faultMeta.Reporter {
				if strings.Contains(faultOrg.Confirms, "+"+faultMeta.Reporter) {
					continue
				} else {
					faultMeta.Confirms = faultMeta.Confirms + "|+" + faultMeta.Reporter
				}
			} else {
				continue
			}

			if strings.Count(faultMeta.Confirms, "+") > 2 {
				faultMeta.Status = nodetypes.FaultStatusConfirmed
				confirmedFaults = append(confirmedFaults, faultMeta)
			}
		} else {
			faultMeta.Confirms = "+" + faultMeta.Reporter
			faultMeta.Penalty = 0
			faultMeta.FaultId = ""
			faultMeta.Status = nodetypes.FaultStatusConfirming
			faultMeta.Reporter = msg.Creator
		}

		k.Logger(ctx).Error("faultMeta:", "faultMeta", faultMeta)
		reportedFaults = append(reportedFaults, faultMeta)
		k.node.SetFault(ctx, faultMeta)
	}

	faultIds := make([]string, 0)
	if len(reportedFaults) > 0 {
		for _, fault := range reportedFaults {
			faultIds = append(faultIds, fault.FaultId)
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.FaultsReportedEventType,
				sdk.NewAttribute("provider", msg.Provider),
				sdk.NewAttribute("faults-ids", strings.Join(faultIds, ",")),
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

	return &types.MsgReportFaultsResponse{
		FaultIds: faultIds,
	}, nil
}

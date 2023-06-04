package keeper

import (
	"context"
	"strings"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RecoverFaults(goCtx context.Context, msg *types.MsgRecoverFaults) (*types.MsgRecoverFaultsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.node.GetNode(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s", msg.Creator)
	}

	if msg.Creator == msg.Provider {
		if node.Status != nodetypes.NODE_STATUS_SERVE_STORAGE {
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidStatus, "%s", msg.Creator)
		}
	} else {
		fishmenInfo := k.node.FishmenInfo(ctx)
		if !strings.Contains(fishmenInfo, node.Creator) {
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidFinshmen, "%s is not a fishmen", msg.Creator)
		}

		// if node.Status != nodetypes.NODE_STATUS_SERVE_FISHING {
		// 	return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidStatus, "fishing is not enabled on %s", msg.Creator)
		// }
	}

	pool, foundPool := k.node.GetPool(ctx)
	if !foundPool {
		return nil, sdkerrors.Wrap(types.ErrorGetPoolInfoFailed, "")
	}

	declaredFaults := make([]*nodetypes.Fault, 0)
	recoveredFaults := make([]*nodetypes.Fault, 0)
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

		faultOrg, found := k.node.GetFaultBySpAndShardId(ctx, fault.Provider, fault.ShardId)
		faultMeta := &nodetypes.Fault{
			DataId:   fault.DataId,
			ShardId:  fault.ShardId,
			CommitId: fault.CommitId,
			Provider: fault.Provider,
		}
		if found {
			faultMeta.Reporter = faultOrg.Reporter
			faultMeta.Penalty = faultOrg.Penalty
			if msg.Provider == msg.Creator && faultOrg.Provider == msg.Creator {
				faultMeta.Status = nodetypes.FaultStatusRecovering
			} else {
				if strings.Contains(faultOrg.Confirms, "-"+msg.Creator) {
					continue
				} else {
					faultMeta.Confirms = "|-" + msg.Creator
				}
			}

			if strings.Count(faultMeta.Confirms, "+") == strings.Count(faultMeta.Confirms, "-") {
				recoveredFaults = append(recoveredFaults, faultMeta)

				pledge, found := k.node.GetPledge(ctx, msg.Creator)
				if found {
					penalty := pool.AccRewardPerByte.Amount.MulInt64(int64(orderMeta.Size_) * int64(faultMeta.Penalty))
					if pledge.RewardDebt.Amount.LT(penalty) {
						pledge.RewardDebt.Amount.SetInt64(0)
					} else {
						pledge.RewardDebt.Amount.Sub(penalty)
					}

					k.node.RemoveFault(ctx, faultMeta.FaultId)
					continue
				}
			}
		} else {
			continue
		}

		declaredFaults = append(declaredFaults, faultMeta)
		k.node.SetFault(ctx, faultMeta)
	}

	if len(declaredFaults) > 0 {
		faultIds := ""
		for index, fault := range declaredFaults {
			if index > 0 {
				faultIds = faultIds + "," + fault.FaultId
			} else {
				faultIds = fault.FaultId
			}
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.FaultsRecoverDeclaredEventType,
				sdk.NewAttribute("provider", msg.Provider),
				sdk.NewAttribute("faults-ids", faultIds),
			),
		)
	}

	if len(recoveredFaults) > 0 {
		faultIds := ""
		for index, fault := range recoveredFaults {
			if index > 0 {
				faultIds = faultIds + "," + fault.FaultId
			} else {
				faultIds = fault.FaultId
			}
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.FaultsRecoveredEventType,
				sdk.NewAttribute("provider", msg.Provider),
				sdk.NewAttribute("faults-ids", faultIds),
			),
		)
	}

	return &types.MsgRecoverFaultsResponse{}, nil
}

package keeper

import (
	"context"
	"math/big"
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
		if node.Status&nodetypes.NODE_STATUS_SERVE_STORAGE == 0 {
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

		if orderMeta.DataId != fault.DataId || !strings.Contains(orderMeta.Commit, fault.CommitId) {
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
			OrderId:  fault.OrderId,
			ShardId:  fault.ShardId,
			CommitId: fault.CommitId,
			Provider: fault.Provider,
		}
		if found {
			if faultOrg.DataId != faultMeta.DataId || faultOrg.OrderId != faultMeta.OrderId || faultOrg.ShardId != faultMeta.ShardId {
				continue
			}

			faultMeta.FaultId = faultOrg.FaultId
			faultMeta.Status = faultOrg.Status
			faultMeta.Reporter = faultOrg.Reporter
			faultMeta.Penalty = faultOrg.Penalty
			faultMeta.Confirms = faultOrg.Confirms
			if msg.Provider == msg.Creator && faultOrg.Provider == msg.Creator {
				faultMeta.Status = nodetypes.FaultStatusRecovering
			} else {
				if strings.Contains(faultOrg.Confirms, "-"+msg.Creator) {
					// continue
					faultMeta.Confirms = faultMeta.Confirms + "|-" + msg.Creator
				} else {
					if faultMeta.Status == nodetypes.FaultStatusRecovering {
						faultMeta.Confirms = faultMeta.Confirms + "|-" + msg.Creator
					} else {
						continue
					}
				}
			}

			if strings.Count(faultMeta.Confirms, "+") == strings.Count(faultMeta.Confirms, "-") {
				recoveredFaults = append(recoveredFaults, faultMeta)
				pledge, found := k.node.GetPledge(ctx, faultMeta.Provider)
				if found {
					// pay penalty from reward > rewardDebt > pledge
					penalty := pool.AccRewardPerByte.Amount.MulInt64(int64(orderMeta.Size_) * int64(faultMeta.Penalty))
					reward := pool.AccRewardPerByte.Amount.MulInt64(int64(orderMeta.Size_) * int64(faultMeta.Penalty))
					if pledge.Reward.Amount.LT(penalty) {
						penalty = penalty.Sub(pledge.Reward.Amount)
						pledge.Reward.Amount.SetInt64(0)

						if pledge.RewardDebt.Amount.LT(penalty) {
							penalty = penalty.Sub(pledge.Reward.Amount)
							pledge.RewardDebt.Amount.SetInt64(0)

							pledged := sdk.NewDecCoinFromCoin(pledge.TotalStoragePledged)
							if pledged.Amount.LT(penalty) {
								reward = reward.Sub(penalty.Sub(pledged.Amount))
								pledge.TotalStoragePledged.Amount.BigInt().Set(big.NewInt(0))
							} else {
								pledge.TotalStoragePledged.Amount.Sub(penalty.TruncateInt())
							}
						} else {
							pledge.RewardDebt.Amount.Sub(penalty)
						}
					} else {
						pledge.Reward.Amount.Sub(penalty)
					}

					// some penalty amount goes to insurance pool for data lossing clients' claims
					// the reporter all the confirmers share the rest of the penalty as reward
					insurance := reward.MulInt64(40).QuoInt64(100)
					reporterReward := reward.MulInt64(10).QuoInt64(100)
					confirmerReward := reward.Sub(reporterReward).Sub(insurance)

					insuranceTotal, found := k.node.GetFishingReward(ctx, nodetypes.InsuranceKey)
					if found {
						total := insuranceTotal.Add(insurance)
						k.node.SetFishingReward(ctx, faultOrg.Reporter, &total)
					} else {
						k.node.SetFishingReward(ctx, faultOrg.Reporter, &insurance)
					}

					reporterRewardTotal, found := k.node.GetFishingReward(ctx, faultOrg.Reporter)
					if found {
						total := reporterRewardTotal.Add(reporterReward)
						k.node.SetFishingReward(ctx, faultOrg.Reporter, &total)
					} else {
						k.node.SetFishingReward(ctx, faultOrg.Reporter, &reporterReward)
					}

					confirmers := strings.ReplaceAll(faultOrg.Confirms, "+", "")
					confirmers = strings.ReplaceAll(confirmers, "-", "")
					confirmersCount := len(strings.Split(confirmers, "|"))
					if confirmersCount > 0 {
						confirmerRewardShare := confirmerReward.QuoInt64(int64(confirmersCount))
						for _, confirmer := range strings.Split(confirmers, "|") {
							rewardTotal, found := k.node.GetFishingReward(ctx, confirmer)
							if found {
								total := rewardTotal.Add(confirmerRewardShare)
								k.node.SetFishingReward(ctx, confirmer, &total)
							} else {
								k.node.SetFishingReward(ctx, confirmer, &confirmerRewardShare)
							}
						}
					}

					k.node.SetPledge(ctx, pledge)
					k.node.RemoveFault(ctx, faultMeta)
					continue
				}
			}
		} else {
			continue
		}

		declaredFaults = append(declaredFaults, faultMeta)
		k.node.SetFault(ctx, faultMeta)
	}

	faultIds := make([]string, 0)
	if len(declaredFaults) > 0 {
		for _, fault := range declaredFaults {
			faultIds = append(faultIds, fault.FaultId)
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.FaultsRecoverDeclaredEventType,
				sdk.NewAttribute("provider", msg.Provider),
				sdk.NewAttribute("faults-ids", strings.Join(faultIds, ",")),
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

	return &types.MsgRecoverFaultsResponse{
		FaultIds: faultIds,
	}, nil
}

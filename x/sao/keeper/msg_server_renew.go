package keeper

import (
	"context"
	"fmt"

	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MaxRenewDuration uint64 = 60 * 60 * 24 * 365 * 2

func (k msgServer) Renew(goCtx context.Context, msg *types.MsgRenew) (*types.MsgRenewResponse, error) {
	var sigDid string
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &msg.Proposal
	sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
	if err != nil {
		return nil, err
	}

	isProvider := false
	if msg.Provider == msg.Creator {
		isProvider = true
	} else {
		provider, found := k.node.GetNode(ctx, msg.Provider)
		if found {
			for _, address := range provider.TxAddresses {
				if address == msg.Creator {
					isProvider = true
				}
			}
		}
	}

	if !isProvider {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, msg.Provider: %s", msg.Creator, msg.Provider)
	}

	if proposal.Duration > MaxRenewDuration {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidDuration, "renew duration: %d, max renew duration: %d", proposal.Duration, MaxRenewDuration)
	}

	resp := types.MsgRenewResponse{
		Result: make([]*types.KV, 0),
	}

	pool, found := k.node.GetPool(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrPoolNotFound, "pool not found")
	}

	if err != nil {
		return nil, err
	}
	denom := k.staking.BondDenom(ctx)

	//blockRewardPerByte := pool.RewardPerBlock.Amount.Quo(sdk.NewDec(pool.TotalStorage))

dataLoop:
	for _, dataId := range proposal.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: status.Errorf(codes.NotFound, "FAILED: dataId %s not found", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		if metadata.Owner != sigDid {
			// only the data model owner could renew operations
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrorNoPermission, "FAILED: no permission to renew the model %s", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		if metadata.Status != modeltypes.MetaComplete {
			// metadata status should be completed,
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(modeltypes.ErrInvalidStatus, "FAILED: try to renew uncompleted model %s ", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		order, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrOrderNotFound, "FAILED: invalid order id: %d", metadata.OrderId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		var shards []ordertypes.Shard
		for _, id := range order.Shards {
			shard, found := k.order.GetShard(ctx, id)
			if !found {
				kv := &types.KV{
					K: dataId,
					V: sdkerrors.Wrapf(types.ErrorNoPermission, "FAILED: shardId %s not found", id).Error(),
				}
				resp.Result = append(resp.Result, kv)
				continue dataLoop
			}
			if shard.Status != ordertypes.ShardCompleted && shard.Status != ordertypes.ShardMigrating {
				kv := &types.KV{
					K: dataId,
					V: sdkerrors.Wrapf(types.ErrShardNotCompleted, "FAILED: invalid shard status: %d", shard.Status).Error(),
				}
				resp.Result = append(resp.Result, kv)
				continue dataLoop

			}
			shards = append(shards, shard)
		}

		if order.Status != ordertypes.OrderCompleted {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrOrderNotFound, "FAILED: invalid order status: %d", metadata.Status).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		duration := int64(order.Duration)
		currentHeight := ctx.BlockHeight()
		orderExpiredAt := int64(order.CreatedAt) + duration

		if orderExpiredAt < currentHeight {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrorOrderExpired, "FAILED: metadata should have expired: order has expired at %d", orderExpiredAt).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		// TODO: use real-time unit price instead
		orderUnitPrice := sdk.NewDecWithPrec(1, 6)

		amount, dec := sdk.NewDecCoinFromDec(
			denom,
			orderUnitPrice.
				MulInt64(int64(order.Replica)).
				MulInt64(int64(order.Size_)).
				MulInt64(int64(proposal.Duration))).
			TruncateDecimal()
		if !dec.IsZero() {
			amount = amount.AddAmount(sdk.NewInt(1))
		}
		newOrder := ordertypes.Order{
			Creator:   msg.Creator,
			Owner:     order.Owner,
			Provider:  msg.Provider,
			Cid:       order.Cid,
			Duration:  proposal.Duration,
			Status:    order.Status,
			Replica:   order.Replica,
			Shards:    order.Shards,
			Amount:    amount,
			Size_:     order.Size_,
			Operation: 3,
			CreatedAt: uint64(ctx.BlockHeight()),
			Timeout:   uint64(proposal.Timeout),
			DataId:    order.DataId,
			Commit:    order.Commit,
			UnitPrice: sdk.NewDecCoinFromDec(denom, orderUnitPrice),
		}

		_, err := k.order.RenewOrder(ctx, &newOrder)
		if err != nil {
			kv := &types.KV{
				K: dataId,
				V: err.Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		totalPledgeChange := sdk.NewInt(0)
		var newExpiredAt uint64 = 0
		for _, shard := range shards {
			if shard.Status == ordertypes.ShardMigrating {
				continue
			}
			spAcc := sdk.MustAccAddressFromBech32(shard.Sp)

			//blockRewardPledge := k.node.BlockRewardPledge(proposal.Duration, shard.Size_, sdk.NewDecCoinFromDec(denom, blockRewardPerByte))
			storeRewardPledge := k.node.StoreRewardPledge(proposal.Duration, shard.Size_, newOrder.UnitPrice)

			newPledge, dec := sdk.NewDecCoinFromDec(denom, storeRewardPledge).TruncateDecimal()
			if !dec.IsZero() {
				newPledge = newPledge.AddAmount(sdk.NewInt(1))
			}

			if newPledge.Amount.GT(shard.Pledge.Amount) {
				extraPledge := newPledge.Sub(shard.Pledge)
				spBalance := k.bank.GetBalance(ctx, spAcc, denom)
				if spBalance.IsGTE(extraPledge) {
					k.bank.SendCoinsFromAccountToModule(ctx, spAcc, nodetypes.ModuleName, sdk.Coins{extraPledge})
				} else {
					k.bank.SendCoinsFromAccountToModule(ctx, spAcc, nodetypes.ModuleName, sdk.Coins{spBalance})
					debt := extraPledge.Sub(spBalance)
					pledgeDebt, found := k.node.GetPledgeDebt(ctx, shard.Sp)
					if !found {
						pledgeDebt = nodetypes.PledgeDebt{
							Sp:   shard.Sp,
							Debt: debt,
						}
					} else {
						pledgeDebt.Debt = pledgeDebt.Debt.Add(debt)
					}
					k.node.SetPledgeDebt(ctx, pledgeDebt)
				}
				totalPledgeChange = totalPledgeChange.Add(extraPledge.Amount)

				shard.Pledge = newPledge

				pledge, _ := k.node.GetPledge(ctx, shard.Sp)
				pledge.TotalStoragePledged = pledge.TotalStoragePledged.Add(extraPledge)
				k.node.SetPledge(ctx, pledge)
			}

			renewInfo := ordertypes.RenewInfo{
				OrderId:  newOrder.Id,
				Pledge:   newPledge,
				Duration: proposal.Duration,
			}
			shard.RenewInfos = append(shard.RenewInfos, renewInfo)
			shardExpiredAt := shard.CreatedAt + shard.Duration
			for _, info := range shard.RenewInfos {
				shardExpiredAt += info.Duration
			}
			if shardExpiredAt > newExpiredAt {
				newExpiredAt = shardExpiredAt
			}

			k.order.SetShard(ctx, shard)
		}

		k.model.ExtendMetaDuration(ctx, metadata.DataId, newExpiredAt)
		k.model.UpdateMeta(ctx, newOrder)

		if !totalPledgeChange.IsZero() {
			pool.TotalPledged.Amount = pool.TotalPledged.Amount.Add(totalPledgeChange)
			k.node.SetPool(ctx, pool)
		}

		kv := &types.KV{
			K: dataId,
			V: fmt.Sprintf("SUCCESS: orderId=%d", newOrder.Id),
		}
		resp.Result = append(resp.Result, kv)
	}

	return &resp, nil
}

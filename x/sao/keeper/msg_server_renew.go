package keeper

import (
	"context"
	"fmt"

	markettypes "github.com/SaoNetwork/sao/x/market/types"
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
	if proposal.Owner != "all" {
		sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
		if err != nil {
			return nil, err
		}
	} else {
		sigDid = "all"
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

	ownerAddress, err := k.did.GetCosmosPaymentAddress(ctx, sigDid)
	if err != nil {
		return nil, err
	}
	denom := k.staking.BondDenom(ctx)
	balance := k.bank.GetBalance(ctx, ownerAddress, denom)

	blockRewardPerByte := sdk.NewDecWithPrec(1, 6)

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

		order, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrOrderNotFound, "FAILED: invalid order id: %d", metadata.OrderId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		duration := int64(order.Duration)
		currentHeight := ctx.BlockHeight()
		orderExpiredAt := int64(order.CreatedAt) + duration
		newOrderExpiredAt := currentHeight + int64(proposal.Duration)

		if orderExpiredAt < currentHeight {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrorOrderExpired, "FAILED: metadata should have expired: order has expired at %d", orderExpiredAt).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		if orderExpiredAt >= newOrderExpiredAt {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrorInvalidExpiredAt, "FAILED: invalid new expired height: expired at %d, but new expired at %d", orderExpiredAt, newOrderExpiredAt).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		//newOrder := ordertypes.Order{
		//	Creator:   msg.Creator,
		//	Owner:     order.Owner,
		//	Provider:  msg.Provider,
		//	Cid:       order.Cid,
		//	Duration:  proposal.Duration,
		//	Status:    order.Status,
		//	Replica:   order.Replica,
		//	Shards:    order.Shards,
		//	Size_:     order.Size_,
		//	Commit:    order.Commit,
		//	Operation: 3,
		//	CreatedAt: uint64(ctx.BlockHeight()),
		//	Timeout:   uint64(proposal.Timeout),
		//	DataId:    order.DataId,
		//	UnitPrice: order.UnitPrice,
		//}

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
			shards = append(shards, shard)
			//shard.Pledge.(currentHeight - order.CreatedAt)
		}

		orderPledge := order.UnitPrice.Amount.MulInt64(duration).MulInt64(int64(order.Replica)).MulInt64(int64(order.Size_))
		leftDec := sdk.NewDecCoinFromCoin(order.Amount).Amount.Sub(orderPledge)
		// calculate new amount
		newOrderPledge := order.UnitPrice.Amount.MulInt64(newOrderExpiredAt - orderExpiredAt).MulInt64(int64(order.Replica)).MulInt64(int64(order.Size_))
		newAmount := sdk.NewCoin(denom, sdk.NewInt(0))
		if leftDec.LT(newOrderPledge) {
			var dec sdk.DecCoin
			newAmount, dec = sdk.NewDecCoinFromDec(denom, newOrderPledge.Sub(leftDec)).TruncateDecimal()
			if !dec.IsZero() {
				newAmount = newAmount.AddAmount(sdk.NewInt(1))
			}
			if balance.IsLT(newAmount) {
				kv := &types.KV{
					K: dataId,
					V: sdkerrors.Wrapf(types.ErrInsufficientCoin, "FAILED: insufficient coin: need %d", newAmount.Amount.Int64()).Error(),
				}
				resp.Result = append(resp.Result, kv)
				continue
			} else {
				balance = balance.Sub(newAmount)
				k.bank.SendCoinsFromAccountToModule(ctx, ownerAddress, markettypes.ModuleName, sdk.Coins{newAmount})
			}
		}

		totalPledgeChange := sdk.NewInt(0)
		for _, shard := range shards {
			spAcc := sdk.MustAccAddressFromBech32(shard.Sp)

			blockRewardPledge := k.node.BlockRewardPledge(proposal.Duration, shard.Size_, sdk.NewDecCoinFromDec(denom, blockRewardPerByte))
			storeRewardPledge := k.node.StoreRewardPledge(proposal.Duration, shard.Size_, order.UnitPrice)

			newPledge, dec := sdk.NewDecCoinFromDec(denom, blockRewardPledge.Add(storeRewardPledge)).TruncateDecimal()
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
			} else if newPledge.Amount.LT(shard.Pledge.Amount) {
				refundPledge := shard.Pledge.Sub(newPledge)
				k.bank.SendCoinsFromModuleToAccount(ctx, nodetypes.ModuleName, spAcc, sdk.Coins{refundPledge})
				totalPledgeChange = totalPledgeChange.Sub(refundPledge.Amount)
			}
			shard.Pledge = newPledge
			k.order.SetShard(ctx, shard)
		}

		order.Amount = order.Amount.Add(newAmount)
		order.Duration = uint64(newOrderExpiredAt) - order.CreatedAt
		k.order.SetOrder(ctx, order)
		k.model.ExtendMetaDuration(ctx, metadata, uint64(newOrderExpiredAt))

		if !totalPledgeChange.IsZero() {
			pool.TotalPledged.Amount = pool.TotalPledged.Amount.Add(totalPledgeChange)
			k.node.SetPool(ctx, pool)
		}

		kv := &types.KV{
			K: dataId,
			V: fmt.Sprintf("SUCCESS: orderId=%d", order.Id),
		}
		resp.Result = append(resp.Result, kv)
	}

	return &resp, nil
}

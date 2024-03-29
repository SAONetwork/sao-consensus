package keeper

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/SaoNetwork/sao/x/model/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Version(commit string, height int64) string {
	version := bytes.NewBufferString(commit)
	version.WriteByte(26)
	version.WriteString(fmt.Sprintf("%d", height))
	return version.String()
}

func CommitFromVersion(version string) string {
	splited := strings.Split(version, string([]uint8{26}))
	return splited[0]
}

func (k Keeper) NewMeta(ctx sdk.Context, order ordertypes.Order, metadata types.Metadata) error {

	if len(metadata.DataId) != 36 {
		return sdkerrors.Wrapf(types.ErrInvalidDataId, "dataid: %s", metadata.DataId)
	}

	_, found_meta := k.GetMetadata(ctx, metadata.DataId)
	if found_meta {
		return sdkerrors.Wrap(types.ErrDataIdExists, "")
	}

	key := fmt.Sprintf("%s-%s-%s", metadata.Owner, metadata.Alias, metadata.GroupId)

	_, found_model := k.GetModel(ctx, key)
	if found_model {
		return sdkerrors.Wrapf(types.ErrModelExists, "model key: %s", key)
	}

	model := types.Model{
		Key:  key,
		Data: metadata.DataId,
	}

	k.SetModel(ctx, model)

	k.SetMetadata(ctx, metadata)

	k.setDataExpireBlock(ctx, metadata.DataId, order.CreatedAt+order.Duration)

	return nil
}

func (k Keeper) UpdateMeta(ctx sdk.Context, order ordertypes.Order) error {

	if len(order.DataId) != 36 {
		return sdkerrors.Wrapf(types.ErrInvalidDataId, "dataid: %s", order.DataId)
	}

	metadata, foundMeta := k.GetMetadata(ctx, order.DataId)
	if !foundMeta {
		return status.Error(codes.NotFound, "not found")
	}

	isValid := metadata.Owner == order.Owner
	if !isValid {
		for _, readwriteDid := range metadata.ReadwriteDids {
			if readwriteDid == order.Owner {
				isValid = true
				break
			}
		}

		if !isValid {
			return sdkerrors.Wrap(types.ErrorNoPermission, "No permission to update the model")
		}
	}

	switch order.Operation {
	case 1: // new or update

		metadata.Cid = order.Cid

		metadata.Commit = order.Commit
		metadata.Commits = append(metadata.Commits, Version(order.Commit, ctx.BlockHeight()))
		metadata.Orders = append(metadata.Orders, order.Id)
	case 2: // force push, replace last commit
		//lastOrderId := metadata.OrderId
		lastCommit := CommitFromVersion(metadata.Commits[len(metadata.Commits)-1])

		// old order settlement
		shardSet := make(map[uint64]int)
		for {
			if len(metadata.Orders) == 0 {
				break
			}
			lastOrder, foundLastOrder := k.order.GetOrder(ctx, metadata.Orders[len(metadata.Orders)-1])
			if !foundLastOrder {
				return status.Error(codes.NotFound, "last order not found")
			}
			if lastOrder.Commit != lastCommit {
				break
			}

			for _, shardId := range lastOrder.Shards {
				shardSet[shardId] = 1
			}

			err := k.TerminateOrder(ctx, lastOrder)
			if err != nil {
				return err
			}
			metadata.Orders = metadata.Orders[:len(metadata.Orders)-1]
		}

		for shardId, _ := range shardSet {
			k.order.RemoveShard(ctx, shardId)
		}

		// remove old version
		metadata.Cid = order.Cid
		if len(metadata.Commits) > 0 {
			metadata.Commits = metadata.Commits[:len(metadata.Commits)-1]
		}
		metadata.Commit = order.Commit
		metadata.Commits = append(metadata.Commits, Version(order.Commit, ctx.BlockHeight()))
		metadata.Orders = append(metadata.Orders, order.Id)
		k.ResetMetaDuration(ctx, &metadata)
	case 3: // renew
		metadata.OrderId = order.Id
		metadata.Orders = append(metadata.Orders, order.Id)
	default:
		return sdkerrors.Wrap(types.ErrInvalidOperation, "Operation should in [1, 2, 3]")
	}
	metadata.Status = types.MetaComplete
	k.SetMetadata(ctx, metadata)

	return nil
}

func (k Keeper) UpdateMetaStatusAndCommit(ctx sdk.Context, order ordertypes.Order) error {
	metadata, found := k.GetMetadata(ctx, order.DataId)
	if !found {
		return status.Errorf(codes.NotFound, "dataId %s not found", order.DataId)
	}

	if metadata.Status != types.MetaComplete {
		return sdkerrors.Wrapf(types.ErrInvalidStatus, "unexpected meta: %s, status: %d", metadata.DataId, metadata.Status)
	}

	// calculate new duration
	oldExpired := metadata.CreatedAt + metadata.Duration
	newExpired := order.CreatedAt + order.Duration
	if oldExpired < uint64(ctx.BlockHeight()) {
		return status.Error(codes.Aborted, "metadata should have expired")
	} else if oldExpired < newExpired {
		k.removeDataExpireBlock(ctx, metadata.DataId, oldExpired)
		metadata.Duration = newExpired - metadata.CreatedAt
		k.setDataExpireBlock(ctx, metadata.DataId, newExpired)
	}

	metadata.Status = int32(order.Operation)
	metadata.Commit = order.Commit
	metadata.OrderId = order.Id

	k.SetMetadata(ctx, metadata)
	return nil
}

func (k Keeper) DeleteMeta(ctx sdk.Context, dataId string) error {
	metadata, found := k.GetMetadata(ctx, dataId)
	if !found {
		return status.Errorf(codes.NotFound, "dataId %s not found", dataId)
	}

	key := fmt.Sprintf("%s-%s-%s", metadata.Owner, metadata.Alias, metadata.GroupId)
	k.RemoveMetadata(ctx, dataId)
	k.RemoveModel(ctx, key)

	return nil
}

func (k Keeper) UpdatePermission(ctx sdk.Context, owner string, dataId string, readonlyDids []string, readwriteDids []string) error {
	metadata, found := k.GetMetadata(ctx, dataId)
	if !found {
		return status.Errorf(codes.NotFound, "dataId %s not found", dataId)
	}

	if owner != metadata.Owner {
		return sdkerrors.Wrap(types.ErrorNoPermission, "No permission to update the model")
	}

	metadata.ReadonlyDids = readonlyDids
	metadata.ReadwriteDids = readwriteDids

	k.SetMetadata(ctx, metadata)

	return nil
}

func (k Keeper) setDataExpireBlock(ctx sdk.Context, dataId string, expiredAt uint64) {

	expiredData, foundExpiredData := k.GetExpiredData(ctx, expiredAt)

	if !foundExpiredData {
		expiredData = types.ExpiredData{
			Height: expiredAt,
		}
	}

	expiredData.Data = append(expiredData.Data, dataId)

	k.SetExpiredData(ctx, expiredData)
}

// TODO: consider in which other cases should remove data expire block
func (k Keeper) removeDataExpireBlock(ctx sdk.Context, dataId string, expiredAt uint64) {

	expiredData, foundExpiredData := k.GetExpiredData(ctx, expiredAt)

	if !foundExpiredData {
		return
	}

	for idx, id := range expiredData.Data {
		if id == dataId {
			expiredData.Data = append(expiredData.Data[:idx], expiredData.Data[idx+1:]...)
		}
	}

	if len(expiredData.Data) == 0 {
		k.RemoveExpiredData(ctx, expiredData.Height)
	} else {
		k.SetExpiredData(ctx, expiredData)
	}
}

func (k Keeper) TerminateOrder(ctx sdk.Context, order ordertypes.Order) error {
	refund, err := k.market.Withdraw(ctx, order)
	if err != nil {
		return err
	}

	// change pledge and pool status
	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			continue
		}
		if shard.Status == ordertypes.ShardCompleted && shard.OrderId == order.Id {
			err := k.node.ShardRelease(ctx, sdk.MustAccAddressFromBech32(shard.Sp), &shard)
			if err != nil {
				return err
			}
		}
	}

	err = k.order.TerminateOrder(ctx, order.Id, refund)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) CancelOrder(ctx sdk.Context, orderId uint64) error {

	order, _ := k.order.GetOrder(ctx, orderId)

	if k.order.RefundOrder(ctx, orderId) != nil {
		return sdkerrors.Wrapf(ordertypes.ErrorRefundOrder, "refund order failed")
	}

	k.RollbackMeta(ctx, order.DataId)
	k.order.RemoveOrder(ctx, orderId)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(ordertypes.CancelOrderEventType,
			sdk.NewAttribute(ordertypes.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	return nil
}

func (k Keeper) RollbackMeta(ctx sdk.Context, dataId string) {

	metadata, found := k.GetMetadata(ctx, dataId)
	if !found {
		return
	}

	if len(metadata.Commits) == 0 {
		k.RemoveMetadata(ctx, dataId)

		key := fmt.Sprintf("%s-%s-%s", metadata.Owner, metadata.Alias, metadata.GroupId)
		k.RemoveModel(ctx, key)
		return
	}

	metadata.Status = types.MetaComplete
	metadata.Commit = CommitFromVersion(metadata.Commits[len(metadata.Commits)-1])
	metadata.OrderId = metadata.Orders[len(metadata.Orders)-1]
	k.ResetMetaDuration(ctx, &metadata)

	k.SetMetadata(ctx, metadata)
	return
}

func (k Keeper) ResetMetaDuration(ctx sdk.Context, meta *types.Metadata) {
	orders := meta.Orders

	var expiredHeight uint64 = 0
	shardExpiredMap := make(map[uint64]uint64)
	for _, orderId := range orders {
		order, foundOrder := k.order.GetOrder(ctx, orderId)
		if foundOrder {
			for _, shardId := range order.Shards {
				if shardExpiredMap[shardId] == 0 {
					shard, foundShard := k.order.GetShard(ctx, shardId)
					if foundShard && shard.Status == ordertypes.ShardCompleted {
						shardExpiredMap[shardId] = shard.CreatedAt + shard.Duration
						for _, renewInfo := range shard.RenewInfos {
							shardExpiredMap[shardId] += renewInfo.Duration
						}
						if shardExpiredMap[shardId] > expiredHeight {
							expiredHeight = shardExpiredMap[shardId]
						}
					}
				}
			}
		}
	}

	newDuration := expiredHeight - meta.CreatedAt

	if meta.Duration != newDuration {
		k.removeDataExpireBlock(ctx, meta.DataId, meta.CreatedAt+meta.Duration)
		meta.Duration = newDuration
		k.setDataExpireBlock(ctx, meta.DataId, expiredHeight)
	}
}

func (k Keeper) ExtendMetaDuration(ctx sdk.Context, dataId string, expiredAt uint64) {
	meta, _ := k.GetMetadata(ctx, dataId)
	newDuration := expiredAt - meta.CreatedAt
	if meta.Duration < newDuration {
		k.removeDataExpireBlock(ctx, meta.DataId, meta.CreatedAt+meta.Duration)
		meta.Duration = newDuration
		k.setDataExpireBlock(ctx, meta.DataId, expiredAt)
		k.SetMetadata(ctx, meta)
	}
}

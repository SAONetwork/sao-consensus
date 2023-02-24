package keeper

import (
	"bytes"
	"fmt"

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

func (k Keeper) NewMeta(ctx sdk.Context, order ordertypes.Order) error {

	metadata := types.Metadata{
		DataId:     order.Metadata.DataId,
		Owner:      order.Metadata.Owner,
		Alias:      order.Metadata.Alias,
		GroupId:    order.Metadata.GroupId,
		OrderId:    order.Metadata.OrderId,
		Tags:       order.Metadata.Tags,
		Cid:        order.Metadata.Cid,
		ExtendInfo: order.Metadata.ExtendInfo,
		Commit:     order.Metadata.Commit,
		Rule:       order.Metadata.Rule,
		Duration:   order.Metadata.Duration,
		// TODO: ensure createdAt is synced with order.Metadata
		CreatedAt: order.Metadata.CreatedAt,
	}

	if len(metadata.DataId) != 36 {
		return sdkerrors.Wrapf(types.ErrInvalidDataId, "dataid: %s", metadata.DataId)
	}

	_, found_meta := k.GetMetadata(ctx, metadata.DataId)
	if found_meta {
		return sdkerrors.Wrap(types.ErrDataIdExists, "")
	}

	key := fmt.Sprintf("%s-%s-%s", order.Owner, metadata.Alias, metadata.GroupId)

	metadata.Owner = order.Owner

	metadata.OrderId = order.Id

	metadata.Commits = make([]string, 0)

	metadata.Commits = append(metadata.Commits, Version(metadata.DataId, ctx.BlockHeight()))

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

	expiredAt := metadata.CreatedAt + metadata.Duration

	k.setDataExpireBlock(ctx, metadata.DataId, expiredAt)

	return nil
}

func (k Keeper) UpdateMeta(ctx sdk.Context, order ordertypes.Order) error {

	metadata := types.Metadata{
		DataId:     order.Metadata.DataId,
		Owner:      order.Metadata.Owner,
		Alias:      order.Metadata.Alias,
		GroupId:    order.Metadata.GroupId,
		OrderId:    order.Metadata.OrderId,
		Tags:       order.Metadata.Tags,
		Cid:        order.Metadata.Cid,
		ExtendInfo: order.Metadata.ExtendInfo,
		Commit:     order.Metadata.Commit,
		Rule:       order.Metadata.Rule,
		Duration:   order.Metadata.Duration,
		CreatedAt:  order.Metadata.CreatedAt,
	}

	if len(metadata.DataId) != 36 {
		return sdkerrors.Wrapf(types.ErrInvalidDataId, "dataid: %s", metadata.DataId)
	}

	_metadata, found_meta := k.GetMetadata(ctx, metadata.DataId)
	if !found_meta {
		return status.Error(codes.NotFound, "not found")
	}

	isValid := _metadata.Owner == order.Owner
	if !isValid {
		for _, readwriteDid := range _metadata.ReadwriteDids {
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
	case 0:
		return sdkerrors.Wrap(types.ErrInvalidOperation, "Operation should in [1, 2, 3]")
	case 1: // new or update
		_metadata.OrderId = order.Id

		_metadata.Cid = metadata.Cid

		_metadata.Commit = metadata.Commit
		_metadata.Commits = append(_metadata.Commits, Version(metadata.Commit, ctx.BlockHeight()))

		k.SetMetadata(ctx, _metadata)
	case 2: // force push, replace last commit
		lastOrder := _metadata.OrderId

		// remove old expired block
		oldExpired := _metadata.CreatedAt + _metadata.Duration
		if oldExpired > uint64(ctx.BlockHeight()) {
			k.removeDataExpireBlock(ctx, _metadata.DataId, oldExpired)
		}

		// TODO: consider refund
		err := k.order.TerminateOrder(ctx, lastOrder, sdk.Coin{})
		if err != nil {
			return err
		}

		// remove old version
		_metadata.OrderId = order.Id
		_metadata.Cid = metadata.Cid
		if len(_metadata.Commits) > 0 {
			_metadata.Commits = _metadata.Commits[:len(_metadata.Commits)-1]
		}
		_metadata.Commit = metadata.Commit
		_metadata.Commits = append(_metadata.Commits, Version(metadata.Commit, ctx.BlockHeight()))

		k.SetMetadata(ctx, _metadata)
	case 3: // renew
		lastOrder := _metadata.OrderId

		// remove old expired Block
		oldExpired := _metadata.CreatedAt + _metadata.Duration
		if oldExpired > uint64(ctx.BlockHeight()) {
			k.removeDataExpireBlock(ctx, _metadata.DataId, oldExpired)
		}

		// TODO: consider refund
		k.order.TerminateOrder(ctx, lastOrder, sdk.Coin{})

		_metadata.OrderId = order.Id

		k.SetMetadata(ctx, _metadata)
	}

	expiredAt := metadata.CreatedAt + metadata.Duration
	k.setDataExpireBlock(ctx, metadata.DataId, expiredAt)

	return nil
}

func (k Keeper) SettlementWithMeta(ctx sdk.Context, dataId string) error {

	metadata, found := k.GetMetadata(ctx, dataId)
	if !found {
		return status.Errorf(codes.NotFound, "dataId %s not found", dataId)
	}

	order, found := k.order.GetOrder(ctx, metadata.OrderId)
	if !found {
		return status.Errorf(codes.NotFound, "orderId %s not found", metadata.OrderId)
	}

	// change worker status
	refund, err := k.market.Withdraw(ctx, order)
	if err != nil {
		return err
	}
	if !refund.IsZero() {
		return status.Errorf(codes.Aborted, "refund should be zero when withdraw from a finished order")
	}

	// change pledge and pool status
	for sp, _ := range order.Shards {
		err := k.node.OrderRelease(ctx, sdk.MustAccAddressFromBech32(sp), &order)
		if err != nil {
			return err
		}
	}

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

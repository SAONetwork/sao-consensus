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

	metadata.CreatedAt = uint64(ctx.BlockTime().Unix())

	k.SetModel(ctx, model)

	k.SetMetadata(ctx, metadata)

	k.setDataExpireBlock(ctx, metadata.DataId, order.Duration)

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

		err := k.order.TerminateOrder(ctx, lastOrder)
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

		k.order.TerminateOrder(ctx, lastOrder)

		_metadata.OrderId = order.Id

		k.SetMetadata(ctx, _metadata)
	}

	k.setDataExpireBlock(ctx, metadata.DataId, order.Duration)

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

func (k Keeper) setDataExpireBlock(ctx sdk.Context, dataId string, duration uint64) {

	expiredAt := ctx.BlockHeight() + int64(duration)

	expiredData, foundExpiredData := k.GetExpiredData(ctx, uint64(expiredAt))

	if !foundExpiredData {
		expiredData = types.ExpiredData{
			Height: uint64(expiredAt),
		}
	}

	expiredData.Data = append(expiredData.Data, dataId)

	k.SetExpiredData(ctx, expiredData)
}

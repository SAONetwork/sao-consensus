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

	metadata := order.Metadata

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
		return sdkerrors.Wrapf(types.ErrModelExists, "modek key: %s", key)
	}

	model := types.Model{
		Key:  key,
		Data: metadata.DataId,
	}

	metadata.CreatedAt = ctx.BlockTime().Unix()

	metadata.Amount = order.Amount

	k.SetModel(ctx, model)

	k.SetMetadata(ctx, metadata)

	return nil
}

func (k Keeper) UpdateMeta(ctx sdk.Context, order ordertypes.Order) error {

	metadata := order.Metadata

	if len(metadata.DataId) != 36 {
		return sdkerrors.Wrapf(types.ErrInvalidDataId, "dataid: %s", metadata.DataId)
	}

	_metadata, found_meta := k.GetMetadata(ctx, metadata.DataId)
	if !found_meta {
		return status.Error(codes.NotFound, "not found")
	}

	if _metadata.Owner != order.Owner {
		return sdkerrors.Wrap(types.ErrOnlyOwner, "")
	}

	if order.Force {
		// remove old version

	} else {
		_metadata.OrderId = order.Id

		_metadata.Cid = metadata.Cid

		_metadata.Commits = append(_metadata.Commits, Version(metadata.Commit, ctx.BlockHeight()))

		k.SetMetadata(ctx, _metadata)
	}

	return nil
}

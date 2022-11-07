package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/SaoNetwork/sao/x/model/types"
	saotypes "github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NewMeta(ctx sdk.Context, order saotypes.Order) error {
	var metadata types.Metadata
	err := json.Unmarshal([]byte(order.Metadata), &metadata)

	if err != nil {
		return err
	}

	if len(metadata.DataId) != 36 {
		return sdkerrors.Wrapf(types.ErrInvalidDataId, "dataid: %s", metadata.DataId)
	}

	if metadata.Update {

		_metadata, found_meta := k.GetMetadata(ctx, metadata.DataId)
		if !found_meta {
			return status.Error(codes.NotFound, "not found")
		}

		if _metadata.Owner != order.Owner {
			return sdkerrors.Wrap(types.ErrOnlyOwner, "")
		}

		_metadata.Cids = metadata.Cids

		k.SetMetadata(ctx, _metadata)

	} else {

		_, found_meta := k.GetMetadata(ctx, metadata.DataId)
		if found_meta {
			return sdkerrors.Wrap(types.ErrDataIdExists, "")
		}

		key := fmt.Sprintf("%s-%s-%s", order.Owner, metadata.Alias, metadata.GroupId)

		metadata.Owner = order.Owner

		metadata.OrderId = order.Id

		_, found_model := k.GetModel(ctx, key)
		if found_model {
			return sdkerrors.Wrapf(types.ErrModelExists, "modek key: %s", key)
		}

		model := types.Model{
			Key:  key,
			Data: metadata.DataId,
		}

		k.SetModel(ctx, model)

		k.SetMetadata(ctx, metadata)

		return nil
	}
	return nil
}

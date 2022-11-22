package keeper

import (
	"context"

	saodid "github.com/SaoNetwork/sao-did"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/ipfs/go-cid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal := msg.Proposal

	didManager, err := saodid.NewDidManagerWithDid(proposal.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidDid, "")
	}

	proposalBytes, _ := proposal.Marshal()

	signature := saodidtypes.JwsSignature{
		Protected: msg.JwsSignature.Protected,
		Signature: msg.JwsSignature.Signature,
	}

	_, err = didManager.VerifyJWS(saodidtypes.GeneralJWS{
		Payload: base64url.Encode(proposalBytes),
		Signatures: []saodidtypes.JwsSignature{
			signature,
		},
	})

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
	}

	var metadata *ordertypes.Metadata
	var model modeltypes.Metadata
	var found_model bool
	var node nodetypes.Node

	if proposal != nil {
		if proposal.Operation != 0 {
			model, found_model = k.Keeper.model.GetMetadata(ctx, proposal.DataId)
			if !found_model {
				return nil, status.Errorf(codes.NotFound, "dataId %d not found", proposal.DataId)
			}
		}
		if proposal.CommitId != "" && proposal.Operation != 2 {
			metadata = &ordertypes.Metadata{
				DataId:     proposal.DataId,
				Owner:      proposal.Owner,
				Alias:      proposal.Alias,
				GroupId:    proposal.GroupId,
				Tags:       proposal.Tags,
				Cid:        proposal.Cid,
				Commit:     proposal.CommitId,
				ExtendInfo: proposal.ExtendInfo,
				Rule:       proposal.Rule,
			}
		}

		if metadata != nil {
			if metadata.DataId == "" {
				return nil, sdkerrors.Wrap(types.ErrorInvalidDataId, "")
			}
			// check provider
			node, found := k.node.GetNode(ctx, proposal.Provider)
			if !found {
				return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
			}

			// check cid
			_, err := cid.Decode(proposal.Cid)
			if err != nil {
				return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", proposal.Cid)
			}
		}
	}

	if proposal.Size_ == 0 {
		proposal.Size_ = 1
	}

	var order = ordertypes.Order{
		Creator:   msg.Creator,
		Owner:     proposal.Owner,
		Cid:       proposal.Cid,
		Expire:    proposal.Timeout + int32(ctx.BlockHeight()),
		Duration:  proposal.Duration,
		Status:    types.OrderPending,
		Replica:   proposal.Replica,
		Metadata:  metadata,
		Operation: int32(proposal.Operation),
	}

	if node.Creator != "" {
		order.Provider = node.Creator
	}

	var sps []nodetypes.Node

	if order.Provider == msg.Creator || order.Operation == 2 {
		if order.Operation == 0 {
			sps = k.node.RandomSP(ctx, order)
			if order.Replica <= 0 || int(order.Replica) > len(sps) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
			}
		} else if order.Operation > 0 {
			sps = k.FindSPByDataId(ctx, proposal.DataId)
			if order.Operation == 2 {
				oldOrder, _ := k.order.GetOrder(ctx, model.OrderId)
				order.Size_ = oldOrder.Size_
				order.Replica = oldOrder.Replica
			}
			if len(sps) == 0 {
				return nil, sdkerrors.Wrap(types.ErrorInvalidDataId, "")
			}
		}
	}

	if order.Size_ == 0 {
		order.Size_ = 1
	}

	price := sdk.NewInt(1)

	owner_address := k.did.GetCosmosPaymentAddress(ctx, proposal.Owner)

	amount := sdk.NewCoin(sdk.DefaultBondDenom, price.MulRaw(int64(order.Size_)).MulRaw(int64(order.Replica)))
	balance := k.bank.GetBalance(ctx, owner_address, sdk.DefaultBondDenom)

	logger := k.Logger(ctx)

	logger.Debug("order amount1 ###################", "amount", amount, "owner", owner_address, "balance", balance)

	if balance.IsLT(amount) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", amount.Amount.Int64())
	}

	order.Amount = amount

	logger.Debug("########### choose sp ", "sps", sps)

	orderId, err := k.order.NewOrder(ctx, order, sps)
	if err != nil {
		return nil, err
	}

	return &types.MsgStoreResponse{OrderId: orderId}, nil
}

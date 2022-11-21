package keeper

import (
	"context"

	saodid "github.com/SaoNetwork/sao-did"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/ipfs/go-cid"
)

func getAddress() sdk.AccAddress {
	return sdk.MustAccAddressFromBech32("cosmos1r33rtwtgak2erkwq2462l3ed2ry2q0p0427eu9")
}

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal := msg.Proposal

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

	if proposal != nil {
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

	if metadata == nil || metadata.DataId == "" {
		return nil, sdkerrors.Wrap(types.ErrorInvalidDataId, "")
	}
	if proposal.Size_ == 0 {
		proposal.Size_ = 1
	}

	price := sdk.NewInt(1)

	owner_address := getAddress()

	amount := sdk.NewCoin(sdk.DefaultBondDenom, price.MulRaw(int64(proposal.Size_)).MulRaw(int64(proposal.Replica)))
	balance := k.bank.GetBalance(ctx, owner_address, sdk.DefaultBondDenom)

	logger := k.Logger(ctx)

	logger.Debug("order amount ###################", "amount", amount, "owner", owner_address, "balance", balance)

	if balance.IsLT(amount) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", amount.Amount.Int64())
	}

	var order = ordertypes.Order{
		Creator:  msg.Creator,
		Owner:    proposal.Owner,
		Provider: node.Creator,
		Cid:      proposal.Cid,
		Expire:   proposal.Timeout + int32(ctx.BlockHeight()),
		Duration: proposal.Duration,
		Status:   types.OrderPending,
		Replica:  proposal.Replica,
		Metadata: metadata,
		Size_:    proposal.Size_,
		Amount:   amount,
	}

	var sps []nodetypes.Node

	if order.Provider == msg.Creator || order.Operation == 2 {
		if order.Operation == 0 {
			sps = k.node.RandomSP(ctx, order)
		} else if order.Operation > 0 {
			sps = k.FindSPByDataId(ctx, metadata.DataId)
			if len(sps) == 0 {
				return nil, sdkerrors.Wrap(types.ErrorInvalidDataId, "")
			}
		}
		if order.Replica <= 0 || int(order.Replica) > len(sps) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
		}
	}

	orderId, err := k.order.NewOrder(ctx, order, sps)
	if err != nil {
		return nil, err
	}

	return &types.MsgStoreResponse{OrderId: orderId}, nil
}

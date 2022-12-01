package keeper

import (
	"context"
	"encoding/json"

	saodid "github.com/SaoNetwork/sao-did"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
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

const addressPrefix = "cosmos"
const chainAddress = "http://localhost:26657"

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal := &msg.Proposal

	if proposal == nil {
		return nil, status.Errorf(codes.InvalidArgument, "proposal is required")
	}

	if &msg.JwsSignature == nil {
		return nil, status.Errorf(codes.InvalidArgument, "signature is required")
	}

	didManager, err := saodid.NewDidManagerWithDid(proposal.Owner, addressPrefix, chainAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidDid, "")
	}

	proposalBytes, err := json.Marshal(proposal)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidProposal, "")
	}

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
	var found_model bool
	var found_node bool
	var node nodetypes.Node

	if proposal.Operation != 0 {
		_, found_model = k.Keeper.model.GetMetadata(ctx, proposal.DataId)
		if !found_model {
			return nil, status.Errorf(codes.NotFound, "dataId %d not found", proposal.DataId)
		}
	}

	if proposal.CommitId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid commitIdi")
	}

	if proposal.DataId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid dataId")
	}

	// check cid
	_, err = cid.Decode(proposal.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", proposal.Cid)
	}

	// check provider
	node, found_node = k.node.GetNode(ctx, proposal.Provider)
	if !found_node {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

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

	if order.Provider == msg.Creator {
		if order.Operation == 0 {
			sps = k.node.RandomSP(ctx, order)
			if order.Replica <= 0 || int(order.Replica) > len(sps) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
			}
		} else if order.Operation > 0 {
			sps = k.FindSPByDataId(ctx, proposal.DataId)
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

	logger.Error("order amount1 ###################", "amount", amount, "owner", owner_address, "balance", balance)

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

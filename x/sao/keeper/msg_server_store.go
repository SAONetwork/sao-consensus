package keeper

import (
	"context"
	"github.com/ipfs/go-cid"
	"strings"

	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	var sigDid string
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &msg.Proposal
	if proposal == nil {
		return nil, status.Errorf(codes.InvalidArgument, "proposal is required")
	}

	sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
	if err != nil {
		return nil, err
	}

	var metadata modeltypes.Metadata
	var node nodetypes.Node

	if proposal.CommitId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid commitId")
	}

	if proposal.DataId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid dataId")
	}

	if proposal.Operation < 1 || proposal.Operation > 2 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid operation %d", proposal.Operation)
	}

	if proposal.Duration < 3600 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid duration")
	}

	// check cid
	_, err = cid.Decode(proposal.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", proposal.Cid)
	}

	if !strings.Contains(proposal.CommitId, proposal.DataId) {
		// validate the permission for all update operations
		meta, isFound := k.Keeper.model.GetMetadata(ctx, proposal.DataId)
		if !isFound {
			return nil, status.Errorf(codes.NotFound, "metadata :%s not found", proposal.DataId)
		}

		isValid := meta.Owner == sigDid
		if !isValid {
			for _, readwriteDid := range meta.ReadwriteDids {
				if readwriteDid == sigDid {
					isValid = true
					break
				}
			}

			if !isValid {
				return nil, sdkerrors.Wrap(types.ErrorNoPermission, "No permission to update the model")
			}
		}
	}

	var paymentAddress sdk.AccAddress
	if proposal.PaymentDid != "" {
		if !strings.HasPrefix(proposal.PaymentDid, "did:key:") {
			return nil, sdkerrors.Wrapf(types.ErrorNotKid, "got payment did %s", proposal.PaymentDid)
		}
		paymentAddress, err = k.did.GetCosmosPaymentAddress(ctx, proposal.PaymentDid)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "invalid payment did")
		}
		if paymentAddress.String() != msg.Creator {
			return nil, sdkerrors.Wrap(types.ErrorNoPermission, "creator should be payment address of payment DID")
		}
	}

	// check provider
	node, found := k.node.GetNode(ctx, proposal.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	commitId := proposal.CommitId
	lastCommitId := proposal.CommitId
	if strings.Contains(proposal.CommitId, "|") {
		// extract the base version from the proposal
		lastCommitId = strings.Split(proposal.CommitId, "|")[0]
		commitId = strings.Split(proposal.CommitId, "|")[1]
	}

	if proposal.Size_ == 0 {
		proposal.Size_ = 1
	}

	if proposal.Timeout == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid arguments: timeout")
	}

	var order = ordertypes.Order{
		Creator:   msg.Creator,
		Owner:     proposal.Owner,
		Cid:       proposal.Cid,
		Timeout:   uint64(proposal.Timeout),
		Duration:  proposal.Duration,
		Status:    ordertypes.OrderPending,
		Replica:   proposal.Replica,
		DataId:    proposal.DataId,
		Operation: proposal.Operation,
		Size_:     proposal.Size_,
		Commit:    commitId,
	}

	if node.Creator != "" {
		order.Provider = node.Creator
	}

	if proposal.PaymentDid != "" {
		order.PaymentDid = proposal.PaymentDid
	}

	var sps []nodetypes.Node

	isProvider := false

	if paymentAddress.Empty() {
		err = k.did.CreatorIsBoundToDid(ctx, msg.Creator, proposal.Owner)
		if err != nil {
			if order.Provider == msg.Creator && msg.Provider == msg.Creator {
				isProvider = true
			} else if order.Provider == msg.Provider {
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
				return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, msg.Provider: %s", msg.Creator, order.Provider)
			}
		}
	} else {
		isProvider = true
	}

	if isProvider {
		sps, err = k.GetSps(ctx, order, proposal.DataId)
		if err != nil {
			return nil, err
		}
	}

	if order.Size_ == 0 {
		order.Size_ = 1
	}

	denom := k.staking.BondDenom(ctx)
	price := sdk.NewDecWithPrec(1, 6)
	unitPrice := sdk.NewDecCoinFromDec(denom, price)
	order.UnitPrice = unitPrice

	amount, dec := sdk.NewDecCoinFromDec(denom, price.MulInt64(int64(order.Size_)).MulInt64(int64(order.Replica)).MulInt64(int64(order.Duration))).TruncateDecimal()
	if !dec.IsZero() {
		amount = amount.AddAmount(sdk.NewInt(1))
	}

	if paymentAddress.Empty() {
		paymentAddress, err = k.did.GetCosmosPaymentAddress(ctx, proposal.Owner)
		if err != nil {
			return nil, err
		}
	}

	balance := k.bank.GetBalance(ctx, paymentAddress, denom)

	if balance.IsLT(amount) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insufficient coin: need %d", amount.Amount.Int64())
	}

	err = k.bank.SendCoinsFromAccountToModule(ctx, paymentAddress, ordertypes.ModuleName, sdk.Coins{amount})
	if err != nil {
		return nil, err
	}

	order.Amount = amount

	spCreators := make([]string, 0)
	for _, sp := range sps {
		spCreators = append(spCreators, sp.Creator)
	}

	orderId, err := k.order.NewOrder(ctx, &order, spCreators)
	if err != nil {
		return nil, err
	}

	if isProvider {
		k.SetTimeoutOrderBlock(ctx, order, order.CreatedAt+order.Timeout)
	}

	// avoid version conflicts
	meta, found := k.model.GetMetadata(ctx, proposal.DataId)
	if found {
		if meta.OrderId > orderId {
			// report error if order id is less than the latest version
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidCommitId, "invalid commitId: %s, detected version conflicts with order: %d", commitId, meta.OrderId)
		}

		lastOrder, isFound := k.order.GetOrder(ctx, meta.OrderId)
		if isFound {
			if lastOrder.Status != ordertypes.OrderCompleted {
				return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "unexpected last order: %s, status: %d", meta.OrderId, lastOrder.Status)
			}
		} else {
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "invalid last order: %s", meta.OrderId)
		}

		if !strings.Contains(meta.Commit, lastCommitId) {
			// report error if base version is not the latest version
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidCommitId, "invalid commitId: %s, detected version conficts, should be %s", lastCommitId, meta.Commit[:36])
		}

		// set meta status and commit
		err = k.model.UpdateMetaStatusAndCommit(ctx, order)
		if err != nil {
			return nil, err
		}

	} else {
		// new metadata
		metadata = modeltypes.Metadata{
			DataId:       proposal.DataId,
			Owner:        proposal.Owner,
			Alias:        proposal.Alias,
			GroupId:      proposal.GroupId,
			OrderId:      orderId,
			Tags:         proposal.Tags,
			Cid:          proposal.Cid,
			ExtendInfo:   proposal.ExtendInfo,
			Commit:       commitId,
			Rule:         proposal.Rule,
			Duration:     proposal.Duration,
			CreatedAt:    uint64(ctx.BlockHeight()),
			Status:       modeltypes.MetaNew,
			ReadonlyDids: proposal.ReadonlyDids,
		}

		err := k.model.NewMeta(ctx, order, metadata)
		if err != nil {
			return nil, err
		}
	}

	if isProvider {
		shards := make([]*types.ShardMeta, 0)
		for _, id := range order.Shards {
			shard, found := k.order.GetShard(ctx, id)
			if !found {
				return nil, status.Errorf(codes.NotFound, "shard %d not found", id)
			}
			node, node_found := k.node.GetNode(ctx, shard.Sp)
			if !node_found {
				continue
			}
			meta := types.ShardMeta{
				ShardId:  shard.Id,
				Peer:     node.Peer,
				Cid:      shard.Cid,
				Provider: order.Provider,
				Sp:       shard.Sp,
			}
			shards = append(shards, &meta)
		}

		return &types.MsgStoreResponse{
			OrderId: orderId,
			Shards:  shards,
		}, nil
	} else {
		return &types.MsgStoreResponse{
			OrderId: orderId,
		}, nil
	}

}

func (k Keeper) GetSps(ctx sdk.Context, order ordertypes.Order, dataId string) (sps []nodetypes.Node, err error) {

	if order.Operation == 1 {
		sps = k.node.RandomSP(ctx, int(order.Replica), nil, int64(order.Size_))
		if order.Replica <= 0 || int(order.Replica) > len(sps) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
		}
	} else if order.Operation == 2 {
		if order.Replica <= 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0")
		}
		sps = k.FindSPByDataId(ctx, dataId)
		if order.Replica < int32(len(sps)) {
			sps = sps[:order.Replica]
		} else if order.Replica > int32(len(sps)) {
			ignoreList := make([]string, 0)
			for _, sp := range sps {
				ignoreList = append(ignoreList, sp.Creator)
			}
			addSps := k.node.RandomSP(ctx, int(order.Replica)-len(sps), ignoreList, int64(order.Size_))
			sps = append(sps, addSps...)
		}
		if int(order.Replica) > len(sps) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should <= %d", len(sps))
		}
	} else {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidOperation, "unsupported operation %d", order.Operation)
	}
	return sps, nil
}

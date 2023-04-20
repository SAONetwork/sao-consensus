package keeper

import (
	"fmt"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		auth       types.AccountKeeper
		bank       types.BankKeeper
		node       types.NodeKeeper
		order      types.OrderKeeper
		model      types.ModelKeeper
		did        types.DidKeeper
		market     types.MarketKeeper
		staking    types.StakingKeeper
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	auth types.AccountKeeper,
	bank types.BankKeeper,
	node types.NodeKeeper,
	order types.OrderKeeper,
	model types.ModelKeeper,
	did types.DidKeeper,
	market types.MarketKeeper,
	staking types.StakingKeeper,
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		auth:       auth,
		bank:       bank,
		node:       node,
		order:      order,
		model:      model,
		did:        did,
		market:     market,
		staking:    staking,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) FindShardsByOrderId(ctx sdk.Context, orderId uint64) []ordertypes.Shard {
	shards := make([]ordertypes.Shard, 0)

	order, found := k.order.GetOrder(ctx, orderId)

	if !found {
		return shards
	}

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			continue
		}
		shards = append(shards, shard)
	}

	return shards
}

func (k Keeper) FindSPByDataId(ctx sdk.Context, dataId string) []nodetypes.Node {
	nodes := make([]nodetypes.Node, 0)

	model, found := k.model.GetMetadata(ctx, dataId)
	if !found {
		return nodes
	}

	order, found := k.order.GetOrder(ctx, model.OrderId)

	if !found {
		return nodes
	}

	for _, id := range order.Shards {
		shard, found := k.order.GetShard(ctx, id)
		if !found {
			continue
		}
		node, found := k.node.GetNode(ctx, shard.Sp)
		if found {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

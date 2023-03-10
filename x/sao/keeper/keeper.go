package keeper

import (
	"fmt"

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

	shards := k.node.GetMetadataShards(ctx, order.Metadata.DataId, int(order.Replica))

	for _, shard := range shards {
		node, found := k.node.GetNode(ctx, shard.Node)
		if found {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

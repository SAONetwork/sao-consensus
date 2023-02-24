package order

import (
	"github.com/SaoNetwork/sao/x/order/keeper"
	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the expiredOrder
	for _, elem := range genState.ExpiredOrderList {
		k.SetExpiredOrder(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the order
	for _, elem := range genState.OrderList {
		k.SetOrder(ctx, elem)
	}

	// Set order count
	k.SetOrderCount(ctx, genState.OrderCount)
	// Set all the shard
	for _, elem := range genState.ShardList {
		k.SetShard(ctx, elem)
	}

	// Set shard count
	k.SetShardCount(ctx, genState.ShardCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.OrderList = k.GetAllOrder(ctx)
	genesis.OrderCount = k.GetOrderCount(ctx)
	genesis.ShardList = k.GetAllShard(ctx)
	genesis.ShardCount = k.GetShardCount(ctx)
	genesis.ExpiredOrderList = k.GetAllExpiredOrder(ctx)
	// this line is used by starport scaffolding # genesis/module/export
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

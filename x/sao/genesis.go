package sao

import (
	"github.com/SaoNetwork/sao/x/sao/keeper"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the timeoutOrder
	for _, elem := range genState.TimeoutOrderList {
		k.SetTimeoutOrder(ctx, elem)
	}
	// Set all the expiredShard
	for _, elem := range genState.ExpiredShardList {
		k.SetExpiredShard(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TimeoutOrderList = k.GetAllTimeoutOrder(ctx)
	genesis.ExpiredShardList = k.GetAllExpiredShard(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

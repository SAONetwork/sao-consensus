package earn

import (
	"github.com/SaoNetwork/sao/x/earn/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	// Set if defined
	if genState.Pool != nil {
		k.SetPool(ctx, *genState.Pool)
	}
	// Set all the pledge
	for _, elem := range genState.PledgeList {
		k.SetPledge(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all pool
	pool, found := k.GetPool(ctx)
	if found {
		genesis.Pool = &pool
	}
	genesis.PledgeList = k.GetAllPledge(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

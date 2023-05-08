package node

import (
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the node
	for _, elem := range genState.NodeList {
		k.SetNode(ctx, elem)
	}
	// Set all the pledgeDebt
	for _, elem := range genState.PledgeDebtList {
		k.SetPledgeDebt(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	k.SetPool(ctx, *genState.Pool)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.NodeList = k.GetAllNode(ctx)
	genesis.PledgeDebtList = k.GetAllPledgeDebt(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

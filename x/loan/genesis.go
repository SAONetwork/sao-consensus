package loan

import (
	"github.com/SaoNetwork/sao/x/loan/keeper"
	"github.com/SaoNetwork/sao/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.LoanPool != nil {
		k.SetLoanPool(ctx, *genState.LoanPool)
	}
	// Set all the credit
	for _, elem := range genState.CreditList {
		k.SetCredit(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all loanPool
	loanPool, found := k.GetLoanPool(ctx)
	if found {
		genesis.LoanPool = &loanPool
	}
	genesis.CreditList = k.GetAllCredit(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

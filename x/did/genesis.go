package did

import (
	"github.com/SaoNetwork/sao/x/did/keeper"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the accountList
	for _, elem := range genState.AccountListList {
		k.SetAccountList(ctx, elem)
	}
	// Set all the accountAuth
	for _, elem := range genState.AccountAuthList {
		k.SetAccountAuth(ctx, elem)
	}
	// Set all the sidDocument
	for _, elem := range genState.SidDocumentList {
		k.SetSidDocument(ctx, elem)
	}
	// Set all the sidDocumentVersion
	for _, elem := range genState.SidDocumentVersionList {
		k.SetSidDocumentVersion(ctx, elem)
	}
	// Set all the pastSeeds
	for _, elem := range genState.PastSeedsList {
		k.SetPastSeeds(ctx, elem)
	}
	// Set all the paymentAddress
	for _, elem := range genState.PaymentAddressList {
		k.SetPaymentAddress(ctx, elem)
	}
	// Set all the accountId
	for _, elem := range genState.AccountIdList {
		k.SetAccountId(ctx, elem)
	}
	// Set all the did
	for _, elem := range genState.DidList {
		k.SetDid(ctx, elem)
	}
	// Set all the kid
	for _, elem := range genState.KidList {
		k.SetKid(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AccountListList = k.GetAllAccountList(ctx)
	genesis.AccountAuthList = k.GetAllAccountAuth(ctx)
	genesis.SidDocumentList = k.GetAllSidDocument(ctx)
	genesis.SidDocumentVersionList = k.GetAllSidDocumentVersion(ctx)
	genesis.PastSeedsList = k.GetAllPastSeeds(ctx)
	genesis.PaymentAddressList = k.GetAllPaymentAddress(ctx)
	genesis.AccountIdList = k.GetAllAccountId(ctx)
	genesis.DidList = k.GetAllDid(ctx)
	genesis.KidList = k.GetAllKid(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

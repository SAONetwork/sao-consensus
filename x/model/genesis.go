package model

import (
	"github.com/SaoNetwork/sao/x/model/keeper"
	"github.com/SaoNetwork/sao/x/model/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the metadata
	for _, elem := range genState.MetadataList {
		k.SetMetadata(ctx, elem)
	}
	// Set all the model
	for _, elem := range genState.ModelList {
		k.SetModel(ctx, elem)
	}
	// Set all the expiredData
	for _, elem := range genState.ExpiredDataList {
		k.SetExpiredData(ctx, elem)
	}
	// Set all the orderFinish
	for _, elem := range genState.OrderFinishList {
		k.SetOrderFinish(ctx, elem)
	}
	// Set all the expiredOrder
	for _, elem := range genState.ExpiredOrderList {
		k.SetExpiredOrder(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MetadataList = k.GetAllMetadata(ctx)
	genesis.ModelList = k.GetAllModel(ctx)
	genesis.ExpiredDataList = k.GetAllExpiredData(ctx)
	genesis.OrderFinishList = k.GetAllOrderFinish(ctx)
	genesis.ExpiredOrderList = k.GetAllExpiredOrder(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

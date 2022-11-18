package did

import (
	"math/rand"

	"github.com/SaoNetwork/sao/testutil/sample"
	didsimulation "github.com/SaoNetwork/sao/x/did/simulation"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = didsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddBinding = "op_weight_msg_add_binding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddBinding int = 100

	opWeightMsgGetBinding = "op_weight_msg_get_binding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGetBinding int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	didGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&didGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAddBinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddBinding, &weightMsgAddBinding, nil,
		func(_ *rand.Rand) {
			weightMsgAddBinding = defaultWeightMsgAddBinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddBinding,
		didsimulation.SimulateMsgAddBinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgGetBinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGetBinding, &weightMsgGetBinding, nil,
		func(_ *rand.Rand) {
			weightMsgGetBinding = defaultWeightMsgGetBinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGetBinding,
		didsimulation.SimulateMsgGetBinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
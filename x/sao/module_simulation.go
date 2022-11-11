package sao

import (
	"math/rand"

	"github.com/SaoNetwork/sao/testutil/sample"
	saosimulation "github.com/SaoNetwork/sao/x/sao/simulation"
	"github.com/SaoNetwork/sao/x/sao/types"
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
	_ = saosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgStore = "op_weight_msg_store"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStore int = 100

	opWeightMsgCancel = "op_weight_msg_cancel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancel int = 100

	opWeightMsgComplete = "op_weight_msg_complete"
	// TODO: Determine the simulation weight value
	defaultWeightMsgComplete int = 100

	opWeightMsgReject = "op_weight_msg_reject"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReject int = 100

	opWeightMsgTerminate = "op_weight_msg_terminate"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTerminate int = 100

	opWeightMsgReady = "op_weight_msg_ready"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReady int = 100

	opWeightMsgStore1 = "op_weight_msg_store_1"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStore1 int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	saoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&saoGenesis)
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

	var weightMsgStore int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStore, &weightMsgStore, nil,
		func(_ *rand.Rand) {
			weightMsgStore = defaultWeightMsgStore
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStore,
		saosimulation.SimulateMsgStore(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = defaultWeightMsgCancel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancel,
		saosimulation.SimulateMsgCancel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgComplete int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgComplete, &weightMsgComplete, nil,
		func(_ *rand.Rand) {
			weightMsgComplete = defaultWeightMsgComplete
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgComplete,
		saosimulation.SimulateMsgComplete(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReject, &weightMsgReject, nil,
		func(_ *rand.Rand) {
			weightMsgReject = defaultWeightMsgReject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReject,
		saosimulation.SimulateMsgReject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTerminate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTerminate, &weightMsgTerminate, nil,
		func(_ *rand.Rand) {
			weightMsgTerminate = defaultWeightMsgTerminate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTerminate,
		saosimulation.SimulateMsgTerminate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReady int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReady, &weightMsgReady, nil,
		func(_ *rand.Rand) {
			weightMsgReady = defaultWeightMsgReady
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReady,
		saosimulation.SimulateMsgReady(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

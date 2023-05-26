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
	opWeightMsgUpdatePaymentAddress = "op_weight_msg_update_payment_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePaymentAddress int = 100

	opWeightMsgBinding = "op_weight_msg_binding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBinding int = 100

	opWeightMsgUpdate = "op_weight_msg_update"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdate int = 100

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
		AccountIdList: []types.AccountId{
			{
				AccountDid: "0",
			},
			{
				AccountDid: "1",
			},
		},
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

	var weightMsgUpdatePaymentAddress int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePaymentAddress, &weightMsgUpdatePaymentAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePaymentAddress = defaultWeightMsgUpdatePaymentAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePaymentAddress,
		didsimulation.SimulateMsgUpdatePaymentAddress(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBinding, &weightMsgBinding, nil,
		func(_ *rand.Rand) {
			weightMsgBinding = defaultWeightMsgBinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBinding,
		didsimulation.SimulateMsgBinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdate, &weightMsgUpdate, nil,
		func(_ *rand.Rand) {
			weightMsgUpdate = defaultWeightMsgUpdate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdate,
		didsimulation.SimulateMsgUpdate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

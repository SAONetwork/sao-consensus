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

	opWeightMsgUnbinding = "op_weight_msg_unbinding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnbinding int = 100

	opWeightMsgAddAccountAuth = "op_weight_msg_add_account_auth"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAccountAuth int = 100

	opWeightMsgUpdateAccountAuths = "op_weight_msg_update_account_auths"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAccountAuths int = 100

	opWeightMsgUpdateSidDocument = "op_weight_msg_update_sid_document"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSidDocument int = 100

	opWeightMsgAddPastSeed = "op_weight_msg_add_past_seed"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddPastSeed int = 100

	opWeightMsgCleanupSidDocuments = "op_weight_msg_cleanup_sid_documents"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCleanupSidDocuments int = 100

	opWeightMsgCleanupPastSeeds = "op_weight_msg_cleanup_past_seeds"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCleanupPastSeeds int = 100

	opWeightMsgResetStore = "op_weight_msg_reset_store"
	// TODO: Determine the simulation weight value
	defaultWeightMsgResetStore int = 100

	opWeightMsgUpdatePaymentAddress = "op_weight_msg_update_payment_address"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePaymentAddress int = 100

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

	var weightMsgUnbinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnbinding, &weightMsgUnbinding, nil,
		func(_ *rand.Rand) {
			weightMsgUnbinding = defaultWeightMsgUnbinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnbinding,
		didsimulation.SimulateMsgUnbinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddAccountAuth int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddAccountAuth, &weightMsgAddAccountAuth, nil,
		func(_ *rand.Rand) {
			weightMsgAddAccountAuth = defaultWeightMsgAddAccountAuth
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAccountAuth,
		didsimulation.SimulateMsgAddAccountAuth(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAccountAuths int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAccountAuths, &weightMsgUpdateAccountAuths, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAccountAuths = defaultWeightMsgUpdateAccountAuths
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAccountAuths,
		didsimulation.SimulateMsgUpdateAccountAuths(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateSidDocument int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateSidDocument, &weightMsgUpdateSidDocument, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSidDocument = defaultWeightMsgUpdateSidDocument
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSidDocument,
		didsimulation.SimulateMsgUpdateSidDocument(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddPastSeed int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddPastSeed, &weightMsgAddPastSeed, nil,
		func(_ *rand.Rand) {
			weightMsgAddPastSeed = defaultWeightMsgAddPastSeed
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddPastSeed,
		didsimulation.SimulateMsgAddPastSeed(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCleanupSidDocuments int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCleanupSidDocuments, &weightMsgCleanupSidDocuments, nil,
		func(_ *rand.Rand) {
			weightMsgCleanupSidDocuments = defaultWeightMsgCleanupSidDocuments
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCleanupSidDocuments,
		didsimulation.SimulateMsgCleanupSidDocuments(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCleanupPastSeeds int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCleanupPastSeeds, &weightMsgCleanupPastSeeds, nil,
		func(_ *rand.Rand) {
			weightMsgCleanupPastSeeds = defaultWeightMsgCleanupPastSeeds
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCleanupPastSeeds,
		didsimulation.SimulateMsgCleanupPastSeeds(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgResetStore int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgResetStore, &weightMsgResetStore, nil,
		func(_ *rand.Rand) {
			weightMsgResetStore = defaultWeightMsgResetStore
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgResetStore,
		didsimulation.SimulateMsgResetStore(am.accountKeeper, am.bankKeeper, am.keeper),
	))

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

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

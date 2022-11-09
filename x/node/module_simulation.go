package node

import (
	"math/rand"

	"github.com/SaoNetwork/sao/testutil/sample"
	nodesimulation "github.com/SaoNetwork/sao/x/node/simulation"
	"github.com/SaoNetwork/sao/x/node/types"
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
	_ = nodesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgLogin = "op_weight_msg_login"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogin int = 100

	opWeightMsgLogout = "op_weight_msg_logout"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLogout int = 100

	opWeightMsgReset = "op_weight_msg_reset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgReset int = 100

	opWeightMsgClaimReward = "op_weight_msg_claim_reward"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimReward int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nodeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nodeGenesis)
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

	var weightMsgLogin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogin, &weightMsgLogin, nil,
		func(_ *rand.Rand) {
			weightMsgLogin = defaultWeightMsgLogin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogin,
		nodesimulation.SimulateMsgLogin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLogout int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLogout, &weightMsgLogout, nil,
		func(_ *rand.Rand) {
			weightMsgLogout = defaultWeightMsgLogout
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLogout,
		nodesimulation.SimulateMsgLogout(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgReset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgReset, &weightMsgReset, nil,
		func(_ *rand.Rand) {
			weightMsgReset = defaultWeightMsgReset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgReset,
		nodesimulation.SimulateMsgReset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimReward int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimReward, &weightMsgClaimReward, nil,
		func(_ *rand.Rand) {
			weightMsgClaimReward = defaultWeightMsgClaimReward
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimReward,
		nodesimulation.SimulateMsgClaimReward(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

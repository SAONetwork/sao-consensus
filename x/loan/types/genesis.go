package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1
const DefaultDenom string = "sao"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LoanPool: &LoanPool{
			Total:              sdk.NewInt64DecCoin(DefaultDenom, 0),
			LoanedOut:          sdk.NewInt64Coin(DefaultDenom, 0),
			AccInterestPerCoin: sdk.NewInt64DecCoin(DefaultDenom, 0),
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in credit

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

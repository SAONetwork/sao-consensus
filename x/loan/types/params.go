package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyLoanInterest = []byte("LoanInterest")
	// TODO: Determine the default value
	DefaultLoanInterest string = "loan_interest"
)

var (
	KeyMinLiquidityRatio = []byte("MinLiquidityRatio")
	// TODO: Determine the default value
	DefaultMinLiquidityRatio string = "min_liquidity_ratio"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	loanInterest string,
	minLiquidityRatio string,
) Params {
	return Params{
		LoanInterest:      loanInterest,
		MinLiquidityRatio: minLiquidityRatio,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultLoanInterest,
		DefaultMinLiquidityRatio,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLoanInterest, &p.LoanInterest, validateLoanInterest),
		paramtypes.NewParamSetPair(KeyMinLiquidityRatio, &p.MinLiquidityRatio, validateMinLiquidityRatio),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLoanInterest(p.LoanInterest); err != nil {
		return err
	}

	if err := validateMinLiquidityRatio(p.MinLiquidityRatio); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateLoanInterest validates the LoanInterest param
func validateLoanInterest(v interface{}) error {
	loanInterest, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = loanInterest

	return nil
}

// validateMinLiquidityRatio validates the MinLiquidityRatio param
func validateMinLiquidityRatio(v interface{}) error {
	minLiquidityRatio, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minLiquidityRatio

	return nil
}

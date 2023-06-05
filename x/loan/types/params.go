package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyLoanInterest            = []byte("LoanInterest")
	DefaultLoanInterest string = "0.000001"
)

var (
	KeyMinLiquidityRatio            = []byte("MinLiquidityRatio")
	DefaultMinLiquidityRatio string = "0.3"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	interestRatePerBlock string,
	minLiquidityRatio string,
) Params {
	return Params{
		InterestRatePerBlock: interestRatePerBlock,
		MinLiquidityRatio:    minLiquidityRatio,
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
		paramtypes.NewParamSetPair(KeyLoanInterest, &p.InterestRatePerBlock, validateLoanInterest),
		paramtypes.NewParamSetPair(KeyMinLiquidityRatio, &p.MinLiquidityRatio, validateMinLiquidityRatio),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLoanInterest(p.InterestRatePerBlock); err != nil {
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

	dec, err := sdk.NewDecFromStr(loanInterest)
	if err != nil {
		return err
	}
	if dec.IsNegative() {
		return fmt.Errorf("invalid loan interest %v", loanInterest)
	}
	return nil
}

// validateMinLiquidityRatio validates the MinLiquidityRatio param
func validateMinLiquidityRatio(v interface{}) error {
	minLiquidityRatio, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	dec, err := sdk.NewDecFromStr(minLiquidityRatio)
	if err != nil {
		return err
	}
	if dec.IsNegative() || dec.GTE(sdk.NewDec(1)) {
		return fmt.Errorf("invalid min liquidity ratio %v", minLiquidityRatio)
	}
	return nil
}

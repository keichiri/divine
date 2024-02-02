package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyDefaultFeeAmount = []byte("DefaultFeeAmount")
	// TODO: Determine the default value
	DefaultDefaultFeeAmount int32 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	defaultFeeAmount int32,
) Params {
	return Params{
		DefaultFeeAmount: defaultFeeAmount,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultDefaultFeeAmount,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDefaultFeeAmount, &p.DefaultFeeAmount, validateDefaultFeeAmount),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateDefaultFeeAmount(p.DefaultFeeAmount); err != nil {
		return err
	}

	return nil
}

// validateDefaultFeeAmount validates the DefaultFeeAmount param
func validateDefaultFeeAmount(v interface{}) error {
	defaultFeeAmount, ok := v.(int32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = defaultFeeAmount

	return nil
}

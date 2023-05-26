package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
)

var _ types.QueryServer = Keeper{}

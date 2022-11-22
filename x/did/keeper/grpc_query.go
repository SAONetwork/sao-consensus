package keeper

import (
	"github.com/SaoNetwork/sao/x/did/types"
)

var _ types.QueryServer = Keeper{}

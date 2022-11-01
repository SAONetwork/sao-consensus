package keeper

import (
	"github.com/SaoNetwork/sao/x/model/types"
)

var _ types.QueryServer = Keeper{}

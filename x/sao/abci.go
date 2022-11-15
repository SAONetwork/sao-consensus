package sao

import (
	"github.com/SaoNetwork/sao/x/sao/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// clean expired order
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
}

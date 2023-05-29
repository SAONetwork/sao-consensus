package keeper

import (
	"github.com/SaoNetwork/sao/x/loan/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetLoanPool set loanPool in the store
func (k Keeper) SetLoanPool(ctx sdk.Context, loanPool types.LoanPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanPoolKey))
	b := k.cdc.MustMarshal(&loanPool)
	store.Set([]byte{0}, b)
}

// GetLoanPool returns loanPool
func (k Keeper) GetLoanPool(ctx sdk.Context) (val types.LoanPool, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanPoolKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLoanPool removes loanPool from the store
func (k Keeper) RemoveLoanPool(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LoanPoolKey))
	store.Delete([]byte{0})
}

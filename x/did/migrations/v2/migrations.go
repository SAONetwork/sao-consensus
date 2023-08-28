package v2

import (
	saodidparser "github.com/SaoNetwork/sao-did/parser"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	logger := ctx.Logger()
	logger.Debug("migrating market")

	store := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.PaymentAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	kidStore := prefix.NewStore(ctx.KVStore(storeKey), types.KeyPrefix(types.KidKeyPrefix))

	for ; iterator.Valid(); iterator.Next() {

		var paymentAddress types.PaymentAddress
		err := cdc.Unmarshal(iterator.Value(), &paymentAddress)
		if err != nil {
			return err
		}

		did, err := saodidparser.Parse(paymentAddress.Did)
		if err != nil {
			logger.Error("failed to parse did", "did", paymentAddress.Did)
			return err
		}

		if did.Method == "key" {
			kid := types.Kid{
				Address: paymentAddress.Address,
				Kid:     paymentAddress.Did,
			}

			b := cdc.MustMarshal(&kid)
			kidStore.Set(types.KidKey(
				kid.Address,
			), b)
		}

	}

	return nil
}

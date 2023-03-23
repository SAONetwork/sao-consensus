package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpdatePaymentAddress{}, "did/UpdatePaymentAddress", nil)
	cdc.RegisterConcrete(&MsgBinding{}, "did/Binding", nil)
	cdc.RegisterConcrete(&MsgUpdate{}, "did/Update", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePaymentAddress{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBinding{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdate{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

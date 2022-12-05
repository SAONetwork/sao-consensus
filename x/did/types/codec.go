package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddBinding{}, "did/AddBinding", nil)
	cdc.RegisterConcrete(&MsgUnbinding{}, "did/Unbinding", nil)
	cdc.RegisterConcrete(&MsgAddAccountAuth{}, "did/AddAccountAuth", nil)
	cdc.RegisterConcrete(&MsgUpdateAccountAuths{}, "did/UpdateAccountAuths", nil)
	cdc.RegisterConcrete(&MsgUpdateSidDocument{}, "did/UpdateSidDocument", nil)
	cdc.RegisterConcrete(&MsgAddPastSeed{}, "did/AddPastSeed", nil)
	cdc.RegisterConcrete(&MsgResetStore{}, "did/ResetStore", nil)
	cdc.RegisterConcrete(&MsgUpdatePaymentAddress{}, "did/UpdatePaymentAddress", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddBinding{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnbinding{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAccountAuth{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateAccountAuths{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateSidDocument{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddPastSeed{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgResetStore{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePaymentAddress{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

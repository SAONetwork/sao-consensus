package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgStore{}, "sao/Store", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "sao/Cancel", nil)
	cdc.RegisterConcrete(&MsgComplete{}, "sao/Complete", nil)
	cdc.RegisterConcrete(&MsgReject{}, "sao/Reject", nil)
	cdc.RegisterConcrete(&MsgTerminate{}, "sao/Terminate", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStore{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgComplete{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTerminate{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

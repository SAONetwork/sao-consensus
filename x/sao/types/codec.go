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
	cdc.RegisterConcrete(&MsgTerminate{}, "sao/Terminate", nil)
	cdc.RegisterConcrete(&MsgReady{}, "sao/Ready", nil)
	cdc.RegisterConcrete(&MsgRenew{}, "sao/Renew", nil)
	cdc.RegisterConcrete(&MsgUpdataPermission{}, "sao/UpdataPermission", nil)
	cdc.RegisterConcrete(&MsgMigrate{}, "sao/Migrate", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "sao/Deposit", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "sao/Withdraw", nil)
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
		&MsgTerminate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgReady{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRenew{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdataPermission{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMigrate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeposit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdraw{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

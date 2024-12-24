package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPayForBlobs{},
		&evmtypes.MsgEthereumTx{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil))

	registry.RegisterInterface(
		"ethermint.evm.v1.MsgEthereumTx",
		(*evmtypes.TxData)(nil),
	)

	registry.RegisterImplementations(
		(*evmtypes.TxData)(nil),
		&evmtypes.LegacyTx{},
		&evmtypes.AccessListTx{},
		&evmtypes.DynamicFeeTx{},
	)

	registry.RegisterInterface(
		"cosmos.auth.v1beta1.BaseAccount",
		(*authtypes.AccountI)(nil),
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&authtypes.BaseAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

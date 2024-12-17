// APACHE NOTICE
// Sourced with modifications from https://github.com/strangelove-ventures/lens
package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clientTypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	coretypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	celTypes "github.com/nodersteam/probe/client/codec/celestiaorg/celestia-app/x/blob/types"
	tendermintLiquidityTypes "github.com/nodersteam/probe/client/codec/tendermint/liquidity/x/liquidity/types"
)

type Codec struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeCodec(moduleBasics []module.AppModuleBasic) Codec {
	modBasic := module.NewBasicManager(moduleBasics...)
	encodingConfig := MakeCodecConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	modBasic.RegisterLegacyAminoCodec(encodingConfig.Amino)
	modBasic.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	coretypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	transfertypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	clientTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}

// Split out from base codec to not include explicitly.
// Should be included only when needed.
func RegisterTendermintLiquidityInterfaces(aminoCodec *codec.LegacyAmino, registry types.InterfaceRegistry) {
	tendermintLiquidityTypes.RegisterLegacyAminoCodec(aminoCodec)
	tendermintLiquidityTypes.RegisterInterfaces(registry)
}

func MakeCodecConfig() Codec {
	interfaceRegistry := types.NewInterfaceRegistry()
	interfaceRegistry.RegisterInterface("cosmos.crypto.Pubkey", (*cryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
	//interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	coretypes.RegisterInterfaces(interfaceRegistry)
	transfertypes.RegisterInterfaces(interfaceRegistry)
	clientTypes.RegisterInterfaces(interfaceRegistry)
	celTypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return Codec{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             codec.NewLegacyAmino(),
	}
}

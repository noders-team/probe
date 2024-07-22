package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/keys/ed25519"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/keys/ethsecp256k1"
	"github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/keys/secp256k1"
	cryptotypes "github.com/oracleNetworkProtocol/plugchain-sdk-go/crypto/types"
)

// RegisterLegacyAminoCodec registers concrete types on the codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLockTokens{},
		&MsgBeginUnlockingAll{},
		&MsgBeginUnlocking{},
		&MsgExtendLockup{},
		&MsgForceUnlock{},
		//Old Osmosis Lockup types
		&MsgUnlockPeriodLock{},
		&MsgUnlockTokens{},
	)
	registry.RegisterInterface("cosmos.crypto.Pubkey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// legacy amino codecs
var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/liquidity module codec. Note, the
	// codec should ONLY be used in certain instances of tests and for JSON
	// encoding as Amino is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/liquidity
	// and defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

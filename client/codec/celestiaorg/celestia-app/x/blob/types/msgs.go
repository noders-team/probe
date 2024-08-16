package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgPayForBlobs)(nil)
)

// Message types for the liquidity module
const (
	TypeMsgPayForBlobs = "pay_for_blobs"
)

func (msg MsgPayForBlobs) Route() string { return "" }

func (msg MsgPayForBlobs) Type() string { return TypeMsgPayForBlobs }

func (msg MsgPayForBlobs) ValidateBasic() error {
	return nil
}

func (msg MsgPayForBlobs) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgPayForBlobs) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (msg MsgPayForBlobs) GetPoolCreator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return addr
}

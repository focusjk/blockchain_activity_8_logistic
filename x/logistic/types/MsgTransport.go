package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgTransport{}

// MsgTransport - struct for unjailing jailed validator
type MsgTransport struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	Transporter sdk.AccAddress `json:"transporter" yaml:"transporter"`
}

// NewMsgTransport creates a new MsgTransport instance
func NewMsgTransport(creator sdk.AccAddress, transporter sdk.AccAddress) MsgTransport {
	return MsgTransport{
		Creator:     creator,
		Transporter: transporter,
	}
}

const TransportConst = "Transport"

// nolint
func (msg MsgTransport) Route() string { return RouterKey }
func (msg MsgTransport) Type() string  { return TransportConst }
func (msg MsgTransport) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgTransport) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgTransport) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing creator address")
	}
	if msg.Transporter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing transporter address")
	}
	return nil
}

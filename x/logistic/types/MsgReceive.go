package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgReceive{}

// MsgReceive - struct for unjailing jailed validator
type MsgReceive struct {
	Customer sdk.AccAddress `json:"customer" yaml:"customer"` // address of the validator operator
}

// NewMsgReceive creates a new MsgReceive instance
func NewMsgReceive(customer sdk.AccAddress) MsgReceive {
	return MsgReceive{
		Customer: customer,
	}
}

const ReceiveConst = "Receive"

// nolint
func (msg MsgReceive) Route() string { return RouterKey }
func (msg MsgReceive) Type() string  { return ReceiveConst }
func (msg MsgReceive) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Customer)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgReceive) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgReceive) ValidateBasic() error {
	if msg.Customer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing customer address")
	}
	return nil
}

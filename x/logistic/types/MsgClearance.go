package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgClearance{}

// MsgClearance - struct for unjailing jailed validator
type MsgClearance struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgClearance creates a new MsgClearance instance
func NewMsgClearance(creator sdk.AccAddress) MsgClearance {
	return MsgClearance{
		Creator: creator,
	}
}

const ClearanceConst = "Clearance"

// nolint
func (msg MsgClearance) Route() string { return RouterKey }
func (msg MsgClearance) Type() string  { return ClearanceConst }
func (msg MsgClearance) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgClearance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgClearance) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing creator address")
	}
	return nil
}

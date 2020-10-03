package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgUpdateTmp{}

// MsgUpdateTmp - struct for unjailing jailed validator
type MsgUpdateTmp struct {
	Transporter sdk.AccAddress `json:"transporter" yaml:"transporter"`
	Tmp         int            `json:"tmp" yaml:"tmp"`
}

// NewMsgUpdateTmp creates a new MsgUpdateTmp instance
func NewMsgUpdateTmp(transporter sdk.AccAddress, tmp int) MsgUpdateTmp {
	return MsgUpdateTmp{
		Transporter: transporter,
		Tmp:         tmp,
	}
}

const UpdateTmpConst = "UpdateTmp"

// nolint
func (msg MsgUpdateTmp) Route() string { return RouterKey }
func (msg MsgUpdateTmp) Type() string  { return UpdateTmpConst }
func (msg MsgUpdateTmp) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Transporter)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUpdateTmp) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUpdateTmp) ValidateBasic() error {
	if msg.Transporter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing transporter address")
	}
	return nil
}

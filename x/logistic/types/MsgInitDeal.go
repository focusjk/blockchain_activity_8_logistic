package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgInitDeal{}

// MsgInitDeal - struct for unjailing jailed validator
type MsgInitDeal struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	Customer sdk.AccAddress `json:"customer" yaml:"customer"`
	Price    sdk.Coins      `json:"price" yaml:"price"`
	MaxTemp  int            `json:"maxTmp" yaml:"maxTmp"`
	MinTemp  int            `json:"minTmp" yaml:"minTmp"`
}

// NewMsgInitDeal creates a new MsgInitDeal instance
func NewMsgInitDeal(creator sdk.AccAddress, customer sdk.AccAddress, price sdk.Coins, maxTemp int, minTemp int) MsgInitDeal {
	return MsgInitDeal{
		Creator:  creator,
		Customer: customer,
		Price:    price,
		MaxTemp:  maxTemp,
		MinTemp:  minTemp,
	}
}

const InitDealConst = "InitDeal"

// nolint
func (msg MsgInitDeal) Route() string { return RouterKey }
func (msg MsgInitDeal) Type() string  { return InitDealConst }
func (msg MsgInitDeal) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgInitDeal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgInitDeal) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing creator address")
	}
	if msg.Customer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing customer address")
	}
	if msg.MaxTemp <= msg.MinTemp {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "maxTmp/minTmp invalid")
	}
	return nil
}

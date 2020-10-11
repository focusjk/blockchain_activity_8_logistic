// Step 1
// Note: message for intiDeal, creator must provide customer, price, orderid, maxTemp and minTemp

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
	OrderID  string         `json:"orderid" yaml:"orderid"`
	MaxTemp  int            `json:"maxTemp" yaml:"maxTemp"`
	MinTemp  int            `json:"minTemp" yaml:"minTemp"`
}

// NewMsgInitDeal creates a new MsgInitDeal instance
func NewMsgInitDeal(creator sdk.AccAddress, orderId string, price sdk.Coins, customer sdk.AccAddress, maxTemp int, minTemp int) MsgInitDeal {
	return MsgInitDeal{
		Creator:  creator,
		Customer: customer,
		Price:    price,
		OrderID:  orderId,
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
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "maxTemp/minTemp invalid")
	}
	return nil
}

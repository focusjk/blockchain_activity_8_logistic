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
	Price    uint           `json:"price" yaml:"price"`
	MaxTemp  int            `json:"maxTmp" yaml:"maxTmp"`
	MixTemp  int            `json:"minTmp" yaml:"minTmp"`
}

// NewMsgInitDeal creates a new MsgInitDeal instance
func NewMsgInitDeal(creator sdk.AccAddress, customer sdk.AccAddress, price uint, maxTmp int, mixTmp int) MsgInitDeal {
	return MsgInitDeal{
		Creator:  creator,
		Customer: customer,
		Price:    price,
		MaxTemp:  maxTmp,
		MixTemp:  mixTmp,
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
	if msg.Price == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid price")
	}
	if msg.MaxTemp <= msg.MixTemp {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "maxTmp/mixTmp invalid")
	}
	return nil
}

package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StateType int

const (
	Created = iota
	InTransit
	Complete
	Cancel
)

type Deal struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	Transporter sdk.AccAddress `json:"transporter" yaml:"transporter"`
	Customer    sdk.AccAddress `json:"customer" yaml:"customer"`
	Price       uint           `json:"price" yaml:"price"`
	MaxTemp     int            `json:"maxTmp" yaml:"maxTmp"`
	MixTemp     int            `json:"minTmp" yaml:"minTmp"`
	Cancelable  bool           `json:"cancelable" yaml:"cancelable"`
	State       StateType      `json:"state" yaml:"state"`
}

func (d Deal) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Transporter: %s
	Customer: %s
	Price: %d
	MaxTemp: %d
	MixTemp: %d,
	Cancelable: %t,
	State: %s`,
		d.Creator,
		d.Transporter,
		d.Customer,
		d.Price,
		d.MaxTemp,
		d.MixTemp,
		d.Cancelable,
		string(d.State),
	))
}

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// TODO: Register the modules msgs
	cdc.RegisterConcrete(MsgInitDeal{}, "logistic/InitDeal", nil)
	cdc.RegisterConcrete(MsgTransport{}, "logistic/Transport", nil)
	cdc.RegisterConcrete(MsgUpdateTmp{}, "logistic/UpdateTmp", nil)
	cdc.RegisterConcrete(MsgReceive{}, "logistic/Receive", nil)
	cdc.RegisterConcrete(MsgReject{}, "logistic/Reject", nil)
	cdc.RegisterConcrete(MsgClearance{}, "logistic/Clearance", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

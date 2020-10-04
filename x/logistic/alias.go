package logistic

import (
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	NewMsgInitDeal  = types.NewMsgInitDeal
	NewMsgTransport = types.NewMsgTransport
	NewMsgUpdateTmp = types.NewMsgUpdateTmp
	NewMsgReceive   = types.NewMsgReceive
	NewMsgReject    = types.NewMsgReject
	NewMsgClearance = types.NewMsgClearance
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgInitDeal  = types.MsgInitDeal
	MsgTransport = types.MsgTransport
	MsgUpdateTmp = types.MsgUpdateTmp
	MsgReceive   = types.MsgReceive
	MsgReject    = types.MsgReject
	MsgClearance = types.MsgClearance
)

package types

// logistic module event types
const (
	EventTypeInitDeal  = "InitDeal"
	EventTypeTransport = "Transport"
	EventTypeUpdateTmp = "UpdateTmp"
	EventTypeReceive   = "Receive"
	EventTypeReject    = "Reject"

	AttributeCreator     = "creator"
	AttributeTransporter = "transporter"
	AttributeCustomer    = "customer"
	AttributePrice       = "price"
	AttributeMaxTmp      = "maxTmp"
	AttributeMinTmp      = "minTmp"
	AttributeCancelable  = "cancelable"
	AttributeUpdateTmp   = "updateTmp"

	AttributeValueCategory = ModuleName
)

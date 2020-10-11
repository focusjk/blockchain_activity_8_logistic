// Step 6: create handler for MsgReject
// Note: this file is use to initialize deal when receive MsgReject

package logistic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

func handlerMsgReject(ctx sdk.Context, k keeper.Keeper, msg types.MsgReject) (*sdk.Result, error) {
	// get deal with orderid
	currentDeal, err := k.GetDeal(ctx, msg.OrderID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	// check if caller is customer
	if msg.Customer.String() != currentDeal.Customer.String() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Fail reject, invalid customer")
	}
	// check if state is valid
	if currentDeal.State != types.InTransit {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}
	// check if deal is cancelable
	if !currentDeal.Cancelable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot reject deal since product is good")
	}

	// set state to complete and update deal
	currentDeal.State = types.Cancel
	k.SetDeal(ctx, currentDeal)

	// set event (for logging transaction)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeReject),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Customer.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

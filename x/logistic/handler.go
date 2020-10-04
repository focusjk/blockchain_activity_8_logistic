package logistic

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/types"
)

// NewHandler creates an sdk.Handler for all the logistic type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgInitDeal:
			return handlerMsgInitDeal(ctx, k, msg)
		case MsgTransport:
			return handlerMsgTransport(ctx, k, msg)
		case MsgUpdateTmp:
			return handlerMsgUpdateTmp(ctx, k, msg)
		case MsgReceive:
			return handlerMsgReceive(ctx, k, msg)
		case MsgReject:
			return handlerMsgReject(ctx, k, msg)
		// case MsgClearance:
		// 	return handlerMsgClearance(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handlerMsgInitDeal(ctx sdk.Context, k Keeper, msg MsgInitDeal) (*sdk.Result, error) {
	var deal = types.Deal{
		Creator:  msg.Creator,
		Customer: msg.Customer,
		Price:    msg.Price,
		MaxTemp:  msg.MaxTemp,
		MinTemp:  msg.MinTemp,
		State:    types.Created,
	}
	_, err := k.GetDeal(ctx, deal.Creator)
	if err == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal of this creator already exists")
	}
	// moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// sdkError := k.CoinKeeper.SendCoins(ctx, scavenge.Creator, moduleAcct, scavenge.Reward)
	// if sdkError != nil {
	// 	return nil, sdkError
	// }

	k.SetDeal(ctx, deal)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeInitDeal),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeCreator, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeCustomer, msg.Customer.String()),
			sdk.NewAttribute(types.AttributePrice, msg.Price.String()),
			sdk.NewAttribute(types.AttributeMaxTmp, string(msg.MaxTemp)),
			sdk.NewAttribute(types.AttributeMinTmp, string(msg.MinTemp)),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handlerMsgTransport(ctx sdk.Context, k Keeper, msg MsgTransport) (*sdk.Result, error) {
	var deal = types.Deal{
		Creator:     msg.Creator,
		Transporter: msg.Transporter,
		State:       types.InTransit,
	}
	currentDeal, err := k.GetDeal(ctx, deal.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	if currentDeal.State != types.Created {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}
	k.SetDeal(ctx, deal)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeTransport),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeCreator, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeTransporter, msg.Transporter.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handlerMsgUpdateTmp(ctx sdk.Context, k Keeper, msg MsgUpdateTmp) (*sdk.Result, error) {
	var deal = types.Deal{
		Transporter: msg.Transporter,
	}
	currentDeal, err := k.GetDeal(ctx, deal.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	if currentDeal.State != types.InTransit {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}
	if msg.Tmp > currentDeal.MaxTemp || msg.Tmp < currentDeal.MinTemp {
		deal.Cancelable = true
	}

	k.SetDeal(ctx, deal)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeUpdateTmp),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Transporter.String()),
			sdk.NewAttribute(types.AttributeTransporter, msg.Transporter.String()),
			sdk.NewAttribute(types.AttributeUpdateTmp, string(msg.Tmp)),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handlerMsgReceive(ctx sdk.Context, k Keeper, msg MsgReceive) (*sdk.Result, error) {
	var deal = types.Deal{
		Customer: msg.Customer,
		State:    types.Complete,
	}
	currentDeal, err := k.GetDeal(ctx, deal.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	if currentDeal.State != types.InTransit {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}

	sdkError := k.CoinKeeper.SendCoins(ctx, currentDeal.Customer, currentDeal.Creator, currentDeal.Price)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetDeal(ctx, deal)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeReceive),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Customer.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handlerMsgReject(ctx sdk.Context, k Keeper, msg MsgReject) (*sdk.Result, error) {
	var deal = types.Deal{
		Customer: msg.Customer,
		State:    types.Cancel,
	}
	currentDeal, err := k.GetDeal(ctx, deal.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}
	if !currentDeal.Cancelable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Cannot reject deal since product is good")
	}

	sdkError := k.CoinKeeper.SendCoins(ctx, currentDeal.Customer, currentDeal.Creator, currentDeal.Price)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetDeal(ctx, deal)
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

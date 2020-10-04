package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/types"
)

// NewQuerier creates a new querier for logistic clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryDeals:
			return listDeals(ctx, k)
		case types.QueryDeal:
			return getDeal(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown logistic query endpoint")
		}
	}
}

// RemovePrefixFromHash removes the prefix from the key
func RemovePrefixFromHash(key []byte, prefix []byte) (hash []byte) {
	hash = key[len(prefix):]
	return hash
}

func listDeals(ctx sdk.Context, k Keeper) ([]byte, error) {
	var dealList types.QueryResDeals

	iterator := k.GetDealsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		deal := RemovePrefixFromHash(iterator.Key(), []byte(types.DealPrefix))
		dealList = append(dealList, string(deal))
	}

	res, err := codec.MarshalJSONIndent(k.cdc, dealList)
	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getDeal(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	creator := path[0]
	Deal, err := k.GetDeal(ctx, creator)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, Deal)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

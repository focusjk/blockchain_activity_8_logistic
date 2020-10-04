package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/earth2378/logistic/x/logistic/types"
)

// Keeper of the logistic store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a logistic keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetDeal(ctx sdk.Context, deal types.Deal) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(deal)
	key := []byte(types.DealPrefix + string(deal.Creator))
	store.Set(key, bz)
}

func (k Keeper) GetDeal(ctx sdk.Context, creator string) (types.Deal, error) {
	store := ctx.KVStore(k.storeKey)
	var deal types.Deal
	byteKey := []byte(types.DealPrefix + creator)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &deal)
	if err != nil {
		return deal, err
	}
	return deal, nil
}

func (k Keeper) GetDealsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.DealPrefix))
}

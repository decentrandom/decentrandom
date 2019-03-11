package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine.
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the rand Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Setter method for the round
func (k Keeper) SetRound(ctx sdk.Context, id string, round Round) {
	if round.Id.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(round))
}

// Getter method for the round
func (k Keeper) GetRound(vtx sdk.Context, id string) Round {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(id)) {
		return NewRound()
	}
	bz := store.Get([]byte(id))
	var round Round
	k.cdc.MustUnmarshalBinaryBare(bz, &round)
	return round
}

func (k Keeper) SetOwner(ctx sdk.Context, id string, owner sdk.AccAddress) {
	store := ctx.KVStore(k.ownersStoreKey)
	store.Set([]byte(id), owner)
}

func (k Keeper) HasOwner(ctx sdk.Context, id string) bool {
	return !k.GetRound(ctx, id).Owner.Empty()
}

func (k Keeper) GetOwner(ctx sdk.Context, id string) sdk.AccAddress {
	return k.GetRound(ctx, id).Owner
}

func (k Keeper) GetIdsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byteP{})
}

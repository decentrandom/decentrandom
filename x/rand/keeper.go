package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper -
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper -
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// SetRound is setter method for the round
func (k Keeper) SetRound(ctx sdk.Context, id string, round Round) {
	if len(id) == 0 || round.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(round))
}

// GetRound is getter method for the round
func (k Keeper) GetRound(ctx sdk.Context, id string) Round {
	store := ctx.KVStore(k.storeKey)
	var round Round

	if !store.Has([]byte(id)) {
		return round
	}
	bz := store.Get([]byte(id))

	k.cdc.MustUnmarshalBinaryBare(bz, &round)
	return round
}

// GetOwner -
func (k Keeper) GetOwner(ctx sdk.Context, id string) sdk.AccAddress {
	return k.GetRound(ctx, id).Owner
}

// SetTargets -
func (k Keeper) SetTargets(ctx sdk.Context, id string, targets []string) {
	round := k.GetRound(ctx, id)
	round.Targets = targets
	k.SetRound(ctx, id, round)
}

// GetIDsIterator -
func (k Keeper) GetIDsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper - struct 선언
type Keeper struct {
	coinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper - cerates new instances of the rand Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// SetRound - round setter
func (k Keeper) SetRound(ctx sdk.Context, id string, round Round) {
	if len(id) == 0 || round.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(round))
}

// GetRound - round getter
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

// GetOwner - gets the owner
func (k Keeper) GetOwner(ctx sdk.Context, id string) sdk.AccAddress {
	return k.GetRound(ctx, id).Owner
}

// GetTargets - gets targets
func (k Keeper) GetTargets(ctx sdk.Context, id string) []string {
	return k.GetRound(ctx, id).Targets
}

// SetTargets - sets targets
func (k Keeper) SetTargets(ctx sdk.Context, id string, targets []string) {
	round := k.GetRound(ctx, id)
	round.Targets = targets
	k.SetRound(ctx, id, round)
}

// SetNonce - sets the nonce
func (k Keeper) SetNonce(ctx sdk.Context, id string, nonce string) {
	round := k.GetRound(ctx, id)
	round.Nonce = nonce
	k.SetRound(ctx, id, round)
}

// SetSeeds - sets seeds
func (k Keeper) SetSeeds(ctx sdk.Context, height int64, seeds []string, sealedSeeds []string) {
	// to-do
}

// GetIDsIterator - get an iterator over all rounds
func (k Keeper) GetIDsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// GetHeightsIterator - get an iterator over all heights
func (k Keeper) GetHeightsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

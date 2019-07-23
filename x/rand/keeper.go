package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper -
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec
}

// NewKeeper -
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// SetRound -
func (k Keeper) SetRound(ctx sdk.Context, id string, round Round) {
	if len(id) == 0 || round.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(round))
}

// GetRound -
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

// GetTargets -
func (k Keeper) GetTargets(ctx sdk.Context, id string) []string {
	return k.GetRound(ctx, id).Targets
}

// SetTargets -
func (k Keeper) SetTargets(ctx sdk.Context, id string, targets []string) {
	round := k.GetRound(ctx, id)
	round.Targets = targets
	k.SetRound(ctx, id, round)
}

// SetNonce -
func (k Keeper) SetNonce(ctx sdk.Context, id string, nonce string) {
	round := k.GetRound(ctx, id)
	round.Nonce = nonce
	k.SetRound(ctx, id, round)
}

// GetIDsIterator -
func (k Keeper) GetIDsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

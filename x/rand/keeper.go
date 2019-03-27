package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper - struct 선언
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey

	cdc *codec.Codec
}

// NewKeeper - keeper 생성
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// SetRound - 라운드 setter
func (k Keeper) SetRound(ctx sdk.Context, id string, round Round) {
	if len(id) == 0 || round.Owner.Empty() {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id), k.cdc.MustMarshalBinaryBare(round))
}

// GetRound - 라운드 getter
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

// GetOwner - 라운드 소유자 getter
func (k Keeper) GetOwner(ctx sdk.Context, id string) sdk.AccAddress {
	return k.GetRound(ctx, id).Owner
}

// GetTargets - 라운드 모집단 getter
func (k Keeper) GetTargets(ctx sdk.Context, id string) []string {
	return k.GetRound(ctx, id).Targets
}

// SetTargets - 라운드 모집단 setter
func (k Keeper) SetTargets(ctx sdk.Context, id string, targets []string) {
	round := k.GetRound(ctx, id)
	round.Targets = targets
	k.SetRound(ctx, id, round)
}

// SetNonce - 라운드 Nonce setter
func (k Keeper) SetNonce(ctx sdk.Context, id string, nonce int16) {
	round := k.GetRound(ctx, id)
	round.Nonce = nonce
	k.SetRound(ctx, id, round)
}

// GetIDsIterator - 전체 라운드 ID getter
func (k Keeper) GetIDsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

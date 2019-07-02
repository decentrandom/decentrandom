package slashing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)

// StoreKey -
const StoreKey = "slashing"

// Keeper -
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	sk         types.StakingKeeper
	paramspace params.Subspace

	// codespace
	codespace sdk.CodespaceType
}

package slashing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)

const StoreKey = "slashing"

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	sk         types.StakingKeeper
	paramspace params.Subspace

	// codespace
	codespace sdk.CodespaceType
}

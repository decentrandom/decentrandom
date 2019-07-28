package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	//abci "github.com/tendermint/tendermint/abci/types"
)

// GenesisState - song genesis state
type GenesisState struct {
	RoundRecords []Round `json:"round_records"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(roundRecords []Round) GenesisState {
	return GenesisState{RoundRecords: nil}
}

// ValidateGenesis validates genesis state
func ValidateGenesis(data GenesisState) error {
	fmt.Println("OMG")
	return nil
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		RoundRecords: []Round{},
	}
}

// Checks whether 2 GenesisState structs are equivalent.
/*func (data GenesisState) Equal(data2 GenesisState) bool {
	b1 := MsgCdc.MustMarshalBinaryBare(data)
	b2 := MsgCdc.MustMarshalBinaryBare(data2)
	return bytes.Equal(b1, b2)
}*/

// Returns if a GenesisState is empty or has data in it
/*func (data GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return data.Equal(emptyGenState)
}*/

// InitGenesis -
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Round
	iterator := k.GetIDsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var round Round
		round = k.GetRound(ctx, id)
		records = append(records, round)
	}
	return GenesisState{RoundRecords: records}
}

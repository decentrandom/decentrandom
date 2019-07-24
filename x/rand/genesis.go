package rand

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - song genesis state
type GenesisState struct {
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState() GenesisState {
	return GenesisState{}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{}
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

// ValidateGenesis validates genesis state
func ValidateGenesis(data GenesisState) error {
	return nil
}

// InitGenesis -
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {

}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return GenesisState{}
}
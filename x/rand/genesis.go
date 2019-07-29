package rand

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	//abci "github.com/tendermint/tendermint/abci/types"
)

// GenesisState -
type GenesisState struct {
	StartingBlockHeight uint64 `json:"starting_block_height"`
	Rounds              Rounds `json:"rounds"`
}

// NewGenesisState -
func NewGenesisState(startingBlockHeight uint64) GenesisState {
	return GenesisState{
		StartingBlockHeight: 1,
	}
}

// ValidateGenesis -
func ValidateGenesis(data GenesisState) error {
	return nil
}

// DefaultGenesisState -
func DefaultGenesisState() GenesisState {
	return GenesisState{
		StartingBlockHeight: 1,
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
	/*for _, record := range data.Rounds {
		keeper.SetRound(ctx, record.ID, record)
	}*/
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	/*var records []Round
	iterator := k.GetIDsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var round Round
		round = k.GetRound(ctx, id)
		records = append(records, round)
	}*/
	return GenesisState{StartingBlockHeight: 1}
}

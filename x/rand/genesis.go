package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// GenesisState -
type GenesisState struct {
	RoundRecords []Round `json:"round_records"`
}

// NewGenesisState -
func NewGenesisState(roundRecords []Round) GenesisState {
	return GenesisState{RoundRecords: nil}
}

// ValidateGenesis -
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.RoundRecords {
		if record.Owner == nil {
			return fmt.Errorf("Invalid RoundRecord: Value: %s. Error: Missing Owner", record.ID)
		}
	}
	return nil
}

// DefaultGenesisState -
func DefaultGenesisState() GenesisState {
	return GenesisState{
		RoundRecords: []Round{},
	}
}

// InitGenesis -
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.RoundRecords {
		keeper.SetRound(ctx, record.ID, record)
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis -
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

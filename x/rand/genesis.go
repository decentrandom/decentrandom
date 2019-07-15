package rand

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	RoundRecords []Round `json:"round_records"`
}

func NewGenesisState(roundRecords []Round) GenesisState {
	return GenesisState{RoundRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.RoundRecords {
		if record.Owner == nil {
			return fmt.Errorf("Invalid RoundRecord: Value: %s. Error: Missing Owner", record.Id)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		RoundRecords: []Round{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.RoundRecords {
		keeper.SetRound()(ctx, record.Id, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Round
	iterator := k.GetRoundsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var round Round
		round = k.GetRound(ctx, id)
		records = append(records, round)
	}
	return GenesisState{RoundRecords: records}
}

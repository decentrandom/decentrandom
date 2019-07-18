package app

import (
	"encoding/json"
)

// GenesisState -
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState -
func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

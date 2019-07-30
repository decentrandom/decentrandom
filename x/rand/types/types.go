package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/*
// Seed -
type Seed struct {
	Height           string   `json:"height"`
	SeedHashes       []string `json:"seed_hashes"`
	SealedSeedHashes []string `json:"sealed_seed_hashes"`
	ValidatorPubKey  string   `jsong:"validator_pub_key"`
}

func (s Seed) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Height: %s
	SeedHashes: %v
	SealedSeedHashes: %v
	ValidatorPubKey: %s
	`, s.Height, s.SeedHashes, s.SealedSeedHashes, s.ValidatorPubKey))
}
*/

// Round -
type Round struct {
	ID            string         `json:"id"`
	Difficulty    uint8          `json:"difficulty"`
	Owner         sdk.AccAddress `json:"owner"`
	Nonce         string         `json:"nonce"`
	NonceHash     string         `json:"nonce_hash"`
	Targets       []string       `json:"targets"`
	ScheduledTime time.Time      `json:"scheduled_time"`
}

func (r Round) String() string {
	timeString := r.ScheduledTime.Local()
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Difficulty: %d
Nonce: %s
NonceHash: %s
Targets: %v
ScheduledTime: %s
`, r.Owner, r.Difficulty, r.Nonce, r.NonceHash, r.Targets, timeString.Format("2006-01-02 15:04:05 +0900")))
}

// QueryResRoundIDs -
type QueryResRoundIDs []string

func (n QueryResRoundIDs) String() string {
	return strings.Join(n[:], "\n")
}

// Rounds -
type Rounds []*Round

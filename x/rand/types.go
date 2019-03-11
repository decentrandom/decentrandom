package rand

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//var DefaultRoundDifficulty = 1

type Round struct {
	Difficulty    int16          `json:"difficulty"`
	Owner         sdk.AccAddress `json:"owner"`
	Nonce         int16          `json:"nonce"`
	NonceHash     string         `json:"nonce_hash"`
	Targets       []string       `json:"targets"`
	ScheduledTime time.Time      `json:"scheduled_time"`
	SeedHeights   []int64        `json:"seed_heights"`
}

func (r Round) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Difficulty: %d
Nonce: %d
NonceHash: %s`, r.Owner, r.Difficulty, r.Nonce, r.NonceHash))
}

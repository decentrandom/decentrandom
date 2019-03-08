package rand

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//var DefaultRoundDifficulty = 1

type Round struct {
	Difficulty  int16          `json:"difficulty"`
	Owner       sdk.AccAddress `json:"owner"`
	Nonce       int16          `json:"nonce"`
	NonceHash   string         `json:"nonce_hash"`
	Targets     []string       `json:"targets"`
	RunTime     time.Time      `json:"time"`
	SeedHeights []int64        `json:"seed_heights"`
}

func NewRound() Round {
	return Round{}
}

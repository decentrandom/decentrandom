package rand

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName        = "rand"
	RouterKey         = ModuleName
	StoreKey          = ModuleName
	QuerierRoute      = ModuleName
	DefaultParamspace = ModuleName
)

// Round - 라운드 기본 구조체
type Round struct {
	Difficulty    int16          `json:"difficulty"`
	Owner         sdk.AccAddress `json:"owner"`
	Nonce         int16          `json:"nonce"`
	NonceHash     string         `json:"nonce_hash"`
	Targets       []string       `json:"targets"`
	ScheduledTime time.Time      `json:"scheduled_time"`
}

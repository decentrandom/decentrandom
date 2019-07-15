package rand

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// golint
const (
	ModuleName        = "rand"
	RouterKey         = ModuleName
	StoreKey          = ModuleName
	QuerierRoute      = ModuleName
	DefaultParamspace = ModuleName
)

// Round - 라운드 기본 구조체
type Round struct {
	Id            string         `json:"id"`
	Difficulty    uint8          `json:"difficulty"`
	Owner         sdk.AccAddress `json:"owner"`
	Nonce         string         `json:"nonce"`
	NonceHash     string         `json:"nonce_hash"`
	Targets       []string       `json:"targets"`
	ScheduledTime time.Time      `json:"scheduled_time"`
}

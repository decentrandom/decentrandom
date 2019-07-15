package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

// implement fmt.Stringer
func (r Round) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Nonce: %s
NonceHash: %s`, r.Owner, r.Nonce, r.NonceHash))
}

package rand

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	cmn "github.com/tendermint/tendermint/libs/common"
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
	Difficulty    uint8          `json:"difficulty"`
	Owner         sdk.AccAddress `json:"owner"`
	Nonce         cmn.HexBytes   `json:"nonce"`
	NonceHash     cmn.HexBytes   `json:"nonce_hash"`
	Targets       []string       `json:"targets"`
	ScheduledTime time.Time      `json:"scheduled_time"`
}

package app

import (
	"encoding/json"
	"errors"
	"fmt"
	//"github.com/decentrandom/decentrandom/types/assets"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// GenesisState - 제네시스 상태 구조체
type GenesisState struct {
	AuthData     auth.GenesisState      `json:"auth"`
	BankData     bank.GenesisState      `json:"bank"`
	Accounts     []*auth.BaseAccount    `json:"accounts"`
	StakingData  staking.GenesisState   `json:"staking"`
	SlashingData slashsing.GenesisState `json:"slashing"`
}

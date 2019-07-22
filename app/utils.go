package app

import (
	"io"

	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// DONTCOVER

// NewRandAppUNSAFE -
func NewRandAppUNSAFE(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) (rapp *RandApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	rapp = NewRandApp(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	return rapp, rapp.keyMain, rapp.keyStaking, rapp.stakingKeeper
}
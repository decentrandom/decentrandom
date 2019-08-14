package app

import (
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// DONTCOVER

// NewRandAppUNSAFE -
func NewRandAppUNSAFE(logger log.Logger, db dbm.DB, invCheckPeriod uint) (rapp *RandApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	rapp = NewRandApp(logger, db, invCheckPeriod)
	return rapp, rapp.keyMain, rapp.keyStaking, rapp.stakingKeeper
}

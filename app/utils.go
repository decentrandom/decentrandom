package app

import (
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// NewRandAppUNSAFE -
func NewRandAppUNSAFE(logger log.Logger, db dbm.DB, invCheckPeriod uint) (rapp *RandApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	rapp = NewRandApp(logger, db, invCheckPeriod)
	return rapp, rapp.keys[bam.MainStoreKey], rapp.keys[staking.StoreKey], rapp.stakingKeeper
}

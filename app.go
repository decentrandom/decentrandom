package app

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/ibc"

	"github.com/decentrandom/decentrandom/x/rand"
)

const (
	appName = "decentrandom"
)

type decentRandomApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	accountKeeper auth.AccountKeeper
	bankKeeper    bank.Keeper
}

func NewDecentRandomApp(logger log.Logger, db dbm.DB) *decentRandomApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

}

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	return cdc
}

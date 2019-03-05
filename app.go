package app

import (
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/decentrandom/decentrandom/x/rand"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"
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

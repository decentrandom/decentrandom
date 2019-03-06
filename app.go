package app

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/examples/democoin/x/simplestake"

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
)

const (
	appName = "decentrandom"
)

// Extended ABCI application
type decentRandomApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// keys to access the substores
	capKeyMainStore    *sdk.KVStoreKey
	capKeyAccountStore *sdk.KVStoreKey
	capKeyPowStore     *sdk.KVStoreKey
	capKeyIBCStore     *sdk.KVStoreKey
	capKeyStakingStore *sdk.KVStoreKey

	// keepers
	feeCollectionKeeper auth.FeeCollectionKeeper
	coinKeeper          bank.Keeper
	coolKeeper          cool.Keeper
	powKeeper           pow.Keeper
	ibcMapper           ibc.Mapper
	stakeKeeper         simplestake.Keeper

	// Manage getting and setting accounts
	accountMapper auth.AccountMapper
}

func NewDecentRandomApp(logger log.Logger, db dbm.DB) *decentRandomApp {

	// Create app-level codec for txs and acctouns.
	var cdc = MakeCodec()

	// Create application object.
	var app = &decentRandomApp{
		BaseApp:            bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc)),
		cdc:                cdc,
		capKeyMainStore:    sdk.NewKVStoreKey("main"),
		capKeyAccountStore: sdk.NewKVStoreKey("acc"),
		capKeyPowStore:     sdk.NewKVStoreKey("pow"),
		capKeyIBCStore:     sdk.NewKVStoreKey("ibc"),
		capKeyStakingStore: sdk.NewKVStoreKey("stake"),
	}

	// Define the accountMapper.
	app.accountMapper = auth.NewAccountMapper(
		cdc,
		app.capKeyAccountStore,
		types.ProtoAppAccount,
	)

	// Add handlers.
	app.coinKeeper = bank.NewKeeper(app.accountMapper)
	app.coolKeeper = cool.NewKeeper(app.capKeyMainStore, app.coinKeeper, app.RegisterCodespace(cool.DefaultCodespace))
	app.powKeeper = pow.NewKeeper(app.capKeyPowStore, pow.NewConfig("pow", int64(1)), app.coinKeeper, app.RegisterCodespace(pow.DefaultCodespace))
	app.ibcMapper = ibc.NewMapper(app.cdc, app.capKeyIBCStore, app.RegisterCodespace(ibc.DefaultCodespace))
	app.stakeKeeper = simplestake.NewKeeper(app.capKeyStakingStore, app.coinKeeper, app.RegisterCodespace(simplestake.DefaultCodespace))
	app.Router().
		AddRoute("bank", bank.NewHandler(app.coinKeeper)).
		AddRoute("cool", cool.NewHandler(app.coolKeeper)).
		AddRoute("pow", app.powKeeper.Handler).
		AddRoute("sketchy", sketchy.NewHandler()).
		AddRoute("ibc", ibc.NewHandler(app.ibcMapper, app.coinKeeper)).
		AddRoute("simplestake", simplestake.NewHandler(app.stakeKeeper))

	// Initialize BaseApp.
	app.SetInitChainer(app.initChainerFn(app.coolKeeper, app.powKeeper))
	app.MountStoresIAVL(app.capKeyMainStore, app.capKeyAccountStore, app.capKeyPowStore, app.capKeyIBCStore, app.capKeyStakingStore)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountMapper, app.feeCollectionKeeper))

	err := app.LoadLatestVersion(app.capKeyMainStore)
	if err != nil {
		cmn.Exit(err.Error())
	}

	app.Seal()

	return app
}

func MakeCodec() *codec.Codec {
	var cdc = codec.NewCodec()
	return cdc
}

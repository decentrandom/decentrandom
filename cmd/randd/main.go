package main

import (
	"encoding/json"
	//"fmt"
	"io"
	//"io/ioutil"
	//"os"
	//"path/filepath"
	//"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	//"github.com/cosmos/cosmos-sdk/x/auth"
	//"github.com/cosmos/cosmos-sdk/x/bank"
	randInit "github.com/decentrandom/decentrandom/cmd/init"
	"github.com/decentrandom/decentrandom/types/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrandom/decentrandom/app"
	abci "github.com/tendermint/tendermint/abci/types"
	//cfg "github.com/tendermint/tendermint/config"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"

	randServer "github.com/decentrandom/decentrandom/server"
)

const flagAssertInvariantsBlockly = "assert-invariants-blockly"

var assertInvariantsBlockly bool

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(util.Bech32PrefixAccAddr, util.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(util.Bech32PrefixValAddr, util.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(util.Bech32PrefixConsAddr, util.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "randd",
		Short:             "DecentRandom Daemon (server)",
		PersistentPreRunE: randServer.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(randInit.InitCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.CollectGenTxsCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.TestnetFilesCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.GenTxCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.ValidateGenesisCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.AddGenesisAccountCmd(ctx, cdc))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))

	server.AddCommands(ctx, cdc, rootCmd, newApp, appExporter())

	executor := cli.PrepareBaseCmd(rootCmd, "DR", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewRandApp(
		logger, db, traceStore, true, assertInvariantsBlockly,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		rApp := app.NewRandApp(logger, db, traceStore, false, false)
		err := rApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	rApp := app.NewRandApp(logger, db, traceStore, true, false)
	return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/decentrandom/decentrandom/app"
	randInit "github.com/decentrandom/decentrandom/cmd/init"
	randServer "github.com/decentrandom/decentrandom/server"
	"github.com/decentrandom/decentrandom/types/util"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

const flagAssertInvariantsBlockly = "assert-invariants-blockly"

var assertInvariantsBlockly bool

// main -
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

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "DR", app.DefaultNodeHome)
	rootCmd.PersistentFlags().BoolVar(&assertInvariantsBlockly, flagAssertInvariantsBlockly,
		false, "Assert registered invariants on a blockly basis")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

// newApp -
func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewRandApp(
		logger, db, traceStore, true, assertInvariantsBlockly,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

// exportAppStateAndTMValidators -
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

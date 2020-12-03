package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tm-db"
	//"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/decentrandom/decentrandom/app"
	//randInit "github.com/decentrandom/decentrandom/cmd/init"
	"github.com/decentrandom/decentrandom/types/util"

	//"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	//"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// randd custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(util.Bech32PrefixAccAddr, util.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(util.Bech32PrefixValAddr, util.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(util.Bech32PrefixConsAddr, util.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "randd",
		Short:             "Rand Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "DR", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod, 0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewRandApp(
		logger, db, invCheckPeriod,
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		rApp := app.NewRandApp(logger, db, uint(1))
		err := rApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	rApp := app.NewRandApp(logger, db, uint(1))
	return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

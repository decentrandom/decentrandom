package main

import (
	"encoding/json"
	//"fmt"
	"io"
	// "io/ioutil"
	//"os"
	//"path/filepath"
	// "strings"

	//"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	//"github.com/cosmos/cosmos-sdk/x/auth"
	//"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/decentrandom/decentrandom/types/util"
	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/cli"
	//"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"

	//gaiaInit "github.com/cosmos/cosmos-sdk/cmd/gaia/init"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/decentrandom/decentrandom"
	randInit "github.com/decentrandom/decentrandom/cmd/init"
	abci "github.com/tendermint/tendermint/abci/types"
	//cfg "github.com/tendermint/tendermint/config"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	flagOverwrite = "overwrite"
)

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
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(randInit.InitCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.AddGenesisAccountCmd(ctx, cdc))
	rootCmd.AddCommand(randInit.GenTxCmd(ctx, cdc))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "DR", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewRandApp(logger, db)
}

// exportAppStateAndTMValidators -
func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		rApp := app.NewRandApp(logger, db)
		err := rApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	rApp := app.NewRandApp(logger, db)
	return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

// SimpleAppGenTx -
func SimpleAppGenTx(cdc *codec.Codec, pk crypto.PubKey) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {

	addr, secret, err := server.GenerateCoinKey()
	if err != nil {
		return
	}

	bz, err := cdc.MarshalJSON(struct {
		Addr sdk.AccAddress `json:"addr"`
	}{addr})
	if err != nil {
		return
	}

	appGenTx = json.RawMessage(bz)

	bz, err = cdc.MarshalJSON(map[string]string{"secret": secret})
	if err != nil {
		return
	}

	cliPrint = json.RawMessage(bz)

	validator = tmtypes.GenesisValidator{
		PubKey: pk,
		Power:  10,
	}

	return
}

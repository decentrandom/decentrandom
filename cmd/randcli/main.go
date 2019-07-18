package main

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/rakyll/statik/fs"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/decentrandom/decentrandom/app"
	_ "github.com/decentrandom/decentrandom/client/lcd/statik"
	"github.com/decentrandom/decentrandom/types/util"
	"github.com/decentrandom/decentrandom/version"
	"github.com/decentrandom/decentrandom/x/rand"
	randClient "github.com/decentrandom/decentrandom/x/rand/client"
	randrest "github.com/decentrandom/decentrandom/x/rand/client/rest"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	at "github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	crisisclient "github.com/cosmos/cosmos-sdk/x/crisis/client"
	distcmd "github.com/cosmos/cosmos-sdk/x/distribution"
	distClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	dist "github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	sl "github.com/cosmos/cosmos-sdk/x/slashing"
	slashingClient "github.com/cosmos/cosmos-sdk/x/slashing/client"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	st "github.com/cosmos/cosmos-sdk/x/staking"
	stakingClient "github.com/cosmos/cosmos-sdk/x/staking/client"
	staking "github.com/cosmos/cosmos-sdk/x/staking/client/rest"

	amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"
)

// main -
func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(util.Bech32PrefixAccAddr, util.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(util.Bech32PrefixValAddr, util.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(util.Bech32PrefixConsAddr, util.Bech32PrefixConsPub)
	config.Seal()

	mc := []sdk.ModuleClients{
		distClient.NewModuleClient(distcmd.StoreKey, cdc),
		stakingClient.NewModuleClient(st.StoreKey, cdc),
		slashingClient.NewModuleClient(sl.StoreKey, cdc),
		randClient.NewModuleClient(rand.StoreKey, cdc),
		crisisclient.NewModuleClient(sl.StoreKey, cdc),
	}

	rootCmd := &cobra.Command{
		Use:   "randcli",
		Short: "Command line interface for interacting with randd",
	}

	rootCmd.PersistentFlags().String(client.FlagChainID, "", "tendermint node's chain ID")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(app.DefaultCLIHome),
		queryCmd(cdc, mc),
		txCmd(cdc, mc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
		client.NewCompletionCmd(rootCmd, true),
	)

	executor := cli.PrepareMainCmd(rootCmd, "DR", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		fmt.Printf("Failed executing CLI command : %s, closing...\n", err)
		os.Exit(1)
	}
}

// queryCmd -
func queryCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
		authcmd.GetAccountCmd(at.StoreKey, cdc),
	)

	/*
	for _, m := range mc {
		mQueryCmd := m.GetQueryCmd()
		if mQueryCmd != nil {
			queryCmd.AddCommand(mQueryCmd)
		}
	}
	*/

	// add modules' query commands
	app.ModuleBasics.AddQueryCommands(queryCmd, cdc)

	return queryCmd
}

// txCmd -
func txCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		client.LineBreak,
		authcmd.GetSignCommand(cdc),
		authcmd.GetMultiSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		tx.GetEncodeCommand(cdc),
		client.LineBreak,
	)

	/*
	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}
	*/

	// add modules' tx commands
	app.ModuleBasics.AddTxCommands(txCmd, cdc)

	return txCmd
}

// registerRoutes -
func registerRoutes(rs *lcd.RestServer) {

	client.RegisterRoutes(rs.CliCtx, rs.Mux)
	authrest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
	app.ModuleBasics.RegisterRESTRoutes(rs.CliCtx, rs.Mux)
}

// initConfig -
func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}

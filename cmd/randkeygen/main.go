package main

import (
	"fmt"
	//"net/http"
	"os"
	//"path"

	//"github.com/rakyll/statik/fs"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"

	"github.com/decentrandom/decentrandom/app"
	_ "github.com/decentrandom/decentrandom/client/lcd/statik"
	"github.com/decentrandom/decentrandom/types/util"

	//"github.com/decentrandom/decentrandom/version"
	//"github.com/decentrandom/decentrandom/x/rand"
	//randClient "github.com/decentrandom/decentrandom/x/rand/client"
	//randrest "github.com/decentrandom/decentrandom/x/rand/client/rest"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"

	//"github.com/cosmos/cosmos-sdk/client/lcd"
	//"github.com/cosmos/cosmos-sdk/client/rpc"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tendermint/libs/cli"
)

// main -
func main() {
	cobra.EnableCommandSorting = false

	//cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(util.Bech32PrefixAccAddr, util.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(util.Bech32PrefixValAddr, util.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(util.Bech32PrefixConsAddr, util.Bech32PrefixConsPub)
	config.Seal()

	/*
		mc := []sdk.ModuleClients{
			distClient.NewModuleClient(distcmd.StoreKey, cdc),
			stakingClient.NewModuleClient(st.StoreKey, cdc),
			slashingClient.NewModuleClient(sl.StoreKey, cdc),
			randClient.NewModuleClient(rand.StoreKey, cdc),
			crisisclient.NewModuleClient(sl.StoreKey, cdc),
		}
	*/

	rootCmd := &cobra.Command{
		Use:   "randkeygen",
		Short: "DecentRandom 주소 생성기",
	}

	/*
		rootCmd.PersistentFlags().String(client.FlagChainID, "", "텐더민트 노드의 체인 아이디")
		rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
			return initConfig(rootCmd)
		}
	*/

	rootCmd.AddCommand(
		keys.Commands(),
		client.NewCompletionCmd(rootCmd, true),
	)

	executor := cli.PrepareMainCmd(rootCmd, "DR", app.DefaultKeygenHome)
	err := executor.Execute()
	if err != nil {
		fmt.Printf("명령어 실행 실패 : %s, 종료합니다...\n", err)
		os.Exit(1)
	}
}

/*
// queryCmd -
func queryCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "",
		Aliases: []string{"q"},
		Short:   "쿼리 하부 명령어",
	}

	queryCmd.AddCommand(

		authcmd.GetAccountCmd(at.StoreKey, cdc),
	)

	for _, m := range mc {
		mQueryCmd := m.GetQueryCmd()
		if mQueryCmd != nil {
			queryCmd.AddCommand(mQueryCmd)
		}
	}

	return queryCmd
}

// txCmd -
func txCmd(cdc *amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "트랜잭션 하부 명령어",
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

	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}

	return txCmd
}

// CLIVersionRequestHandler -
func CLIVersionRequestHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf("{\"version\": \"%s\"}", version.Version)))
}

// registerRoutes -
func registerRoutes(rs *lcd.RestServer) {

	rs.Mux.HandleFunc("/version", CLIVersionRequestHandler).Methods("GET")
	registerSwaggerUI(rs)

	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, at.StoreKey)
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
*/

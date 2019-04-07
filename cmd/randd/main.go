package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/decentrandom/decentrandom/types/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"

	gaiaInit "github.com/cosmos/cosmos-sdk/cmd/gaia/init"
	sdk "github.com/cosmos/cosmos-sdk/types"
	app "github.com/decentrandom/decentrandom"
	abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"
)

// DefaultNodeHome - 기본 노드 홈
var DefaultNodeHome = os.ExpandEnv("$HOME/.randd")

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

	rootCmd.AddCommand(InitCmd(ctx, cdc))
	rootCmd.AddCommand(AddGenesisAccountCmd(ctx, cdc))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "DR", DefaultNodeHome)
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
	rApp := app.NewRandApp(logger, db, traceStore, true)
	return rApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}

// InitCmd - init 명령어
func InitCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "제네시스 설정, priv-validator 파일, P2P 노드 파일을 초기화 합니다.",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			chainID := viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}

			_, pk, err := gaiaInit.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			var appState json.RawMessage
			genFile := config.GenesisFile()

			if !viper.GetBool(flagOverwrite) && common.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}

			genesis := app.GenesisState{
				AuthData: auth.DefaultGenesisState(),
				BankData: bank.DefaultGenesisState(),
			}

			appState, err = codec.MarshalJSONIndent(cdc, genesis)
			if err != nil {
				return err
			}

			_, _, validator, err := SimpleAppGenTx(cdc, pk)
			if err != nil {
				return err
			}

			if err = gaiaInit.ExportGenesisFile(genFile, chainID, []tmtypes.GenesisValidator{validator}, appState); err != nil {
				return err
			}

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

			fmt.Printf("%s - 설정 및 부트스트랩 파일을 초기화하였습니다....\n", viper.GetString(cli.HomeFlag))
			return nil
		},
	}

	cmd.Flags().String(cli.HomeFlag, DefaultNodeHome, "노드의 홈 디렉토리")
	cmd.Flags().String(client.FlagChainID, "", "제네시스 파일 체인 ID 입니다. 공백으로 남겨두면 임의로 생성합니다.")
	cmd.Flags().BoolP(flagOverwrite, "o", false, "genesis.json 파일을 덮어씁니다.")

	return cmd
}

// AddGenesisAccountCmd - add-genesis-account 명령어
func AddGenesisAccountCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account [address] [coins[,coins]]",
		Short: "제네시스 파일에 계정을 추가합니다.",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(`
		제네시스 파일을 추가하여 CLI를 통해 코인과 함께 체인을 구동할 수 있도록 합니다.:
		$ randd add-genesis-account cosmos1tse7r2fadvlrrgau3pa0ss7cqh55wrv6y9alwh 1000STAKE,1000mycoin
		`),
		RunE: func(_ *cobra.Command, args []string) error {
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			coins.Sort()

			var genDoc tmtypes.GenesisDoc
			config := ctx.Config
			genFile := config.GenesisFile()
			if !common.FileExists(genFile) {
				return fmt.Errorf("%s 파일이 존재하지 않습니다. `gaiad init` 를 먼저 실행하시기 바랍니다", genFile)
			}
			genContents, err := ioutil.ReadFile(genFile)
			if err != nil {
			}

			if err = cdc.UnmarshalJSON(genContents, &genDoc); err != nil {
				return err
			}

			var appState app.GenesisState
			if err = cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
				return err
			}

			for _, stateAcc := range appState.Accounts {
				if stateAcc.Address.Equals(addr) {
					return fmt.Errorf("이미 %v 계정에 대한 정보를 가지고 있습니다", addr)
				}
			}

			acc := auth.NewBaseAccountWithAddress(addr)
			acc.Coins = coins
			appState.Accounts = append(appState.Accounts, &acc)
			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return err
			}

			return gaiaInit.ExportGenesisFile(genFile, genDoc.ChainID, genDoc.Validators, appStateJSON)
		},
	}

	return cmd
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

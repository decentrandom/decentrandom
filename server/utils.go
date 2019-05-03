package server

import (
	"encoding/json"
	"errors"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/version"

	tcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
	pvm "github.com/tendermint/tendermint/privval"
)

// PersistentPreRunEFn -
func PersistentPreRunEFn(context *server.Context) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == version.VersionCmd.Name() {
			return nil
		}
		config, err := interceptLoadConfig()
		if err != nil {
			return err
		}
		err = validateConfig(config)
		if err != nil {
			return err
		}
		logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
		logger, err = tmflags.ParseLogLevel(config.LogLevel, logger, cfg.DefaultLogLevel())
		if err != nil {
			return err
		}
		if viper.GetBool(cli.TraceFlag) {
			logger = log.NewTracingLogger(logger)
		}
		logger = logger.With("module", "main")
		context.Config = config
		context.Logger = logger
		return nil
	}
}

// interceptLoadConfig -
func interceptLoadConfig() (conf *cfg.Config, err error) {
	tmpConf := cfg.DefaultConfig()
	err = viper.Unmarshal(tmpConf)
	if err != nil {
		panic(err)
	}
	rootDir := tmpConf.RootDir
	configFilePath := filepath.Join(rootDir, "config/config.toml")

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		conf, _ = tcmd.ParseConfig()
		conf.ProfListenAddress = "localhost:6060"
		conf.P2P.RecvRate = 5120000
		conf.P2P.SendRate = 5120000
		conf.TxIndex.IndexAllTags = true
		conf.Consensus.TimeoutCommit = 5 * time.Second
		cfg.WriteConfigFile(configFilePath, conf)
	}

	if conf == nil {
		conf, err = tcmd.ParseConfig()
	}

	randConfigFilePath := filepath.Join(rootDir, "config/randd.toml")
	if _, err := os.Stat(randConfigFilePath); os.IsNotExist(err) {
		randConf, _ := config.ParseConfig()
		config.WriteConfigFile(randConfigFilePath, randConf)
	}

	viper.SetConfigName("randd")
	err = viper.MergeInConfig()

	return
}

// validateConfig -
func validateConfig(conf *cfg.Config) error {
	if !conf.Consensus.CreateEmptyBlocks {
		return errors.New("config option CreateEmptyBlocks = false is currently unsupported")
	}
	return nil
}

// AddCommands -
func AddCommands(
	ctx *server.Context, cdc *codec.Codec,
	rootCmd *cobra.Command,
	appCreator server.AppCreator, appExport server.AppExporter) {

	rootCmd.PersistentFlags().String("log_level", ctx.Config.LogLevel, "Log level")

	tendermintCmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint subcommands",
	}

	tendermintCmd.AddCommand(
		server.ShowNodeIDCmd(ctx),
		server.ShowValidatorCmd(ctx),
		server.ShowAddressCmd(ctx),
		server.VersionCmd(ctx),
	)

	rootCmd.AddCommand(
		server.StartCmd(ctx, appCreator),
		server.UnsafeResetAllCmd(ctx),
		client.LineBreak,
		tendermintCmd,
		server.ExportCmd(ctx, cdc, appExport),
		client.LineBreak,
		version.VersionCmd,
	)
}

// InsertKeyJSON -
func InsertKeyJSON(cdc *codec.Codec, baseJSON []byte, key string, value json.RawMessage) ([]byte, error) {
	var jsonMap map[string]json.RawMessage

	if err := cdc.UnmarshalJSON(baseJSON, &jsonMap); err != nil {
		return nil, err
	}

	jsonMap[key] = value
	bz, err := codec.MarshalJSONIndent(cdc, jsonMap)

	return json.RawMessage(bz), err
}

// ExternalIP -
func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if skipInterface(iface) {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ip := addrToIP(addr)
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

// TrapSignal -
func TrapSignal(cleanupFunc func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		switch sig {
		case syscall.SIGTERM:
			defer cleanupFunc()
			os.Exit(128 + int(syscall.SIGTERM))
		case syscall.SIGINT:
			defer cleanupFunc()
			os.Exit(128 + int(syscall.SIGINT))
		}
	}()
}

// UpgradeOldPrivValFile -
func UpgradeOldPrivValFile(config *cfg.Config) {
	if _, err := os.Stat(config.OldPrivValidatorFile()); !os.IsNotExist(err) {
		if oldFilePV, err := pvm.LoadOldFilePV(config.OldPrivValidatorFile()); err == nil {
			oldFilePV.Upgrade(config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())
		}
	}
}

// skipInterface -
func skipInterface(iface net.Interface) bool {
	if iface.Flags&net.FlagUp == 0 {
		return true
	}
	if iface.Flags&net.FlagLoopback != 0 {
		return true
	}
	return false
}

// addrToIP -
func addrToIP(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	return ip
}

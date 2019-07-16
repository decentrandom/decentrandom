package client

import (
	"github.com/spf13/cobra"

	randcmd "github.com/decentrandom/decentrandom/x/rand/client/cli"

	"github.com/cosmos/cosmos-sdk/client"

	amino "github.com/tendermint/go-amino"
)

// ModuleClient -
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

// NewModuleClient -
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd -
func (mc ModuleClient) GetQueryCmd() *cobra.Command {

	randQueryCmd := &cobra.Command{
		Use:   "rand",
		Short: "Querying commands for the rand module",
	}

	randQueryCmd.AddCommand(client.GetCommands(
		randcmd.GetCmdRoundInfo(mc.storeKey, mc.cdc),
	)...)

	return randQueryCmd
}

// GetTxCmd -
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	randTxCmd := &cobra.Command{
		Use:   "rand",
		Short: "Transactions subcommands",
	}

	randTxCmd.AddCommand(client.PostCommands(
		randcmd.GetCmdNewRound(mc.cdc),
		randcmd.GetCmdDeployNonce(mc.cdc),
		randcmd.GetCmdAddTargets(mc.cdc),
		randcmd.GetCmdRemoveTargets(mc.cdc),
	)...)

	return randTxCmd
}

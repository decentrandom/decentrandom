package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	randcmd "github.com/decentrandom/decentrandom/x/rand/client/cli"
	"github.com/gogo/protobuf/codec"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	codec    *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) getQueryCmd() *cobra.Command {
	// Group rand queries under a subcommand
	randQueryCmd := &cobra.Command{
		Use:   "rand",
		Short: "Querying commands for the rand module",
	}

	randQueryCmd.AddCommand(client.GetCommands(
	/*
		to-do
	*/
	)...)

	return randQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	randTxCmd := &cobra.Command{
		Use:   "rand",
		Short: "Rand transactions subcommands",
	}

	randTxCmd.AddCommand(client.PostCommands()...)

	return randTxCmd
}

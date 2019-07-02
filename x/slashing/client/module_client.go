package client

import (
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/decentrandom/decentrandom/x/slashing/client/cli"
)

// ModuleClient -
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

// NewModuleClient
func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd -
func (mc ModuleClient) GetQueryCmd() *cobra.Command {

	slashingQueryCmd := &cobra.Command{
		Use:   slashing.ModuleName,
		Short: "Querying commands for the slashing module",
	}

	slashingQueryCmd.AddCommand(
		client.GetCommands(
			cli.GetCmdQuerySigningInfo(mc.storeKey, mc.cdc),
			cli.GetCmdQueryParams(mc.cdc),
		)...,
	)

	return slashingQueryCmd
}

// GetTxCmd -
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	slashingTxCmd := &cobra.Command{
		Use:  slashing.ModuleName,
		Shor: "Slashing tx subcommands",
	}

	slashingTxCmd.AddCommand(client.PostCommands(
		cli.GetCmdUnjail(mc.cdc),
	)...)

	return slashingTxCmd
}

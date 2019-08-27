package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/decentrandom/decentrandom/x/rand/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetQueryCmd -
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	randQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the rand module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	randQueryCmd.AddCommand(client.GetCommands(
		GetCmdRoundInfo(cdc),
		GetCmdRoundIDs(cdc),
	)...)
	return randQueryCmd
}

// GetCmdRoundInfo -
func GetCmdRoundInfo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round-info [id]",
		Short: "get information of certain round",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", types.QuerierRoute, id), nil)
			if err != nil {
				//fmt.Printf("Cannot receive round %s data\nError : %s \nqueryRoute : %s\n", string(id), err.Error(), queryRoute)
				return nil
			}

			var out types.Round
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdRoundIDs -
func GetCmdRoundIDs(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round-ids",
		Short: "Get round IDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round_ids", types.QuerierRoute), nil)
			if err != nil {
				fmt.Printf("Cannot receive IDs .\n")
				return nil
			}

			var out types.QueryResRoundIDs
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)

		},
	}
}

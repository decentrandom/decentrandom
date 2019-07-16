package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/decentrandom/decentrandom/x/rand"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetCmdRoundInfo -
func GetCmdRoundInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round_info [id]",
		Short: "get information of certain round",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("Cannot receive round %s data %s \n", string(id), string(err))
				return nil
			}

			var out rand.Round
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdRoundIDs -
func GetCmdRoundIDs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round_ids",
		Short: "Get round IDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round_ids", queryRoute), nil)
			if err != nil {
				fmt.Printf("Cannot receive IDs .\n")
				return nil
			}

			var out rand.QueryResRoundIDs
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)

		},
	}
}

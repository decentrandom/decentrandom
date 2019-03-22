package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/decentrandom/decentrandom/x/rand"
	"github.com/spf13/cobra"
)

// GetCmdRoundInfo -
func GetCmdRoundInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round_info [id]",
		Short: "Query round info of ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not get round - %s \n", string(id))
				return nil
			}

			var out rand.Round
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdRounds -
func GetCmdRounds(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:	"rounds",
		Short:	"rounds",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/rounds", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query rounds\n")
				return nil
			}

			var out rand.QueryResRounds
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)

		}
	}
}
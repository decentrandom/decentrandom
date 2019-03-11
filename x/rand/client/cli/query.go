package cli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decentrandom/decentrandom/x/rand"
	"github.com/spf13/cobra"
)

func GetCmdIds(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "ids",
		Short: "ids",
		//Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprinf("custom/$s/ids", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query ids\n")
				return nil
			}

			var out rand.QueryResIds
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

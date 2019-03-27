package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/decentrandom/decentrandom/x/rand"
	"github.com/spf13/cobra"
)

// GetCmdRoundInfo - 라운드 정보
func GetCmdRoundInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round_info [id]",
		Short: "ID에 해당하는 라운드 정보 요청",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("ID %s에 해당하는 라운드 정보를 받아오지 못했습니다. \n", string(id))
				return nil
			}

			var out rand.Round
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdRoundIDs - 라운드 ID 리스트
func GetCmdRoundIDs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "round_ids",
		Short: "라운드 ID 받아오기",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round_ids", queryRoute), nil)
			if err != nil {
				fmt.Printf("라운드 ID 내역을 받아오지 못했습니다.\n")
				return nil
			}

			var out rand.QueryResRoundIDs
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)

		},
	}
}

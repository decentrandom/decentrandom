package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/decentrandom/decentrandom/x/rand/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/tendermint/tendermint/crypto/tmhash"
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
	randQueryCmd.AddCommand(flags.GetCommands(
		GetCmdRoundInfo(cdc),
		GetCmdRoundIDs(cdc),
		GetCmdHashNonce(cdc),
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

// GetCmdHashNonce -
func GetCmdHashNonce(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "hash-nonce [nonce]",
		Short: "get hash string from a nonce",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			nonce := args[0]

			// Hashing Nonce
			hasher := tmhash.New()
			nonceVector := []byte(nonce)
			_, hashError := hasher.Write(nonceVector)
			if hashError != nil {
				return hashError
			}
			bz := tmhash.Sum(nonceVector)
			nonceHash := hex.EncodeToString(bz)

			nonceStruct := types.Nonce{Nonce: nonce, NonceHash: nonceHash}

			//cdc.MustUnmarshalJSON(res, &nonceStruct)
			return cliCtx.PrintOutput(nonceStruct)
		},
	}
}

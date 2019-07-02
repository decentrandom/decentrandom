package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
)

// GetCmdQuerySigningInfo -
func GetCmdQuerySigningInfo(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signing-info --validator-conspub [validator-conspub]",
		Short: "Query a validator's signing information",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			consPubKey := viper.GetString(FlagConsensusPubKeyValidator)
			pk, err := sdk.GetConsPubKeyBech32(consPubKey)
			if err != nil {
				return err
			}

			consAddr := sdk.ConsAddress(pk.Address())
			key := slashing.GetValidatorSigningInfoKey(consAddr)

			res, err := cliCtx.QueryStore(key, storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("Validator %s not found in slashing store", consAddr)
			}

			var signingInfo slashing.ValidatorSigningInfo
			cdc.MustUnmarshalBinaryLengthPrefixed(res, &signingInfo)

			return cliCtx.PrintOutput(signingInfo)
		},
	}

	cmd.Flags().String(FlagConsensusPubKeyValidator, "", "validator's consensus public key")

	cmd.MarkFlagRequired(FlagConsensusPubKeyValidator)

	return cmd
}

// GetCmdQueryParams -
func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "params",
		Shor: "Query the current slashing parameters",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/parameters", slashing.QuerierRoute)
			res, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params slashing.Params
			cdc.MustUnmarshalJSON(res, &params)
			return cliCtx.PrintOutput(params)
		},
	}
}

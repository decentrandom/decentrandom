package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/decentrandom/decentrandom/x/rand/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/tendermint/tendermint/crypto/merkle"
)

/*
type hashItem []byte
*/

/*
func (hI hashItem) Hash() []byte {
	return []byte(hI)
}
*/

// GetTxCmd -
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	randTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Rand transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	randTxCmd.AddCommand(flags.PostCommands(
		GetCmdNewRound(cdc),
		GetCmdAddTargets(cdc),
		GetCmdUpdateTargets(cdc),
	)...)

	return randTxCmd
}

// GetCmdNewRound -
func GetCmdNewRound(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "new-round [difficulty] [nonce_hash] [target1,target2,...,target(n)] [deposit] [scheduled_time]",
		Short: "Create new round data",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			/*
				if err := cliCtx.EnsureAccountExists(); err != nil {
					return err
				}
			*/

			// string to uint8
			difficulty64, errConvert := strconv.ParseInt(args[0], 8, 8)
			if errConvert != nil {
				panic(errConvert)
			}
			difficulty := uint8(difficulty64)

			// string to slice
			cleaned := strings.Replace(args[2], ",", " ", -1)
			targets := strings.Fields(cleaned)

			// string to time.Time
			var scheduledTime time.Time
			if args[4] != "" {
				var err error
				// Sample : "2014-09-12T11:45:26.371Z"
				scheduledTime, err = time.Parse(time.RFC3339, args[4])

				if err != nil {
					panic(err)
				}

			} else {

				scheduledTime = time.Now()
			}

			// Create ID
			roundArgs := make([][]byte, 7)
			roundArgs[0] = []byte(args[0])
			roundArgs[1] = []byte(args[1])
			roundArgs[2] = []byte(args[2])
			roundArgs[3] = []byte(args[3])
			roundArgs[4] = []byte(args[4])
			roundArgs[5] = []byte(cliCtx.GetFromAddress().String())
			roundArgs[6] = []byte(time.Now().String())

			rootHash := merkle.SimpleHashFromByteSlices(roundArgs)

			depositAmount, _ := sdk.NewIntFromString(args[3])

			msg := types.NewMsgNewRound(fmt.Sprintf("%X", []byte(rootHash)), difficulty, cliCtx.GetFromAddress(), args[1], targets, sdk.NewCoin("urand", depositAmount), scheduledTime)
			errValidate := msg.ValidateBasic()
			if errValidate != nil {
				return errValidate
			}

			//cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})

		},
	}
}

// GetCmdDeployNonce -
func GetCmdDeployNonce(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-nonce [id] [nonce]",
		Short: "Deploy Nonce to network",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			/*
				if err := cliCtx.EnsureAccountExists(); err != nil {
					return err
				}
			*/

			msg := types.NewMsgDeployNonce(args[0], cliCtx.GetFromAddress(), args[1])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			//cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

// GetCmdAddTargets -
func GetCmdAddTargets(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		/*
			****** important
			might be changed like this, target1, target2, target3, ....
		*/
		Use:   "add-targets [id] [target1,target2,...,target(n)]",
		Short: "Insert target data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			/*
				if err := cliCtx.EnsureAccountExists(); err != nil {
					return err
				}
			*/

			// string to slice
			cleaned := strings.Replace(args[1], ",", " ", -1)
			strSlice := strings.Fields(cleaned)

			msg := types.NewMsgAddTargets(args[0], cliCtx.GetFromAddress(), strSlice)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			//cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

// GetCmdUpdateTargets -
func GetCmdUpdateTargets(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "update-targets [id] [target1,target2,...,target(n)]",
		Short: "Update target data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			inBuf := bufio.NewReader(cmd.InOrStdin())

			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			/*
				if err := cliCtx.EnsureAccountExists(); err != nil {
					return err
				}
			*/

			// 컴마로 구분된 string을 slice로 변환
			cleaned := strings.Replace(args[1], ",", " ", -1)
			strSlice := strings.Fields(cleaned)

			msg := types.NewMsgUpdateTargets(args[0], cliCtx.GetFromAddress(), strSlice)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			//cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

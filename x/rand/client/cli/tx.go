package cli

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/decentrandom/decentrandom/x/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

// GetCmdNewRound -
func GetCmdNewRound(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "new-round [difficulty] [nonce] [targets] [scheduled_time]",
		Short: "set the value associate with a round that you want to initialize",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			difficulty64, errConvert := strconv.ParseInt(args[0], 16, 16)
			if errConvert != nil {
				panic(errConvert)
			}
			difficulty := int16(difficulty64)

			newID := "test"             // ***** important : to-do
			nonceHash := "hashed_nonce" // ***** important : to-do

			var targets []string // ***** important : to-do

			layout := "2014-09-12T11:45:26.371Z"          // ***** important : to-do
			str := "2014-11-12T11:45:26.371Z"             // ***** important : to-do
			scheduledTime, err := time.Parse(layout, str) // ***** important : to-do
			if err != nil {
				panic(err)
			}
			msg := rand.NewMsgNewRound(newID, difficulty, cliCtx.GetFromAddress(), nonceHash, targets, scheduledTime)
			errValidate := msg.ValidateBasic()
			if errValidate != nil {
				return errValidate
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})

		},
	}
}

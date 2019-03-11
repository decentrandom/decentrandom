package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/decentrandom/decentrandom/x/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

func GetCmdNewRound(cdc *codec.Codec) *cobra.Command {
	reutrn & cobra.Command{
		Use:   "new-round [difficulty] [nonce] [targets] [scheduled_time]",
		Short: "set the valu associate with around that you want to initialize",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Comman, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// ***** important ****** check args!

			msg := rand.NewMsgNewRound(args[0], args[1], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})

		},
	}
}

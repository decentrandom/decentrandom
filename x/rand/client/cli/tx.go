package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/crypto/tmhash"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/decentrandom/decentrandom/x/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
)

type hashItem []byte

func (hI hashItem) Hash() []byte {
	return []byte(hI)
}

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

			newID := "test" // ***** important : to-do

			roundArgs := make([][]byte, 5)
			for i := 0; i < 5; i++ {
				roundArgs[i] = hashItem(cmn.RandBytes(tmhash.Size))
			}

			rootHash := SimpleProofsFromByteSlices(roundArgs)

			// Nonce를 주소와 함께 SHA256으로 해시
			hasher := tmhash.New()
			nonceVector := []byte(fmt.Sprintf("%s%s", args[1], cliCtx.GetFromAddress()))
			hasher.Write(nonceVector)
			bz := tmhash.Sum(nonceVector)
			nonceHash := hex.EncodeToString(bz)

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

// GetCmdDeployNonce -
func GetCmdDeployNonce(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deploy-nonce [id] [nonce]",
		Short: "set the nonce associated with a ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nonceInt64, errInt := strconv.ParseInt(args[1], 10, 64)
			if errInt != nil {
				return errInt
			}
			nonceInt16 := int16(nonceInt64)

			msg := rand.NewMsgDeployNonce(args[0], cliCtx.GetFromAddress(), nonceInt16)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

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
		Use:   "add-targets [id] [value]",
		Short: "add targets",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// first, clean/remove the comma
			cleaned := strings.Replace(args[1], ",", " ", -1)

			// convert 'clened' comma separated string to slice
			strSlice := strings.Fields(cleaned)

			msg := rand.NewMsgAddTargets(args[0], cliCtx.GetFromAddress(), strSlice)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

// GetCmdRemoveTargets -
func GetCmdRemoveTargets(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		/*
			****** important
			might be changed like this, target1, target2, target3, ....
		*/
		Use:   "remove-targets [id] [value]",
		Short: "remove targets",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// first, clean/remove the comma
			cleaned := strings.Replace(args[1], ",", " ", -1)

			// convert 'clened' comma separated string to slice
			strSlice := strings.Fields(cleaned)

			msg := rand.NewMsgRemoveTargets(args[0], cliCtx.GetFromAddress(), strSlice)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
}

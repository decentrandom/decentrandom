package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"

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

// GetCmdNewRound - 신규 라운드 생성
func GetCmdNewRound(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "new-round [difficulty] [nonce] [target1,target2,...,target(n)] [scheduled_time]",
		Short: "set the value associate with a round that you want to initialize",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// string 타입의 난이도를 int16으로 변경
			difficulty64, errConvert := strconv.ParseInt(args[0], 16, 16)
			if errConvert != nil {
				panic(errConvert)
			}
			difficulty := int16(difficulty64)

			// Nonce를 주소와 함께 SHA256으로 해시
			hasher := tmhash.New()
			nonceVector := []byte(fmt.Sprintf("%s%s", args[1], cliCtx.GetFromAddress()))
			hasher.Write(nonceVector)
			bz := tmhash.Sum(nonceVector)
			nonceHash := hex.EncodeToString(bz)

			// 컴마로 구분된 string을 slice로 변환
			cleaned := strings.Replace(args[2], ",", " ", -1)
			targets := strings.Fields(cleaned)

			// string 타입의 시간을 time.Time 으로 변환
			layout := "2014-09-12T11:45:26.371Z"
			scheduledTime, err := time.Parse(layout, args[3])
			if err != nil {
				panic(err)
			}

			// ID 생성, 파라메터 값들과 계정의 머클트리 해시를 이용
			roundArgs := make([][]byte, 5)
			roundArgs[0] = []byte(args[0])
			roundArgs[1] = []byte(args[1])
			roundArgs[2] = []byte(args[2])
			roundArgs[3] = []byte(args[3])
			roundArgs[4] = []byte(cliCtx.GetFromAddress().String())

			rootHash := merkle.SimpleHashFromByteSlices(roundArgs)

			msg := rand.NewMsgNewRound(fmt.Sprintf("%X", []byte(rootHash)), difficulty, cliCtx.GetFromAddress(), nonceHash, targets, scheduledTime)
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

			// string 타입의 Nonce를 int16으로 변환
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
		Use:   "add-targets [id] [target1,target2,...,target(n)]",
		Short: "add targets",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// 컴마로 구분된 string을 slice로 변환
			cleaned := strings.Replace(args[1], ",", " ", -1)
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

// GetCmdRemoveTargets - 타겟 데이터 삭제
func GetCmdRemoveTargets(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "remove-targets [id] [target1,target2,...,target(n)]",
		Short: "remove targets",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// 컴마로 구분된 string을 slice로 변환
			cleaned := strings.Replace(args[1], ",", " ", -1)
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

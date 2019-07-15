package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/decentrandom/decentrandom/x/rand"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"

	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

/*
type hashItem []byte
*/

/*
func (hI hashItem) Hash() []byte {
	return []byte(hI)
}
*/

// GetCmdNewRound - 신규 라운드 생성
func GetCmdNewRound(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "new-round [difficulty] [nonce] [target1,target2,...,target(n)] [scheduled_time]",
		Short: "신규 라운드 생성을 위한 명령어입니다. 날짜는 yyyy-mm-ddThh:mm:ss.iiiZ 형식으로 기입해야 합니다.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			// string 타입의 난이도를 uint8으로 변경
			difficulty64, errConvert := strconv.ParseInt(args[0], 8, 8)
			if errConvert != nil {
				panic(errConvert)
			}
			difficulty := uint8(difficulty64)

			// Nonce를 해시
			hasher := tmhash.New()
			nonceVector := []byte(args[1])
			_, hashError := hasher.Write(nonceVector)
			if hashError != nil {
				return hashError
			}
			bz := tmhash.Sum(nonceVector)
			nonceHash := hex.EncodeToString(bz)

			// 컴마로 구분된 string을 slice로 변환
			cleaned := strings.Replace(args[2], ",", " ", -1)
			targets := strings.Fields(cleaned)

			// string 타입의 시간을 time.Time 으로 변환
			var scheduledTime time.Time
			if args[3] != "" {
				var err error
				// Sample : "2014-09-12T11:45:26.371Z"
				scheduledTime, err = time.Parse(time.RFC3339, args[3])

				if err != nil {
					panic(err)
				}

			} else {

				scheduledTime = time.Now()
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
		Short: "논스를 배포하기 위한 명령어 입니다. 라운드 소유자만 실행할 수 있습니다.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			msg := rand.NewMsgDeployNonce(args[0], cliCtx.GetFromAddress(), args[1])
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
		Short: "모집단 입력을 위한 명령어입니다. 라운드 소유자만 실행할 수 있습니다.",
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
		Short: "기 입력된 모집단을 삭제하기 위한 명령어입니다. 라운드 소유자만 실행할 수 있습니다.",
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

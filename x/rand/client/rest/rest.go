package rest

import (
	"fmt"
	"time"
	//"google.golang.org/grpc/balancer/base"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/decentrandom/decentrandom/x/rand"

	"github.com/gorilla/mux"
)

const (
	restRound = "round"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), newRoundHandler(cdc, cliCtx, storeName)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), addTargetsHandler(cdc, cliCtx, storeName)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), deployNonceHandler(cdc, cliCtx, storeName)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), removeTargetsHandler(cdc, cliCtx, storeName)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds/{%s}/round", storeName, restRound), roundHandler(cdc, cliCtx, storeName)).Methods("GET")

}

// Query Handler(s)
// roundHandler - 라운드 정보 handler
func roundHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restRound]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

// TX Handlers
// newRoundReq - 라운드 생성 요청 구조체
type newRoundReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Difficulty    int16        `json:"difficulty"`
	NonceHash     string       `json:"nonce_hash"`
	ID            string       `json:"id"`
	Owner         string       `json:"owner"`
	Targets       []string     `json:"targets"`
	ScheduledTime time.Time    `jsong:"scheduled_time"`
}

// newRoundHandler - 라운드 생성 handler
func newRoundHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req newRoundReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "요청을 해독하는데 실패했습니다.")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := rand.NewMsgNewRound(req.ID, req.Difficulty, addr, req.NonceHash, req.Targets, req.ScheduledTime)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

// deployNonceReq - Nonce 배포 구조체
type deployNonceReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Nonce   int16        `json:"nonce"`
	ID      string       `json:"id"`
	Owner   string       `json:"owner"`
}

// deployNonceHandler - Nonce 배포 handler
func deployNonceHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deployNonceReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "요청을 해독하는데 실패했습니다.")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := rand.NewMsgDeployNonce(req.ID, addr, req.Nonce)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})
		//clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)

	}
}

// addTargetsReq - 라운드 모집단 추가 구조체
type addTargetsReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Targets []string     `json:"targets"`
	ID      string       `json:"id"`
	Owner   string       `json:"owner"`
}

// addTargetsHandler - 라운드 모집단 추가 handler
func addTargetsHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addTargetsReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "요청을 해독하는데 실패했습니다.")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := rand.NewMsgAddTargets(req.ID, addr, req.Targets)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

// removeTargetsReq - 라운드 모집단 삭제 구조체
type removeTargetsReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Targets []string     `json:"targets"`
	ID      string       `json:"id"`
	Owner   string       `json:"owner"`
}

// removeTargetsHandler - 라운드 모집단 삭제 handler
func removeTargetsHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req removeTargetsReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "요청을 해독하는데 실패했습니다.")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := rand.NewMsgRemoveTargets(req.ID, addr, req.Targets)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

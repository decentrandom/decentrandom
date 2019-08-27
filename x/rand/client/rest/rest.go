package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	//"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/decentrandom/decentrandom/x/rand/types"
)

const (
	restRound = "round"
)

// RegisterRoutes -
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), newRoundHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), addTargetsHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), deployNonceHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), updateTargetsHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds/{%s}/round", storeName, restRound), roundHandler(cliCtx, storeName)).Methods("GET")
}

// Query Handler(s)
// roundHandler -
func roundHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//vars := mux.Vars(r)
		//paramType := vars[restRound]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/rounds", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

// TX Handlers
// newRoundReq -
type newRoundReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Difficulty    uint8        `json:"difficulty"`
	NonceHash     string       `json:"nonce_hash"`
	ID            string       `json:"id"`
	Owner         string       `json:"owner"`
	Targets       []string     `json:"targets"`
	ScheduledTime time.Time    `jsong:"scheduled_time"`
}

// newRoundHandler -
func newRoundHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req newRoundReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
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

		msg := types.NewMsgNewRound(req.ID, req.Difficulty, addr, req.NonceHash, req.Targets, req.ScheduledTime)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

// deployNonceReq -
type deployNonceReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Nonce   string       `json:"nonce"`
	ID      string       `json:"id"`
	Owner   string       `json:"owner"`
}

// deployNonceHandler -
func deployNonceHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deployNonceReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
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

		msg := types.NewMsgDeployNonce(req.ID, addr, req.Nonce)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
		//clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)

	}
}

// addTargetsReq -
type addTargetsReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Targets []string     `json:"targets"`
	ID      string       `json:"id"`
	Owner   string       `json:"owner"`
}

// addTargetsHandler -
func addTargetsHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req addTargetsReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
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

		msg := types.NewMsgAddTargets(req.ID, addr, req.Targets)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

// updateTargetsReq -
type updateTargetsReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Targets []string     `json:"targets"`
	ID      string       `json:"id"`
	Owner   string       `json:"owner"`
}

// updateTargetsHandler -
func updateTargetsHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateTargetsReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Cannot read request")
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

		msg := types.NewMsgUpdateTargets(req.ID, addr, req.Targets)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})

	}
}

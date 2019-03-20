package rest

import (
	"fmt"
	"google.golang.org/grpc/balancer/base"
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
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), newRoundHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), addTargetsHandler(cdc, cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), deployNonceHandler(cdc, cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), removeTargetsHandler(cdc, cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds/{%s}/round", storeName, restName), roundHandler(cdc, cliCtx, storeName)).Methods("GET")

}

// Query Handler(s)
// roundHandler -
func roundHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restRound]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", storeNamem paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

// TX Handlers
// newRoundReq -
type newRoundReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Difficulty    int16        `json:"difficulty"`
	NonceHash     string       `json:"nonce_hash"`
	ID            string       `json:"id"`
	Owner         string       `json:"owner"`
	Targets       []string     `json:"targets"`
	ScheduledTime time.Time    `jsong:"scheduled_time"`
}

// newRoundHandler -
func newRoundHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req newRoundReq

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
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

		msg := rand.NewMsgNewRound(req.ID, req.Difficulty, req.Owner, req.NonceHash, req.Targets, req.ScheduledTime)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientrest.CompleteAndBroadcastTxREST(w, cliCtx, baseReq, []sdk.Msg{msg}, cdc)

	}
}

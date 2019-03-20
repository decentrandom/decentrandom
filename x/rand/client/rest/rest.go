package rest

import (
	"fmt"
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
	restName = "round"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), newRoundHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), addTargetsHandler(cdc, cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), deployNonceHandler(cdc, cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), removeTargetsHandler(cdc, cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/rounds/{%s}/round", storeName, restName), roundHandler(cdc, cliCtx, storeName)).Methods("GET")

}

// newRoundReq -
type newRoundReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID      string       `json:"id"`
}

// newRoundHandler -
func newRoundHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		paramType := vars[restName]

		res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/round/%s", storeName, paramType), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound.err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)

	}
}

package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	clientrest "github.com/decentrandom/decentrandom/client/rest"
	"github.com/decentrandom/decentrandom/x/rand"

	"github.com/gorilla/mux"
)

const (
	restName = "rand"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/rounds", storeName), NewRoundHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/$s/rounds", storeName), AddTargetsHandler(cdc, cliCtx)).Methods("PUT")
}

func newRoundHandler(cdc *codec.Codec, cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	/*
		to-do
	*/
}

type newRoundReq struct {
	/*
		to-do
	*/
}

package rand

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// QueryRound -
const (
	QueryRound = "round"
)

// NewQuerier -
func NewQuerier(keeper Keeper) sdk.Querier {

	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryRound:
			return queryRound(ctx, path[1:], req, keeper)

		default:
			return nil, sdk.ErrUnknownRequest("unknown rand query endpoint.")
		}
	}
}

func queryRound(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]
	round := keeper.GetRound(ctx, id)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, round)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil

}

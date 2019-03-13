package rand

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// QueryIds - query for IDs
const (
	QueryIds = "ids"
)

// NewQuerier -
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryIds:
			return queryIds(ctx, req, keeper)

		default:
			return nil, sdk.ErrUnknownRequest("unknown rand query endpoint.")
		}
	}
}

func queryIds(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {

	var idsList QueryResIDs

	iterator := keeper.GetIDsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		idsList = append(idsList, id)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, idsList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// QueryResIDs -
type QueryResIDs []string

func (n QueryResIDs) String() string {
	return strings.Join(n[:], "\n")
}

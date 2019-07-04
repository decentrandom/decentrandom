package types

import (
	"github.com/golang/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/types"
)

func (c Context) BlockHeader() abci.Header {
	return c.Value(contextKeyBlockHeader).(abci.Header)
}

func (c Context) WithBlockHeader(header abci.Header) Context {
	var _ proto.Message = &header // for cloning.
	return c.withValue(contextKeyBlockHeader, header)
}
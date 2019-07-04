package types

import (
	"github.com/golang/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/types"
)

type Header struct {
	abci.Header
	SealedSeedHash1 []byte
	SealedSeedHash2 []byte
	SealedSeedHash3 []byte
	SealedSeedHash4 []byte
	SealedSeedHash5 []byte
	SeedHash1       []byte
	SeedHash2       []byte
	SeedHash3       []byte
	SeedHash4       []byte
	SeedHash5       []byte
}

func (c Context) BlockHeader() abci.Header {
	return c.Value(contextKeyBlockHeader).(abci.Header)
}

func (c Context) WithBlockHeader(header abci.Header) Context {
	var _ proto.Message = &header // for cloning.
	return c.withValue(contextKeyBlockHeader, header)
}
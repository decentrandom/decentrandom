package abci

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

// Custom Header
type Header struct {
	abci.Header
	SealedSeedHash1 []byte
	SealedSeedHash2 []byte
	SealedSeedHash3 []byte
	SealedSeedHash4 []byte
	SealedSeedHash5 []byte

	SeedHash1 []byte
	SeedHash2 []byte
	SeedHash3 []byte
	SeedHash4 []byte
	SeedHash5 []byte
}

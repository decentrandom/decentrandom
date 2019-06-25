package types

import (
	"sync"
	"time"

	cmn "github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/version"
)

// Block defines the atomic unit of a Tendermint blockchain.
type Block struct {
	mtx          sync.Mutex
	Header       `json:"header"`
	tmtypes.Data `json:"data"`
	Evidence     tmtypes.EvidenceData `json:"evidence"`
	LastCommit   *tmtypes.Commit      `json:"last_commit"`
}

// Header defines the structure of a Tendermint block header.
// NOTE: changes to the Header should be duplicated in:
// - header.Hash()
// - abci.Header
// - /docs/spec/blockchain/blockchain.md
type Header struct {
	// basic block info
	Version  version.Consensus `json:"version"`
	ChainID  string            `json:"chain_id"`
	Height   int64             `json:"height"`
	Time     time.Time         `json:"time"`
	NumTxs   int64             `json:"num_txs"`
	TotalTxs int64             `json:"total_txs"`

	// prev block info
	LastBlockID tmtypes.BlockID `json:"last_block_id"`

	// hashes of rand data
	SealedSeedHash1 cmn.HexBytes `json:"sealed_seed_hash_1"`
	SealedSeedHash2 cmn.HexBytes `json:"sealed_seed_hash_2"`
	SealedSeedHash3 cmn.HexBytes `json:"sealed_seed_hash_3"`
	SealedSeedHash4 cmn.HexBytes `json:"sealed_seed_hash_4"`
	SealedSeedHash5 cmn.HexBytes `json:"sealed_seed_hash_5"`
	SeedHash1       cmn.HexBytes `json:"seed_hash_1"`
	SeedHash2       cmn.HexBytes `json:"seed_hash_2"`
	SeedHash3       cmn.HexBytes `json:"seed_hash_3"`
	SeedHash4       cmn.HexBytes `json:"seed_hash_4"`
	SeedHash5       cmn.HexBytes `json:"seed_hash_5"`

	// hashes of block data
	LastCommitHash cmn.HexBytes `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       cmn.HexBytes `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     cmn.HexBytes `json:"validators_hash"`      // validators for the current block
	NextValidatorsHash cmn.HexBytes `json:"next_validators_hash"` // validators for the next block
	ConsensusHash      cmn.HexBytes `json:"consensus_hash"`       // consensus params for current block
	AppHash            cmn.HexBytes `json:"app_hash"`             // state after txs from the previous block
	LastResultsHash    cmn.HexBytes `json:"last_results_hash"`    // root hash of all results from the txs from the previous block

	// consensus info
	EvidenceHash    cmn.HexBytes    `json:"evidence_hash"`    // evidence included in the block
	ProposerAddress tmtypes.Address `json:"proposer_address"` // original proposer of the block
}

package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

func (c Context) BlockHeader() abci.Header {
	return c.Value(contextKeyBlockHeader).(abci.Header)
}

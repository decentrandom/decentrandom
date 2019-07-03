package types

import (
	"context"
)

// Context -
type Context struct {
	context.Context
	pst *thePast
	gen int
}

// NewContext -
func NewContext(ms Multistore, header abci.Header, isCheckTx bool, logger log.Logger) Context {

}

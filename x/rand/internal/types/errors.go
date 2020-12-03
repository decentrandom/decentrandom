package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// ErrCodeRoundNotFound -
	ErrCodeRoundNotFound = 1
)

var (
	// DefaultCodespace -
	DefaultCodespace = ModuleName

	// ErrRoundNotFound -
	ErrRoundNotFound = sdkerrors.Register(ModuleName, ErrCodeRoundNotFound, "round not found")
)

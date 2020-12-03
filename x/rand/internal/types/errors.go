package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ErrCodeRoundNotFound = 1
)

var (
	DefaultCodespace = ModuleName

	ErrRoundNotFound = sdkerrors.Register(ModuleName, ErrCodeRoundNotFound, "round not found")
)

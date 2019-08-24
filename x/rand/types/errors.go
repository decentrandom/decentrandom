package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultCodespace -
	DefaultCodespace sdk.CodespaceType = ModuleName

	// CodeRoundDoesNotExist -
	CodeRoundDoesNotExist sdk.CodeType = 101
)

// ErrRoundDoesNotExist -
func ErrRoundDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeRoundDoesNotExist, "Round does not exist")
}

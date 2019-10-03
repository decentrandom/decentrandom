package types

import "strings"

// QueryResRoundIDs -
type QueryResRoundIDs []string

func (r QueryResRoundIDs) String() string {
	return strings.Join(r[:], "\n")
}

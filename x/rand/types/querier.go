package types

import "strings"

// QueryResRoundIDs -
type QueryResRoundIDs struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResRoundIDs) String() string {
	return r.Value
}

// QueryRoundInfo -
type QueryRoundInfo []string

// implement fmt.Stringer
func (n QueryRoundInfo) String() string {
	return strings.Join(n[:], "\n")
}

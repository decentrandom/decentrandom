package assets

const (
	// MicroRandDenom -
	MicroRandDenom = "urand"

	// MicroUnit -
	MicroUnit = int64(1e6)
)

// IsValidDenom -
func IsValidDenom(denom string) bool {
	return denom == MicroRandDenom
}

package assets

// nolint: golint
const (
	MicroRandDenom = "urand"

	MicroUnit = int64(1e6)
)

// IsValidDenom -
func IsValidDenom(denom string) bool {
	return denom == MicroRandDenom
}

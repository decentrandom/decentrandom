package assets

// nolint: golint
const (
	MicroRandDenom = "mrand"

	MicroUnit = int64(1e6)
)

// IsValidDenom -
func IsValidDenom(denom string) bool {
	return denom == MicroRandDenom
}

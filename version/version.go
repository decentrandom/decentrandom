package version

import (
	"fmt"
	"runtime"
)

var (
	Commit    = ""
	Version   = ""
	GoSumHash = ""
	BuildTags = ""
)

// versionInfo -
type versionInfo struct {
	DecentRandom string `json:"decent_random"`
	GitCommit    string `json:"commit"`
	GoSumHash    string `json:"gosum_hash"`
	BuildTags    string `json:"build_tags"`
	GoVersion    string `json:"go"`
}

// String -
func (v versionInfo) String() string {
	return fmt.Sprintf(`decent-random: %s
git commit: %s
go.sum hash: %s
build tags: %s
%s`, v.DecentRandom, v.GitCommit, v.GoSumHash, v.BuildTags, v.GoVersion)
}

// newVersionInfo -
func newVersionInfo() versionInfo {
	return versionInfo{
		Version,
		Commit,
		GoSumHash,
		BuildTags,
		fmt.Sprintf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)}
}

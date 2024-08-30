package version

import (
	"fmt"
	"runtime"
)

type VersionInfo struct {
	Version   string
	GoVersion string
	GitBranch string
	BuildTime string
	GitCommit string
}

var (
	Info = VersionInfo{GoVersion: runtime.Version()}
)

func (a *VersionInfo) About() {
	fmt.Printf(
		"Version: %s GoVersion: %s GitBranch: %s BuildTime: %s GitCommit: %s\n",
		a.Version,
		a.GoVersion,
		a.GitBranch,
		a.BuildTime,
		a.GitCommit,
	)
}

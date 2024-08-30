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
	Info = VersionInfo{
		Version:   Version,
		GoVersion: runtime.Version(),
		GitBranch: GitBranch,
		BuildTime: BuildTime,
		GitCommit: GitBranch,
	}
	Version   string
	GoVersion string
	GitBranch string
	BuildTime string
	GitCommit string
)

func About() {
	fmt.Printf(
		"Version: %s GoVersion: %s GitBranch: %s BuildTime: %s GitCommit: %s\n",
		Info.Version,
		Info.GoVersion,
		Info.GitBranch,
		Info.BuildTime,
		Info.GitCommit,
	)
}

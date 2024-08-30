package version

import "fmt"

type About struct {
	Version   string
	GoVersion string
	GitBranch string
	BuildTime string
	GitCommit string
}

func (a About) String() {
	fmt.Printf(
		"Version: %s GoVersion: %s GitBranch: %s BuildTime: %s GitCommit: %s",
		a.Version,
		a.GoVersion,
		a.GitBranch,
		a.BuildTime,
		a.GitCommit,
	)
}

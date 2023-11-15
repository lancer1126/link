package version

import "fmt"

const (
	series = "v0.5"
)

var (
	version   = "unknown"
	commitID  = "unknown"
	buildDate = "unknown"
	buildTags = "unknown"
)

type BuildInfo struct {
	Series    string `json:"series"`
	Version   string `json:"version"`
	Sum       string `json:"sum"`
	BuildDate string `json:"build_date"`
	BuildTags string `json:"build_tags"`
}

func Info() string {
	return fmt.Sprintf("paopao %s (build:%s %s)", version, commitID, buildDate)
}

func ReadBuildInfo() *BuildInfo {
	return &BuildInfo{
		Series:    series,
		Version:   version,
		Sum:       commitID,
		BuildDate: buildDate,
		BuildTags: buildTags,
	}
}

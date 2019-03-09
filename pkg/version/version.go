package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

var (
	//versionFile = "/src/github.com/kubernaut/kubernaut/VERSION"
	GitCommit      string
	ReleaseVersion = "master"
)

type Version struct {
	Version   string
	Commit    string
	BuildDate string
	GoVersion string
	GOOS      string
	GOArch    string
}

func GetVersion() *Version {
	return &Version{
		Version:   ReleaseVersion,
		Commit:    GitCommit,
		BuildDate: time.Now().UTC().String(),
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
		GOArch:    runtime.GOARCH,
	}
}

func GetVersionShort() string {
	return fmt.Sprintf("kubernaut %s", GetVersion().Version)
}

func GetVersionJSON() string {
	verBytes, err := json.Marshal(GetVersion())

	if err != nil {
		return `{"Version": "unknown"}`
	}

	return string(verBytes)
}

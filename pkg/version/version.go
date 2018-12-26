package version

import (
	"fmt"
	"os/exec"
	"strings"
)

var (
	version string
)

func GetVersion() string {
	if version == "" {
		out, _ := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		version = fmt.Sprintf("dev-%s", strings.Trim(string(out), "\r\n"))
	}
	return version
}

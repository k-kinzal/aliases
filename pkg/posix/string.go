package posix

import (
	"os/exec"
	"strings"
)

// see: http://pubs.opengroup.org/onlinepubs/009695399/utilities/xcu_chap02.html
func String(cmd exec.Cmd) string {
	return strings.Join(cmd.Args, " ")
}

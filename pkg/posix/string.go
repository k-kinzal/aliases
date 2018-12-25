package posix

import (
	"fmt"
	"os/exec"
	"strings"
)

// see: http://pubs.opengroup.org/onlinepubs/009695399/utilities/xcu_chap02.html
func String(cmd exec.Cmd) string {
	return fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args[1:], " "))
}

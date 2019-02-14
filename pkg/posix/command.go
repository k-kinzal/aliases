package posix

import (
	"fmt"
	"os/exec"
	"path"
	"strings"
)

// Cmd is the base of the POSIX commands.
type Cmd struct {
	*exec.Cmd
}

// String returns command string.
func (cmd *Cmd) String() string {
	return fmt.Sprintf("%s %s", path.Base(cmd.Cmd.Args[0]), strings.Join(cmd.Cmd.Args[1:], " "))
}

// Command creates a new posix.Cmd.
func Command(name string, arg ...string) *Cmd {
	return &Cmd{exec.Command(name, arg...)}
}

package docker

import (
	"os"
	"os/exec"
)

func NewRunCmd(opt *RunOpts) *exec.Cmd {
	cmd := exec.Command("docker", "run")

	cmd.Args = append(cmd.Args, opt.toArguments()...)

	cmd.Env = os.Environ()

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

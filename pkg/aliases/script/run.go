package script

import (
	"os"

	"github.com/k-kinzal/aliases/pkg/logger"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/posix"
)

// Run aliases script.
func (script *Script) Run(args []string, opt docker.RunOption) error {
	for _, relative := range script.relative {
		if _, err := relative.Write(); err != nil {
			return err
		}
	}

	dockerCmdString := script.docker(args, opt).String()
	logger.Debug(dockerCmdString)
	cmd := posix.Shell(dockerCmdString)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

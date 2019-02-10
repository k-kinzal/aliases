package script

import (
	"os"

	"github.com/k-kinzal/aliases/pkg/aliases"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/posix"
)

// Run aliases script.
func (script *Script) Run(ctx aliases.Context, args []string, opt docker.RunOption) error {
	for _, relative := range script.relative {
		if _, err := relative.Write(ctx); err != nil {
			return err
		}
	}

	cmd := posix.Shell(script.docker(args, opt).String())
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

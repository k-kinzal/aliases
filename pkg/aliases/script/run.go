package script

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k-kinzal/aliases/pkg/logger"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/posix"
)

// Run aliases script.
func (script *Script) Run(args []string, opt docker.RunOption) error {
	path, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		return err
	}

	for _, relative := range script.relative {
		if err := relative.Write(path); err != nil {
			return err
		}
	}

	cmd := posix.Shell(script.docker(args, opt).String())
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = append(cmd.Env, fmt.Sprintf("ALIASES_EXPORT_PATH=%s", path))

	logger.Debug(cmd.String())

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

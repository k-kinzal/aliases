package export

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	pathes "path"
	"path/filepath"

	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/posix"
)

const tmpl = `#!/bin/sh

if [ -p /dev/stdin ]; then
  cat - | %s "$@"
  exit $?
else
  %s "$@"
  exit $?
fi
`

func Script(ctx context.Context, commands map[string]exec.Cmd) error {
	if err := os.RemoveAll(ctx.ExportPath()); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}
	if err := os.Mkdir(ctx.ExportPath(), 0755); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}

	for path, cmd := range commands {
		str := posix.String(cmd)
		writePath := filepath.Join(ctx.ExportPath(), pathes.Base(path))
		content := fmt.Sprintf(tmpl, str, str)
		if err := ioutil.WriteFile(writePath, []byte(content), 0755); err != nil {
			return fmt.Errorf("runtime error: %s", err)
		}
	}

	return nil
}

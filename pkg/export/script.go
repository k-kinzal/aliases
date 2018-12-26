package export

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	pathes "path"
	"path/filepath"
	"strings"

	"github.com/k-kinzal/aliases/pkg/context"
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

func Script(ctx *context.Context, commands map[string]exec.Cmd) error {
	if err := os.RemoveAll(ctx.GetExportPath()); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}
	if err := os.Mkdir(ctx.GetExportPath(), 0755); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}

	for path, cmd := range commands {
		str := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args[1:], " "))
		writePath := filepath.Join(ctx.GetExportPath(), pathes.Base(path))
		content := fmt.Sprintf(tmpl, str, str)
		ioutil.WriteFile(writePath, []byte(content), 0755)
	}

	return nil
}

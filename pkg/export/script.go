package export

import (
	"fmt"
	"io/ioutil"
	"os/exec"

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

func Script(path string, cmd exec.Cmd) error {
	str := posix.String(cmd)
	content := fmt.Sprintf(tmpl, str, str)
	if err := ioutil.WriteFile(path, []byte(content), 0755); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}

	return nil
}

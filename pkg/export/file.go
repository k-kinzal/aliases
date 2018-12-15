package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg/conf"
	"github.com/k-kinzal/aliases/pkg/docker"
	"io/ioutil"
	"path"
	"strings"
)

var (
	tmpl = `#!/bin/sh

if [ -p /dev/stdin ]; then
  cat - | %s "$@"
  exit $?
else
  %s "$@"
  exit $?
fi
`
)

func writeFiles(conf *conf.CommandConf, dir string) {
	cmd := docker.NewRunCmd(&conf.DockerRunOpts)
	cmdStr := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))
	writePath := fmt.Sprintf("%s/%s", dir, path.Base(conf.Path))
	content := fmt.Sprintf(tmpl, cmdStr, cmdStr)
	ioutil.WriteFile(writePath, []byte(content), 0755)
}

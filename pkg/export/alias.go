package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
	"github.com/k-kinzal/aliases/pkg/docker"
	"os"
	"path"
	"strings"
)


func Aliases(conf *aliases.AliasesConf, ctx *aliases.Context) {
	dir := ctx.GetBinaryPath(conf.Hash)
	os.Remove(dir)
	os.Mkdir(dir, 0755)

	for _, c := range conf.Commands {
		for _, dep := range c.Dependencies {
			writeFiles(dep, dir)
		}
		cmd := docker.NewRunCmd(&c.DockerRunOpts)
		fmt.Printf(fmt.Sprintf("alias %s='%s %s'\n", path.Base(cmd.Path) , cmd.Path, strings.Join(cmd.Args, " ")))
	}
}
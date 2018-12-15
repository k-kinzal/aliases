package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg/conf"
	"github.com/k-kinzal/aliases/pkg/context"
	"github.com/k-kinzal/aliases/pkg/docker"
	"os"
	"path"
	"strings"
)


func Aliases(ctx *context.Context, conf *conf.AliasesConf) {
	dir := ctx.GetBinaryPath()
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
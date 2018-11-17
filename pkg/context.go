package aliases

import (
	"fmt"
	"os"
	"os/user"
)

type Context struct {
	homePath string
	confPath string
}

func (ctx *Context)GetHomePath() string {
	if ctx.homePath == "" {
		ctx.homePath = os.Getenv("ALIASES_HOME")
		if ctx.homePath == "" {
			usr, _ := user.Current()
			ctx.homePath = fmt.Sprintf("%s/.aliases", usr.HomeDir)
		}

		if _, err := os.Stat(ctx.homePath); os.IsNotExist(err) {
			os.Mkdir(ctx.homePath, 0755)
		}
	}

	return ctx.homePath
}

func (ctx *Context)GetBinaryPath(hash string) string {
	return fmt.Sprintf("%s/%s", ctx.GetHomePath(), hash)
}

func (ctx *Context)GetConfPath() string {
	if ctx.confPath == "" {
		cwd, _ := os.Getwd()
		ctx.confPath = fmt.Sprintf("%s/aliases.yaml", cwd)
	}

	return ctx.confPath
}

func NewContext(homePath string, confPath string) *Context {
	return &Context{homePath, confPath}
}

package context

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"os/user"
)

type Context struct {
	homePath string
	confPath string
	binaryPath string
}

func NewContext(
		homePath string,
		confPath string,
		binaryPath string) *Context {
	return &Context{
		homePath: homePath,
		confPath: confPath,
		binaryPath: binaryPath,
	}
}

func (ctx *Context)GetHomePath() string {
	if ctx.homePath != "" {
		return ctx.homePath
	}

	ctx.homePath = os.Getenv("ALIASES_HOME")
	if ctx.homePath == "" {
		usr, _ := user.Current()
		ctx.homePath = fmt.Sprintf("%s/.aliases", usr.HomeDir)
	}

	if _, err := os.Stat(ctx.homePath); os.IsNotExist(err) {
		os.Mkdir(ctx.homePath, 0755)
	}

	return ctx.homePath
}

func (ctx *Context)GetConfPath() string {
	if ctx.confPath != "" {
		return ctx.confPath
	}

	cwd, _ := os.Getwd()
	ctx.confPath = fmt.Sprintf("%s/aliases.yaml", cwd)

	if _, err := os.Stat(ctx.confPath); os.IsNotExist(err) {
		ctx.confPath = fmt.Sprintf("%s/aliases.yaml", ctx.GetHomePath())
	}

	return ctx.confPath
}


func (ctx *Context)GetBinaryPath() string {
	if ctx.binaryPath != "" {
		return ctx.binaryPath
	}

	hash := uuid.NewMD5(uuid.UUID{}, []byte(ctx.GetHomePath())).String()
	ctx.binaryPath = fmt.Sprintf("%s/%s", ctx.GetHomePath(), hash)

	return ctx.binaryPath
}
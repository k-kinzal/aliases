package aliases

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/k-kinzal/aliases/pkg/types"
)

type Context interface {
	HomePath() string
	ConfPath() string
	ExportPath() string
	BinaryPath() string

	MakeHomeDir() error
	MakeExportDir() error
	MakeBinaryDir() error
}

type GlobalContext struct {
	homePath   string
	confPath   string
	exportPath string
	binaryPath string
}

func (ctx *GlobalContext) HomePath() string {
	if ctx.homePath == "" {
		usr, _ := user.Current()
		ctx.homePath = fmt.Sprintf("%s/.aliases", usr.HomeDir)
	}
	return ctx.homePath
}

func (ctx *GlobalContext) ConfPath() string {
	if ctx.confPath == "" {
		cwd, _ := os.Getwd()
		ctx.confPath = fmt.Sprintf("%s/aliases.yaml", cwd)
		if _, err := os.Stat(ctx.confPath); os.IsNotExist(err) {
			ctx.confPath = fmt.Sprintf("%s/aliases.yaml", ctx.HomePath())
		}
	}
	return ctx.confPath
}

func (ctx *GlobalContext) ExportPath() string {
	if ctx.exportPath == "" {
		ctx.exportPath = fmt.Sprintf("%s/%s", ctx.HomePath(), types.MD5(ctx.ConfPath()))
	}
	return ctx.exportPath
}

func (ctx *GlobalContext) BinaryPath() string {
	if ctx.binaryPath == "" {
		ctx.binaryPath = path.Join(ctx.HomePath(), "docker")
	}
	return ctx.binaryPath
}

func (ctx *GlobalContext) MakeHomeDir() error {
	if _, err := os.Stat(ctx.HomePath()); os.IsNotExist(err) {
		if err := os.Mkdir(ctx.HomePath(), 0755); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *GlobalContext) MakeExportDir() error {
	if err := os.RemoveAll(ctx.ExportPath()); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}
	if err := os.Mkdir(ctx.ExportPath(), 0755); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}
	return nil
}

func (ctx *GlobalContext) MakeBinaryDir() error {
	if _, err := os.Stat(ctx.BinaryPath()); os.IsNotExist(err) {
		if err := os.Mkdir(ctx.BinaryPath(), 0755); err != nil {
			return err
		}
	}
	return nil
}

func NewContext(homePath string, confPath string) (Context, error) {
	ctx := new(GlobalContext)
	ctx.homePath = homePath
	ctx.confPath = confPath

	return ctx, nil
}

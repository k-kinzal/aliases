package aliases

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/k-kinzal/aliases/pkg/logger"
)

type Context interface {
	HomePath() string
	ConfPath() string
	ExportPath() string

	MakeHomeDir() error
	MakeExportDir() error

	DockerBinaryPath() string
	DockerHost() string
	DockerSocketPath() *string
	HasDockerSocket() bool
}

type GlobalContext struct {
	homePath       string
	confPath       string
	exportPath     string
	dockerBinPath  string
	dockerHost     string
	dockerSockPath *string
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
		hasher := md5.New()
		_, _ = hasher.Write([]byte(ctx.ConfPath()))
		ctx.exportPath = fmt.Sprintf("%s/%s", ctx.HomePath(), hex.EncodeToString(hasher.Sum(nil)))
	}
	return ctx.exportPath
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

func (ctx *GlobalContext) DockerBinaryPath() string {
	return ctx.dockerBinPath
}

func (ctx *GlobalContext) DockerHost() string {
	return ctx.dockerHost
}

func (ctx *GlobalContext) DockerSocketPath() *string {
	if !strings.HasPrefix(ctx.dockerHost, "unix://") {
		return nil
	}
	sock := strings.TrimPrefix(ctx.dockerHost, "unix://")
	return &sock
}

func (ctx *GlobalContext) HasDockerSocket() bool {
	sock := ctx.DockerSocketPath()
	return sock != nil && *sock != ""
}

func NewContext(homePath string, confPath string) (Context, error) {
	ctx := new(GlobalContext)
	ctx.homePath = homePath
	ctx.confPath = confPath

	cmd := exec.Command("docker")
	if cmd.Path == "docker" {
		return nil, fmt.Errorf("runtime error: docker is not installed. see https://docs.docker.com/install/")
	}
	ctx.dockerBinPath = cmd.Path

	ctx.dockerHost = os.Getenv("DOCKER_HOST")
	if ctx.dockerHost == "" {
		// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-socket-option
		ctx.dockerHost = "unix:///var/run/docker.sock"
	}

	if !strings.HasPrefix(ctx.dockerHost, "unix://") {
		logger.Warnf("%s may not working possibility. Please same path that you use on the host and the host of `DOCKER_HOST`.", ctx.dockerHost)
	} else {
		sock := strings.TrimPrefix(ctx.dockerHost, "unix://")
		if _, err := os.Stat(sock); err != nil {
			return nil, fmt.Errorf("runtime error: %s: no such file. please set DOCKER_HOST", sock)
		}
		ctx.dockerSockPath = &sock
	}

	return ctx, nil
}

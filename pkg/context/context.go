package context

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

const (
	DockerSockTypeSock   = 1
	DockerSockTypeRemote = 2
)

type Context struct {
	homePath         string
	confPath         string
	exportPath       string
	dockerPath       string
	dockerSockPath   string
	dockerRemoteHost string
}

func New(
	homePath string,
	confPath string,
	exportPath string) *Context {
	return &Context{
		homePath:   homePath,
		confPath:   confPath,
		exportPath: exportPath,
	}
}

func (ctx *Context) GetHomePath() string {
	if ctx.homePath != "" {
		return ctx.homePath
	}

	ctx.homePath = os.Getenv("ALIASES_HOME")
	if ctx.homePath == "" {
		usr, _ := user.Current()
		ctx.homePath = fmt.Sprintf("%s/.aliases", usr.HomeDir)
	}

	if _, err := os.Stat(ctx.homePath); os.IsNotExist(err) {
		if err := os.Mkdir(ctx.homePath, 0755); err != nil {
			panic(fmt.Sprintf("runtime error: %s", err)) // FIXME: handling error
		}
	}

	return ctx.homePath
}

func (ctx *Context) GetConfPath() string {
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

func (ctx *Context) GetExportPath() string {
	if ctx.exportPath != "" {
		return ctx.exportPath
	}

	hasher := md5.New()
	hasher.Write([]byte(ctx.GetHomePath()))

	ctx.exportPath = fmt.Sprintf("%s/%s", ctx.GetHomePath(), hex.EncodeToString(hasher.Sum(nil)))

	return ctx.exportPath
}

func (ctx *Context) DockerBinaryPath() string {
	if ctx.dockerPath != "" {
		return ctx.dockerPath
	}
	cmd := exec.Command("docker")
	if cmd.Path == "docker" {
		panic("runtime error: docker is not installed. see https://docs.docker.com/install/") // FIXME: handling error
	}
	ctx.dockerPath = cmd.Path

	return ctx.dockerPath
}

func (ctx *Context) DockerSockType() int {
	switch ctx.DockerSockPath() == "" {
	case true:
		return DockerSockTypeRemote
	case false:
		return DockerSockTypeSock
	}
	return -1
}

func (ctx *Context) DockerSockPath() string {
	if ctx.dockerSockPath != "" {
		return ctx.dockerSockPath
	}
	host := os.Getenv("DOCKER_HOST")
	if host == "" {
		sock := "/var/run/docker.sock"
		host = fmt.Sprintf("unix://%s", sock)
	}
	if strings.HasPrefix(host, "unix://") {
		ctx.dockerSockPath = strings.TrimPrefix(host, "unix://")
		if _, err := os.Stat(ctx.dockerSockPath); err != nil {
			panic(fmt.Sprintf("runtime error: %s: no such file. please set DOCKER_HOST", ctx.dockerSockPath)) // FIXME: handling error
		}
	} else {
		ctx.dockerSockPath = ""
	}
	return ctx.dockerSockPath
}

func (ctx *Context) DockerRemoteHost() string {
	if ctx.dockerRemoteHost != "" {
		return ctx.dockerRemoteHost
	}
	host := os.Getenv("DOCKER_HOST")
	if !strings.HasPrefix(host, "unix://") {
		ctx.dockerRemoteHost = host
	} else {
		ctx.dockerRemoteHost = ""
	}
	return ctx.dockerRemoteHost
}

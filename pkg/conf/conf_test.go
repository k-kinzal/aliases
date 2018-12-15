package conf_test

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg/conf"
	"github.com/k-kinzal/aliases/pkg/context"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"testing"
)

func TestLoadConfFile(t *testing.T) {
	content := `---
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
`
	file, _ := ioutil.TempFile("", "")
	file.Write([]byte(content))
	ctx := context.NewContext("", file.Name(), "")

	cf, err := conf.LoadConfFile(ctx)
	if err != nil {
		t.Errorf("load config error: %v", err)
	}
	if len(cf.Commands) != 1 {
		t.Errorf("expected `1`, but in actual the load configuration length is `%d`", len(cf.Commands))
	}
	command := cf.Commands[0]
	if command.Path != "/usr/local/bin/kubectl" {
		t.Error("/usr/local/bin/kubectl does not exist in load configuration")
	}
	if command.DockerRunOpts.Image != "chatwork/kubectl:${KUBECTL_VERSION:-\"1.11.2\"}" {
		t.Errorf("expected `chatwork/kubectl:${KUBECTL_VERSION:-\"1.11.2\"}`, but in actual `%s` has been set in dockerrunopts.image", command.DockerRunOpts.Image)
	}

}

func TestLoadConfFile_ShouldBeSetDefaultValue(t *testing.T) {
	content := `---
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
`
	file, _ := ioutil.TempFile("", "")
	file.Write([]byte(content))
	ctx := context.NewContext("", file.Name(), "")

	cf, err := conf.LoadConfFile(ctx)
	if err != nil {
		t.Errorf("load config error: %v", err)
	}
	command := cf.Commands[0]
	if *command.DockerRunOpts.Interactive != true {
		t.Errorf("expected `true`, but in actual `%s` has been set in dockerrunopts.interactive", strconv.FormatBool(*command.DockerRunOpts.Interactive))
	}
	if *command.DockerRunOpts.Rm != true {
		t.Errorf("expected `true`, but in actual `%s` has been set in dockerrunopts.rm", strconv.FormatBool(*command.DockerRunOpts.Rm))
	}
	if *command.DockerRunOpts.Network != "host" {
		t.Errorf("expected `host`, but in actual `%s` has been set in dockerrunopts.interactive", *command.DockerRunOpts.Network)
	}
}

func TestUnmarshalConfFile_ShouldBeSetDependenciesWithUnixSock(t *testing.T) {
	os.Setenv("DOCKER_HOST", "unix:///var/run/docker.sock")
	content := `---
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
/usr/local/bin/helmfile:
  image: chatwork/helmfile
  tag: 0.36.1-2.10.0
  dependencies:
  - /usr/local/bin/kubectl
`
	file, _ := ioutil.TempFile("", "")
	file.Write([]byte(content))
	ctx := context.NewContext("", file.Name(), "")

	cf, err := conf.LoadConfFile(ctx)
	if err != nil {
		t.Errorf("load config error: %v", err)
	}
	command := conf.CommandConf{}
	if cf.Commands[0].Path == "/usr/local/bin/helmfile" {
		command = cf.Commands[0]
	} else {
		command = cf.Commands[1]
	}

	if *command.DockerRunOpts.Privileged != true {
		t.Errorf("expected `true`, but in actual `%s` has been set in dockerrunopts.privileged", strconv.FormatBool(*command.DockerRunOpts.Privileged))
	}
	if command.DockerRunOpts.Volume[0] != "/usr/local/bin/docker:/usr/local/bin/docker" {
		t.Errorf("expected `/usr/local/bin/docker:/usr/local/bin/docker`, but in actual `%s` has been set in dockerrunopts.volume[0]", command.DockerRunOpts.Volume[0])
	}
	if command.DockerRunOpts.Volume[1] != "/var/run/docker.sock:/var/run/docker.sock" {
		t.Errorf("expected `/var/run/docker.sock:/var/run/docker.sock`, but in actual `%s` has been set in dockerrunopts.volume[1]", command.DockerRunOpts.Volume[1])
	}
	mauntPath := fmt.Sprintf("%s/%s", ctx.GetExportPath(), path.Base(command.Dependencies[0].Path))
	volume := fmt.Sprintf("%s:%s", mauntPath, command.Dependencies[0].Path)
	if command.DockerRunOpts.Volume[2] != volume {
		t.Errorf("expected `%s`, but in actual `%s` has been set in dockerrunopts.volume[2]", volume, command.DockerRunOpts.Volume[1])
	}
	if command.DockerRunOpts.Env["ALIASES_PWD"] != "${ALIASES_PWD:-$PWD}" {
		t.Errorf("expected `${ALIASES_PWD:-$PWD}`, but in actual `%s` has been set in dockerrunopts.env[\"ALIASES_PWD\"]", command.DockerRunOpts.Env["ALIASES_PWD"])
	}
}

func TestUnmarshalConfFile_ShouldBeSetDependenciesWithHost(t *testing.T) {
	os.Setenv("DOCKER_HOST", "tcp://localhost")
	content := `---
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
/usr/local/bin/helmfile:
  image: chatwork/helmfile
  tag: 0.36.1-2.10.0
  dependencies:
  - /usr/local/bin/kubectl
`
	file, _ := ioutil.TempFile("", "")
	file.Write([]byte(content))
	ctx := context.NewContext("", file.Name(), "")

	cf, err := conf.LoadConfFile(ctx)
	if err != nil {
		t.Errorf("load config error: %v", err)
	}
	command := conf.CommandConf{}
	if cf.Commands[0].Path == "/usr/local/bin/helmfile" {
		command = cf.Commands[0]
	} else {
		command = cf.Commands[1]
	}

	if *command.DockerRunOpts.Privileged != true {
		t.Errorf("expected `true`, but in actual `%s` has been set in dockerrunopts.privileged", strconv.FormatBool(*command.DockerRunOpts.Privileged))
	}
	if command.DockerRunOpts.Volume[0] != "/usr/local/bin/docker:/usr/local/bin/docker" {
		t.Errorf("expected `/usr/local/bin/docker:/usr/local/bin/docker`, but in actual `%s` has been set in dockerrunopts.volume[0]", command.DockerRunOpts.Volume[0])
	}
	mauntPath := fmt.Sprintf("%s/%s", ctx.GetExportPath(), path.Base(command.Dependencies[0].Path))
	volume := fmt.Sprintf("%s:%s", mauntPath, command.Dependencies[0].Path)
	if command.DockerRunOpts.Volume[1] != volume {
		t.Errorf("expected `%s`, but in actual `%s` has been set in dockerrunopts.volume[1]", volume, command.DockerRunOpts.Volume[1])
	}
	if command.DockerRunOpts.Env["DOCKER_HOST"] != "tcp://localhost" {
		t.Errorf("expected `tcp://localhost`, but in actual `%s` has been set in dockerrunopts.env[\"DOCKER_HOST\"]", command.DockerRunOpts.Env["ALIASES_PWD"])
	}
	if command.DockerRunOpts.Env["ALIASES_PWD"] != "${ALIASES_PWD:-$PWD}" {
		t.Errorf("expected `${ALIASES_PWD:-$PWD}`, but in actual `%s` has been set in dockerrunopts.env[\"ALIASES_PWD\"]", command.DockerRunOpts.Env["ALIASES_PWD"])
	}
}

func TestUnmarshalConfFile_ShouldBeExpandEnv(t *testing.T) {
	content := `---
/usr/local/bin/kubectl:
  image: chatwork/kubectl
  tag: 1.11.2
  env:
    KUBECONFIG: $HOME/.kube/config
    $HOST.label: foo
  volume:
  - $HOME/.kube:/root/.kube
  - $PWD:/kube
  workdir: /kube
`
	file, _ := ioutil.TempFile("", "")
	file.Write([]byte(content))
	ctx := context.NewContext("", file.Name(), "")

	cf, err := conf.LoadConfFile(ctx)
	if err != nil {
		t.Errorf("load config error: %v", err)
	}
	command := cf.Commands[0]
	if command.DockerRunOpts.Env["KUBECONFIG"] != "$HOME/.kube/config" {
		t.Errorf("expected `$HOME/.kube/config`, but in actual `%s` has been set in dockerrunopts.env[\"KUBECONFIG\"]", command.DockerRunOpts.Env["KUBECONFIG"])
	}
	envname := fmt.Sprintf("%s.label", os.Getenv("HOST"))
	if command.DockerRunOpts.Env[envname] != "foo" {
		t.Errorf("expected `foo`, but in actual `%s` has been set in dockerrunopts.env[\"%s\"]", command.DockerRunOpts.Env[envname], envname)
	}
	volume := fmt.Sprintf("%s/.kube:/root/.kube", os.Getenv("HOME"))
	if command.DockerRunOpts.Volume[0] != volume {
		t.Errorf("expected `%s`, but in actual `%s` has been set in dockerrunopts.volume[0]", volume, command.DockerRunOpts.Volume[0])
	}
	if command.DockerRunOpts.Volume[1] != "${ALIASES_PWD:-$PWD}:/kube" {
		t.Errorf("expected `${ALIASES_PWD:-$PWD}:/kube`, but in actual `%s` has been set in dockerrunopts.volume[1]", command.DockerRunOpts.Volume[0])
	}
}

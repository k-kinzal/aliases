package script_test

import (
	"io/ioutil"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleScript_Run() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  entrypoint: sh
  env:
    FOO: 1
  dependencies:
  - /path/to/command2:
      image: alpine
      tag: latest
      entrypoint: sh
      env:
        FOO: 1
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}

	opt, err := conf.Get("/path/to/command1")
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		panic(err)
	}
	if err := context.ChangeHomePath(dir); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath(dir); err != nil {
		panic(err)
	}
	for _, binary := range conf.Binaries(context.BinaryPath()) {
		if err := docker.Download(binary.Path, binary.Image, binary.Tag); err != nil {
			panic(err)
		}
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := script.NewScript(client, *opt)
	if err := cmd.Run([]string{"-c", "'/path/to/command2 -c '\"'\"'echo $FOO'\"'\"''"}, docker.RunOption{}); err != nil {
		panic(err)
	}
	// Output:
	// 1
}

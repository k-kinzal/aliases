package script_test

import (
	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleNewScript() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
  entrypoint: sh
  args: [-c]
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}

	opt, err := conf.Get("/path/to/command1")
	if err != nil {
		panic(err)
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := script.NewScript(client, *opt)
	if err := cmd.Run([]string{"-c", "echo 1"}, docker.RunOption{}); err != nil {
		panic(err)
	}
	// Output:
	// 1
}

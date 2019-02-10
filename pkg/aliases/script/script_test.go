package script_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/aliases"
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

	ctx, err := aliases.NewContext("", "")
	if err != nil {
		panic(err)
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := script.NewScript(ctx, client, *opt)
	if err := cmd.Run(ctx, []string{"-c", "echo 1"}, docker.RunOption{}); err != nil {
		panic(err)
	}
	// Output:
	// 1
}

func ExampleScript_Path() {
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

	ctx, err := aliases.NewContext("", "")
	if err != nil {
		panic(err)
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := script.NewScript(ctx, client, *opt)
	fmt.Println(cmd.Path("/tmp"))
	// Output: /tmp/command1
}

func ExampleScript_FileName() {
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

	ctx, err := aliases.NewContext("", "")
	if err != nil {
		panic(err)
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := script.NewScript(ctx, client, *opt)
	fmt.Println(cmd.FileName())
	// Output: command1
}

func ExampleScript_String() {
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

	ctx, err := aliases.NewContext("", "")
	if err != nil {
		panic(err)
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := script.NewScript(ctx, client, *opt)
	fmt.Println(cmd.String())
	// Output: docker run --entrypoint "sh" --interactive --network "host" --rm $(test "$(if tty >/dev/null; then echo true; else echo false; fi)" = "true" && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c
}

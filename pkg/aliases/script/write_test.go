package script_test

import (
	"fmt"
	"io/ioutil"

	"github.com/k-kinzal/aliases/pkg/aliases"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleScript_Write() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
/path/to/command2:
  image: alpine
  tag: latest
/path/to/command3:
  image: alpine
  tag: latest
`
	conf, err := config.Unmarshal([]byte(content))
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

	for _, opt := range conf.Slice() {
		cmd := script.NewScript(ctx, client, opt)
		if _, err := cmd.Write(ctx); err != nil {
			panic(err)
		}
		if _, err := cmd.Write(ctx); err != nil {
			panic(err)
		}
	}

	files, _ := ioutil.ReadDir(ctx.ExportPath())
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Output:
	// command1
	// command2
	// command3
}

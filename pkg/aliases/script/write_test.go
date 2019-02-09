package script_test

import (
	"fmt"
	"io/ioutil"

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

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	path, _ := ioutil.TempDir("/tmp", "")
	for _, opt := range conf.Slice() {
		cmd := script.NewScript(client, opt)
		if err := cmd.Write(path); err != nil {
			panic(err)
		}
		if err := cmd.Write(path); err != nil {
			panic(err)
		}
	}

	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Output:
	// command1
	// command2
	// command3
}

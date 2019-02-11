package script_test

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/k-kinzal/aliases/pkg/aliases/context"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func readDir(path string) []string {
	files := make([]string, 0)
	fileInfo, _ := ioutil.ReadDir(path)
	for _, file := range fileInfo {
		if file.IsDir() {
			continue
		}
		files = append(files, file.Name())
	}

	sort.Strings(files)

	return files
}

func ExampleScript_Write() {
	content := `
/path/to/command1:
  image: alpine
  tag: latest
/path/to/command2:
  image: alpine
  tag: latest
  dependencies:
  - /path/to/command3:
      image: alpine
      tag: latest
`
	conf, err := config.Unmarshal([]byte(content))
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath(dir); err != nil {
		panic(err)
	}

	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	for _, opt := range conf.Slice() {
		cmd := script.NewScript(client, opt)
		if _, err := cmd.Write(); err != nil {
			panic(err)
		}
		if _, err := cmd.Write(); err != nil {
			panic(err)
		}
	}

	for _, file := range readDir(context.ExportPath()) {
		fmt.Println(file)
	}

	// Output:
	// command1
	// command2
}

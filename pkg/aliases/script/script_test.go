package script_test

import (
	"fmt"
	"io/ioutil"

	"github.com/k-kinzal/aliases/pkg/aliases/config"
	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleNewScript() {
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

func ExampleScript_Path() {
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
	fmt.Println(cmd.Path("/tmp"))
	// Output: /tmp/command1
}

func ExampleScript_FileName() {
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
	fmt.Println(cmd.FileName())
	// Output: command1
}

func ExampleScript_StringWithOverride() {
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
	fmt.Println(cmd.StringWithOverride([]string{"-c", "echo 1"}, docker.RunOption{Env: map[string]string{"FOO": "1"}}))
	// Output: docker run --entrypoint "sh" --env FOO="1" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "echo 1"
}

func ExampleScript_String() {
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
	fmt.Println(cmd.String())
	// Output: docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c
}

func ExampleScript_Shell() {
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
	shell, err := cmd.Shell([]string{}, docker.RunOption{})
	if err != nil {
		panic(err)
	}
	fmt.Println(shell.String())
	// Output:
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "$@"
	//   exit $?
	// fi
}

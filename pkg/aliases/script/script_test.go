package script_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleScript_Write() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleScript_Write"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleScript_Write/export"); err != nil {
		panic(err)
	}
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image: "alpine",
			Tag:   "latest",
			Args:  []string{"-c"},
		},
		Path: yaml.SpecPath("/path/to/command"),
	}

	if err := script.NewScript(spec).Write(client); err != nil {
		panic(err)
	}

	fileinfo, _ := ioutil.ReadDir(context.ExportPath())
	for _, file := range fileinfo {
		content, _ := ioutil.ReadFile(filepath.Join(context.ExportPath(), file.Name()))
		fmt.Printf("%s:\n", file.Name())
		fmt.Println(string(content))
	}
	// Output:
	// command:
	// #!/bin/sh
	//
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run alpine:${COMMAND_VERSION:-"latest"} -c "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run alpine:${COMMAND_VERSION:-"latest"} -c "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run alpine:${COMMAND_VERSION:-"latest"} -c "$@"
	//   exit $?
	// fi
}

func ExampleScript_WriteWithOverride() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleScript_WriteWithOverride"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleScript_WriteWithOverride/export"); err != nil {
		panic(err)
	}
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{Image: "docker", Tag: "18.09.0"},
			Image: "alpine",
			Tag:   "latest",
			Args:  []string{"-c"},
		},
		Path: yaml.SpecPath("/path/to/command1"),
		Dependencies: []*yaml.Option{
			{
				OptionSpec: &yaml.OptionSpec{
					Image: "alpine",
					Tag:   "latest",
					Args:  []string{"-c"},
				},
				Path: yaml.SpecPath("/path/to/command2"),
			},
		},
	}

	entrypoint := "bash"
	if err := script.NewScript(spec).WriteWithOverride(client, []string{"--posix", "-c"}, docker.RunOption{Entrypoint: &entrypoint}); err != nil {
		panic(err)
	}

	fileinfo, _ := ioutil.ReadDir(context.ExportPath())
	for _, file := range fileinfo {
		content, _ := ioutil.ReadFile(filepath.Join(context.ExportPath(), file.Name()))
		fmt.Printf("%s:\n", file.Name())
		fmt.Println(string(content))
	}
	// Output:
	// command1:
	// #!/bin/sh
	//
	// DOCKER_BINARY_PATH="/tmp/aliases/ExampleScript_WriteWithOverride/docker/docker-18-09-0"
	// if [ ! -f "${DOCKER_BINARY_PATH}" ]; then
	//   docker run --entrypoint "" --volume "/tmp/aliases/ExampleScript_WriteWithOverride/docker:/share" docker:18.09.0 sh -c "cp -av $(which docker) /share/docker-18-09-0" >/dev/null
	// fi
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run --entrypoint "bash" --env ALIASES_PWD="${ALIASES_PWD:-$PWD}" --privileged --volume "/tmp/aliases/ExampleScript_WriteWithOverride:/tmp/aliases/ExampleScript_WriteWithOverride" --volume "${DOCKER_BINARY_PATH}:/usr/local/bin/docker" --volume "/var/run/docker.sock:/var/run/docker.sock" --volume "/tmp/aliases/ExampleScript_WriteWithOverride/export/command2:/path/to/command2" alpine:${COMMAND1_VERSION:-"latest"} --posix -c "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run --entrypoint "bash" --env ALIASES_PWD="${ALIASES_PWD:-$PWD}" --privileged --volume "/tmp/aliases/ExampleScript_WriteWithOverride:/tmp/aliases/ExampleScript_WriteWithOverride" --volume "${DOCKER_BINARY_PATH}:/usr/local/bin/docker" --volume "/var/run/docker.sock:/var/run/docker.sock" --volume "/tmp/aliases/ExampleScript_WriteWithOverride/export/command2:/path/to/command2" alpine:${COMMAND1_VERSION:-"latest"} --posix -c "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run --entrypoint "bash" --env ALIASES_PWD="${ALIASES_PWD:-$PWD}" --privileged --volume "/tmp/aliases/ExampleScript_WriteWithOverride:/tmp/aliases/ExampleScript_WriteWithOverride" --volume "${DOCKER_BINARY_PATH}:/usr/local/bin/docker" --volume "/var/run/docker.sock:/var/run/docker.sock" --volume "/tmp/aliases/ExampleScript_WriteWithOverride/export/command2:/path/to/command2" alpine:${COMMAND1_VERSION:-"latest"} --posix -c "$@"
	//   exit $?
	// fi
	//
	// command2:
	// #!/bin/sh
	//
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run alpine:${COMMAND2_VERSION:-"latest"} -c "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run alpine:${COMMAND2_VERSION:-"latest"} -c "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run alpine:${COMMAND2_VERSION:-"latest"} -c "$@"
	//   exit $?
	// fi
}

func ExampleScript_Alias() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleScript_Alias"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleScript_Alias/export"); err != nil {
		panic(err)
	}
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image: "alpine",
			Tag:   "latest",
			Args:  []string{"-c"},
		},
		Path: yaml.SpecPath("/path/to/command"),
	}

	cmd, err := script.NewScript(spec).Alias(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.String())
	// Output: alias command='/tmp/aliases/ExampleScript_Alias/export/command'
}

func ExampleScript_Run() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleScript_Alias"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleScript_Alias/export"); err != nil {
		panic(err)
	}
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image:   "alpine",
			Tag:     "latest",
			Command: (func(s string) *string { return &s })("sh"),
			Args:    []string{"-c"},
		},
		Path: yaml.SpecPath("/path/to/command"),
	}

	if err := script.NewScript(spec).Run(client, []string{"-c", "echo 1"}, docker.RunOption{}); err != nil {
		panic(err)
	}
	// Output: 1
}

//import (
//	"fmt"
//	"io/ioutil"
//
//	"github.com/k-kinzal/aliases/pkg/aliases/context"
//	"github.com/k-kinzal/aliases/pkg/aliases/script"
//	"github.com/k-kinzal/aliases/pkg/docker"
//)
//
//func ExampleNewScript() {
//	dir, err := ioutil.TempDir("/tmp", "")
//	if err != nil {
//		panic(err)
//	}
//	if err := context.ChangeHomePath(dir); err != nil {
//		panic(err)
//	}
//	if err := context.ChangeExportPath(dir); err != nil {
//		panic(err)
//	}
//
//	content := `
///path/to/command1:
//  image: alpine
//  tag: latest
//  entrypoint: sh
//  args: [-c]
//`
//	conf, err := config.Unmarshal([]byte(content))
//	if err != nil {
//		panic(err)
//	}
//
//	opt, err := conf.Get("/path/to/command1")
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := docker.NewClient()
//	if err != nil {
//		panic(err)
//	}
//
//	cmd := script.NewScript(client, *opt)
//	if err := cmd.Run([]string{"-c", "echo 1"}, docker.RunOption{}); err != nil {
//		panic(err)
//	}
//	// Output:
//	// 1
//}
//
//func ExampleScript_Path() {
//	dir, err := ioutil.TempDir("/tmp", "")
//	if err != nil {
//		panic(err)
//	}
//	if err := context.ChangeHomePath(dir); err != nil {
//		panic(err)
//	}
//	if err := context.ChangeExportPath(dir); err != nil {
//		panic(err)
//	}
//
//	content := `
///path/to/command1:
//  image: alpine
//  tag: latest
//  entrypoint: sh
//  args: [-c]
//`
//	conf, err := config.Unmarshal([]byte(content))
//	if err != nil {
//		panic(err)
//	}
//
//	opt, err := conf.Get("/path/to/command1")
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := docker.NewClient()
//	if err != nil {
//		panic(err)
//	}
//
//	cmd := script.NewScript(client, *opt)
//	fmt.Println(cmd.Path("/tmp"))
//	// Output: /tmp/command1
//}
//
//func ExampleScript_FileName() {
//	dir, err := ioutil.TempDir("/tmp", "")
//	if err != nil {
//		panic(err)
//	}
//	if err := context.ChangeHomePath(dir); err != nil {
//		panic(err)
//	}
//	if err := context.ChangeExportPath(dir); err != nil {
//		panic(err)
//	}
//
//	content := `
///path/to/command1:
//  image: alpine
//  tag: latest
//  entrypoint: sh
//  args: [-c]
//`
//	conf, err := config.Unmarshal([]byte(content))
//	if err != nil {
//		panic(err)
//	}
//
//	opt, err := conf.Get("/path/to/command1")
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := docker.NewClient()
//	if err != nil {
//		panic(err)
//	}
//
//	cmd := script.NewScript(client, *opt)
//	fmt.Println(cmd.FileName())
//	// Output: command1
//}
//
//func ExampleScript_StringWithOverride() {
//	dir, err := ioutil.TempDir("/tmp", "")
//	if err != nil {
//		panic(err)
//	}
//	if err := context.ChangeHomePath(dir); err != nil {
//		panic(err)
//	}
//	if err := context.ChangeExportPath(dir); err != nil {
//		panic(err)
//	}
//
//	content := `
///path/to/command1:
//  image: alpine
//  tag: latest
//  entrypoint: sh
//  args: [-c]
//`
//	conf, err := config.Unmarshal([]byte(content))
//	if err != nil {
//		panic(err)
//	}
//
//	opt, err := conf.Get("/path/to/command1")
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := docker.NewClient()
//	if err != nil {
//		panic(err)
//	}
//
//	cmd := script.NewScript(client, *opt)
//	fmt.Println(cmd.StringWithOverride([]string{"-c", "echo 1"}, docker.RunOption{Env: map[string]string{"FOO": "1"}}))
//	// Output: docker run --entrypoint "sh" --env FOO="1" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "echo 1"
//}
//
//func ExampleScript_String() {
//	dir, err := ioutil.TempDir("/tmp", "")
//	if err != nil {
//		panic(err)
//	}
//	if err := context.ChangeHomePath(dir); err != nil {
//		panic(err)
//	}
//	if err := context.ChangeExportPath(dir); err != nil {
//		panic(err)
//	}
//
//	content := `
///path/to/command1:
//  image: alpine
//  tag: latest
//  entrypoint: sh
//  args: [-c]
//`
//	conf, err := config.Unmarshal([]byte(content))
//	if err != nil {
//		panic(err)
//	}
//
//	opt, err := conf.Get("/path/to/command1")
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := docker.NewClient()
//	if err != nil {
//		panic(err)
//	}
//
//	cmd := script.NewScript(client, *opt)
//	fmt.Println(cmd.String())
//	// Output: docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c
//}
//
//func ExampleScript_Shell() {
//	dir, err := ioutil.TempDir("/tmp", "")
//	if err != nil {
//		panic(err)
//	}
//	if err := context.ChangeHomePath(dir); err != nil {
//		panic(err)
//	}
//	if err := context.ChangeExportPath(dir); err != nil {
//		panic(err)
//	}
//
//	content := `
///path/to/command1:
//  image: alpine
//  tag: latest
//  entrypoint: sh
//  args: [-c]
//`
//	conf, err := config.Unmarshal([]byte(content))
//	if err != nil {
//		panic(err)
//	}
//
//	opt, err := conf.Get("/path/to/command1")
//	if err != nil {
//		panic(err)
//	}
//
//	client, err := docker.NewClient()
//	if err != nil {
//		panic(err)
//	}
//
//	cmd := script.NewScript(client, *opt)
//	shell, err := cmd.Shell([]string{}, docker.RunOption{})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(shell.String())
//	// Output:
//	// if [ -p /dev/stdin ]; then
//	//   cat - | docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "$@"
//	//   exit $?
//	// elif [ -f /dev/stdin ]; then
//	//   docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "$@" </dev/stdin
//	//   exit $?
//	// else
//	//   echo "" >/dev/null | docker run --entrypoint "sh" --interactive --network "host" --rm $(tty >/dev/null && echo "--tty") alpine:${COMMAND1_VERSION:-"latest"} -c "$@"
//	//   exit $?
//	// fi
//}
//
////func ExampleScript_Run() {
////	content := `
/////path/to/command1:
////  image: alpine
////  tag: latest
////  entrypoint: sh
////`
////	conf, err := config.Unmarshal([]byte(content))
////	if err != nil {
////		panic(err)
////	}
////
////	opt, err := conf.Get("/path/to/command1")
////	if err != nil {
////		panic(err)
////	}
////
////	dir, err := ioutil.TempDir("/tmp", "")
////	if err != nil {
////		panic(err)
////	}
////	if err := context.ChangeHomePath(dir); err != nil {
////		panic(err)
////	}
////
////	client, err := docker.NewClient()
////	if err != nil {
////		panic(err)
////	}
////
////	cmd := script.NewScript(client, *opt)
////	if err := cmd.Run([]string{"-c", "echo 1"}, docker.RunOption{}); err != nil {
////		panic(err)
////	}
////	// Output:
////	// 1
////}
//
////func readDir(path string) []string {
////	files := make([]string, 0)
////	fileInfo, _ := ioutil.ReadDir(path)
////	for _, file := range fileInfo {
////		if file.IsDir() {
////			continue
////		}
////		files = append(files, file.Name())
////	}
////
////	sort.Strings(files)
////
////	return files
////}
////
////func ExampleScript_Write() {
////	content := `
/////path/to/command1:
////  image: alpine
////  tag: latest
/////path/to/command2:
////  image: alpine
////  tag: latest
////  dependencies:
////  - /path/to/command3:
////      image: alpine
////      tag: latest
////`
////	conf, err := config.Unmarshal([]byte(content))
////	if err != nil {
////		panic(err)
////	}
////
////	dir, err := ioutil.TempDir("/tmp", "")
////	if err != nil {
////		panic(err)
////	}
////	if err := context.ChangeExportPath(dir); err != nil {
////		panic(err)
////	}
////
////	client, err := docker.NewClient()
////	if err != nil {
////		panic(err)
////	}
////
////	for _, opt := range conf.Slice() {
////		cmd := script.NewScript(client, opt)
////		if _, err := cmd.Write(); err != nil {
////			panic(err)
////		}
////		if _, err := cmd.Write(); err != nil {
////			panic(err)
////		}
////	}
////
////	for _, file := range readDir(context.ExportPath()) {
////		fmt.Println(file)
////	}
////
////	// Output:
////	// command1
////	// command2
////}

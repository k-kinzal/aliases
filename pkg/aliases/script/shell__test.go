package script_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleShellAdapter_Command() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleShellAdapter_Command"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleShellAdapter_Command"); err != nil {
		panic(err)
	}
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image:      "alpine",
			Tag:        "latest",
			Entrypoint: (func(str string) *string { return &str })("sh"),
			Args:       []string{"-c"},
		},
	}
	runner := script.AdaptShell(spec)
	cmd, err := runner.Command(client, []string{"-c", "echo 1"}, docker.RunOption{
		Entrypoint: (func(str string) *string { return &str })("bash"),
	}, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.String())
	// Output:
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run --entrypoint "bash" alpine:${._VERSION:-"latest"} -c "echo 1" "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run --entrypoint "bash" alpine:${._VERSION:-"latest"} -c "echo 1" "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run --entrypoint "bash" alpine:${._VERSION:-"latest"} -c "echo 1" "$@"
	//   exit $?
	// fi
}

func ExampleShellAdapter_Command_UseExtendEntrypoint() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleShellAdapter_Command_UseExtendEntrypoint"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleShellAdapter_Command_UseExtendEntrypoint"); err != nil {
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
			Entrypoint: (func(str string) *string { return &str })(`
				#!/bin/sh
				sh -c "$@"
			`),
			Args: []string{"-c"},
		},
	}
	runner := script.AdaptShell(spec)
	cmd, err := runner.Command(client, []string{}, docker.RunOption{}, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.String())
	// Output:
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run --entrypoint "/b82c6c95fa559f017ac62344f961e206" --volume "/tmp/aliases/ExampleShellAdapter_Command_UseExtendEntrypoint/entrypoint/b82c6c95fa559f017ac62344f961e206:/b82c6c95fa559f017ac62344f961e206" alpine:${._VERSION:-"latest"} -c "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run --entrypoint "/b82c6c95fa559f017ac62344f961e206" --volume "/tmp/aliases/ExampleShellAdapter_Command_UseExtendEntrypoint/entrypoint/b82c6c95fa559f017ac62344f961e206:/b82c6c95fa559f017ac62344f961e206" alpine:${._VERSION:-"latest"} -c "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run --entrypoint "/b82c6c95fa559f017ac62344f961e206" --volume "/tmp/aliases/ExampleShellAdapter_Command_UseExtendEntrypoint/entrypoint/b82c6c95fa559f017ac62344f961e206:/b82c6c95fa559f017ac62344f961e206" alpine:${._VERSION:-"latest"} -c "$@"
	//   exit $?
	// fi
}

func ExampleShellAdapter_Command_HasDependencies() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleShellAdapter_Command_HasDependencies"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleShellAdapter_Command_HasDependencies"); err != nil {
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
			Image:      "alpine",
			Tag:        "latest",
			Entrypoint: (func(str string) *string { return &str })("sh"),
			Args:       []string{"-c"},
		},
		Dependencies: []*yaml.Option{
			{
				OptionSpec: &yaml.OptionSpec{
					Image: "alpine",
					Tag:   "latest",
				},
			},
		},
	}
	runner := script.AdaptShell(spec)
	cmd, err := runner.Command(client, []string{}, docker.RunOption{}, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.String())
	// Output:
	// DOCKER_BINARY_PATH="/tmp/aliases/ExampleShellAdapter_Command_HasDependencies/docker/docker-18-09-0"
	// if [ ! -f "${DOCKER_BINARY_PATH}" ]; then
	//   docker run --entrypoint "" --volume "/tmp/aliases/ExampleShellAdapter_Command_HasDependencies/docker:/share" docker:18.09.0 sh -c "cp -av $(which docker) /share/docker-18-09-0" >/dev/null
	// fi
	// if [ -p /dev/stdin ]; then
	//   cat - | docker run --entrypoint "sh" --env ALIASES_PWD="${ALIASES_PWD:-$PWD}" --privileged --volume "/tmp/aliases/ExampleShellAdapter_Command_HasDependencies:/tmp/aliases/ExampleShellAdapter_Command_HasDependencies" --volume "${DOCKER_BINARY_PATH}:/usr/local/bin/docker" --volume "/var/run/docker.sock:/var/run/docker.sock" alpine:${._VERSION:-"latest"} -c "$@"
	//   exit $?
	// elif [ -f /dev/stdin ]; then
	//   docker run --entrypoint "sh" --env ALIASES_PWD="${ALIASES_PWD:-$PWD}" --privileged --volume "/tmp/aliases/ExampleShellAdapter_Command_HasDependencies:/tmp/aliases/ExampleShellAdapter_Command_HasDependencies" --volume "${DOCKER_BINARY_PATH}:/usr/local/bin/docker" --volume "/var/run/docker.sock:/var/run/docker.sock" alpine:${._VERSION:-"latest"} -c "$@" </dev/stdin
	//   exit $?
	// else
	//   echo "" >/dev/null | docker run --entrypoint "sh" --env ALIASES_PWD="${ALIASES_PWD:-$PWD}" --privileged --volume "/tmp/aliases/ExampleShellAdapter_Command_HasDependencies:/tmp/aliases/ExampleShellAdapter_Command_HasDependencies" --volume "${DOCKER_BINARY_PATH}:/usr/local/bin/docker" --volume "/var/run/docker.sock:/var/run/docker.sock" alpine:${._VERSION:-"latest"} -c "$@"
	//   exit $?
	// fi
}

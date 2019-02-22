package script_test

import (
	"fmt"
	"testing"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleDockerRunAdapter_Image() {
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image:      "alpine",
			Tag:        "latest",
			Entrypoint: (func(str string) *string { return &str })("sh"),
			Args:       []string{"-c"},
		},
	}
	runner := script.AdaptDockerRun(spec)
	fmt.Println(runner.Image())
	// Output: alpine:${._VERSION:-"latest"}
}

func ExampleDockerRunAdapter_Args() {
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image:      "alpine",
			Tag:        "latest",
			Entrypoint: (func(str string) *string { return &str })("sh"),
			Args:       []string{"-c"},
		},
	}
	runner := script.AdaptDockerRun(spec)
	fmt.Println(runner.Args())
	// Output: [-c]
}

func ExampleDockerRunAdapter_Option() {
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Image:      "alpine",
			Tag:        "latest",
			Entrypoint: (func(str string) *string { return &str })("sh"),
			Args:       []string{"-c"},
		},
	}
	runner := script.AdaptDockerRun(spec)
	fmt.Println(*runner.Option().Entrypoint)
	// Output: sh
}

func ExampleDockerRunAdapter_Command() {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleDockerRunAdapter_Command"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleDockerRunAdapter_Command"); err != nil {
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
	runner := script.AdaptDockerRun(spec)
	cmd := runner.Command(client, []string{"-c", "echo 1"}, docker.RunOption{
		Entrypoint: (func(str string) *string { return &str })("bash"),
	})
	fmt.Println(cmd.String())
	// Output: docker run --entrypoint "bash" alpine:${._VERSION:-"latest"} -c "echo 1"
}

func TestDockerRunAdapter_Command(t *testing.T) {
	if err := context.ChangeHomePath("/tmp/aliases/ExampleDockerBinaryAdapter_Command"); err != nil {
		panic(err)
	}
	if err := context.ChangeExportPath("/tmp/aliases/ExampleDockerBinaryAdapter_Command"); err != nil {
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
	runner := script.AdaptDockerRun(spec)
	cmd := runner.Command(client, []string{}, docker.RunOption{})

	if cmd.String() != "docker run --entrypoint \"sh\" alpine:${._VERSION:-\"latest\"} -c" {
		t.Errorf("expected `docker run --entrypoint \"sh\" alpine:${._VERSION:-\"latest\"} -c`, but actual is `%s`", cmd.String())
	}
}

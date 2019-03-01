package script_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/aliases/context"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/aliases/script"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
)

func ExampleDockerBinaryAdapter_Image() {
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{Image: "docker", Tag: "18.09.0"},
		},
	}
	bin := script.AdaptDockerBinary(spec)
	fmt.Println(bin.Image())
	// Output: docker
}

func ExampleDockerBinaryAdapter_Tag() {
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{Image: "docker", Tag: "18.09.0"},
		},
	}
	bin := script.AdaptDockerBinary(spec)
	fmt.Println(bin.Tag())
	// Output: 18.09.0
}

func ExampleDockerBinaryAdapter_FileName() {
	spec := yaml.Option{
		OptionSpec: &yaml.OptionSpec{
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{Image: "docker", Tag: "18.09.0"},
		},
	}
	bin := script.AdaptDockerBinary(spec)
	fmt.Println(bin.FileName())
	// Output: docker-18-09-0
}

func ExampleDockerBinaryAdapter_Command() {
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
			Docker: struct {
				Image string `yaml:"image" default:"docker"`
				Tag   string `yaml:"tag" default:"18.09.0"`
			}{Image: "docker", Tag: "18.09.0"},
		},
	}
	bin := script.AdaptDockerBinary(spec)
	fmt.Println(bin.Command(client).String())
	// Output: docker run --entrypoint "" --volume "/tmp/aliases/ExampleDockerBinaryAdapter_Command/docker:/share" docker:18.09.0 sh -c 'cp -av $(which docker) /share/docker-18-09-0'
}

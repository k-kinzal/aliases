package script

import (
	"fmt"
	"strings"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
	"github.com/k-kinzal/aliases/pkg/posix"
)

// DockerBinaryAdapter adapts binary information from spec to a form that can be used in aliases.
type DockerBinaryAdapter yaml.Option

// Image returns image of docker binary.
func (adpt *DockerBinaryAdapter) Image() string {
	spec := (*yaml.Option)(adpt)
	return spec.Docker.Image
}

// Tag returns tag of docker binary.
func (adpt *DockerBinaryAdapter) Tag() string {
	spec := (*yaml.Option)(adpt)
	return spec.Docker.Tag
}

// FileName returns filename of docker binary.
func (adpt *DockerBinaryAdapter) FileName() string {
	filename := fmt.Sprintf("%s:%s", adpt.Image(), adpt.Tag())
	filename = strings.Replace(filename, "/", "-", -1)
	filename = strings.Replace(filename, ":", "-", -1)
	filename = strings.Replace(filename, ".", "-", -1)
	filename = strings.Replace(filename, "_", "-", -1)

	return filename
}

// Command returns a command to download the docker binary.
func (adpt *DockerBinaryAdapter) Command(client *docker.Client) *posix.Cmd {
	image := fmt.Sprintf("%s:%s", adpt.Image(), adpt.Tag())

	cmd := client.Run(image, []string{
		"sh",
		"-c",
		fmt.Sprintf("'cp -av $(which docker) /share/%s'", adpt.FileName()),
	}, docker.RunOption{
		Entrypoint: (func(str string) *string { return &str })(""),
		Volume: []string{
			fmt.Sprintf("%s:/share", context.BinaryPath()),
		},
	})

	return cmd
}

// adaptDockerBinary returns DockerBinaryAdapter.
func adaptDockerBinary(spec yaml.Option) *DockerBinaryAdapter {
	return (*DockerBinaryAdapter)(&spec)
}

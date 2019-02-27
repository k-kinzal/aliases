package script

import "github.com/k-kinzal/aliases/pkg/aliases/yaml"

func AdaptDockerRun(spec yaml.Option) *DockerRunAdapter {
	return adaptDockerRun(spec)
}

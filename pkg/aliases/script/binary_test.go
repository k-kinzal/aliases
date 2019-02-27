package script

import "github.com/k-kinzal/aliases/pkg/aliases/yaml"

func AdaptDockerBinary(spec yaml.Option) *DockerBinaryAdapter {
	return adaptDockerBinary(spec)
}

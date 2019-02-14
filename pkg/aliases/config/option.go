package config

import (
	"fmt"
	"path"
	"strings"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/types"
)

// Option is configuration option.
type Option struct {
	*yaml.OptionSpec
	Namespace    string
	Path         string
	FileName     string
	Dependencies []*Option
}

// Binary returns docker binary info.
func (opt Option) Binary(binaryDir string) *DockerBinary {
	filename := fmt.Sprintf("%s:%s", opt.Docker.Image, opt.Docker.Tag)
	filename = strings.Replace(filename, "/", "-", -1)
	filename = strings.Replace(filename, ":", "-", -1)
	filename = strings.Replace(filename, ".", "-", -1)
	filename = strings.Replace(filename, "_", "-", -1)

	return &DockerBinary{
		Image: opt.Docker.Image,
		Tag:   opt.Docker.Tag,
		Path:  path.Join(binaryDir, filename),
	}
}

// Binaries returns slice of docker binery info.
func (opt *Option) Binaries(binaryDir string) []DockerBinary {
	set := types.NewSet(nil)
	for _, dep := range opt.Dependencies {
		for _, binary := range dep.Binaries(binaryDir) {
			set.Add(binary)
		}
	}

	set.Add(*opt.Binary(binaryDir))

	slice := set.Slice()
	binaries := make([]DockerBinary, len(slice))
	for i := 0; i < len(slice); i++ {
		binaries[i] = slice[i].(DockerBinary)
	}

	return binaries
}

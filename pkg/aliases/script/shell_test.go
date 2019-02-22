package script

import "github.com/k-kinzal/aliases/pkg/aliases/yaml"

func AdaptShell(spec yaml.Option) *ShellAdapter {
	return adaptShell(spec)
}

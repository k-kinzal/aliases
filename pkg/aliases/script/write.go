package script

import (
	"os"
	"path"
	"text/template"

	"github.com/k-kinzal/aliases/pkg/docker"
)

var content = `#!/bin/sh
if [ -p /dev/stdin ]; then
  cat - | {{ .DockerRunCommand }} "$@"
  exit $?
else
  {{ .DockerRunCommand }} "$@"
  exit $?
fi
`

// Write exports aliases script to a file.
func (script *Script) Write(dir string) error {
	return script.WriteWithOverride(dir, nil, docker.RunOption{})
}

// Write exports aliases script to a file with override docker option.
func (script *Script) WriteWithOverride(dir string, args []string, option docker.RunOption) error {
	for _, cmd := range script.relative {
		if err := cmd.Write(dir); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(path.Dir(script.Path(dir)), 0755); err != nil {
		return err
	}

	file, err := os.Create(script.Path(dir))
	if err != nil {
		return err
	}
	defer file.Close()

	if err := os.Chmod(script.Path(dir), 0755); err != nil {
		return err
	}

	tmpl := template.Must(template.New(script.path).Parse(content))

	data := map[string]interface{}{
		"DockerRunCommand": script.docker(args, option).String(),
	}

	if err := tmpl.Execute(file, data); err != nil {
		return err
	}

	return nil
}

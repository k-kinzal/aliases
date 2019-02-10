package script

import (
	"os"
	"path"
	"text/template"

	"github.com/k-kinzal/aliases/pkg/aliases"

	"github.com/k-kinzal/aliases/pkg/docker"
)

// FIXME:
//  `-p /dev/stdin` always becomes true in `docker run -i`.
//  Therefore, we use timeouts at the expense of performance.
var content = `#!/bin/sh
if [ -p /dev/stdin ]; then
  while read -t 1 line; do
    echo $line
  done | {{ .command }} "$@"
  exit $?
else
  {{ .command }} "$@"
  exit $?
fi
`

// Write exports aliases script to a file.
func (script *Script) Write(ctx aliases.Context) (string, error) {
	return script.WriteWithOverride(ctx, nil, docker.RunOption{})
}

// Write exports aliases script to a file with override docker option.
func (script *Script) WriteWithOverride(ctx aliases.Context, args []string, option docker.RunOption) (string, error) {
	for _, cmd := range script.relative {
		if _, err := cmd.Write(ctx); err != nil {
			return "", err
		}
	}

	targetPath := script.Path(ctx.ExportPath())

	if err := os.MkdirAll(path.Dir(targetPath), 0755); err != nil {
		return "", err
	}

	file, err := os.Create(targetPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := os.Chmod(targetPath, 0755); err != nil {
		return "", err
	}

	data := map[string]interface{}{
		"command": script.docker(args, option).String(),
	}

	tmpl := template.Must(template.New(script.path).Parse(content))

	if err := tmpl.Execute(file, data); err != nil {
		return "", err
	}

	return targetPath, nil
}

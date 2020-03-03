package script

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"

	"github.com/k-kinzal/aliases/pkg/types"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
	"github.com/k-kinzal/aliases/pkg/posix"
)

var content = `
{{- if .binary }}
DOCKER_BINARY_PATH="{{ .binary.path }}/{{ .binary.filename }}"
if [ ! -f "${DOCKER_BINARY_PATH}" ]; then
  {{ .binary.command }} >/dev/null
fi
{{- end }}
{{- if .envPrefix }}
TEMPDIR="{{ .temporaryPath }}"
{{- range $prefix := .envPrefix }}
touch ${TEMPDIR}/{{ $prefix | lower | trimSuffix "_" }}.env
for line in $(env | grep "^{{ $prefix }}"); do
  echo "${line#{{ $prefix }}}" >> ${TEMPDIR}/{{ $prefix | lower | trimSuffix "_" }}.env
done
{{- end }}
{{- end }}
{{- if .debug }}
echo "\033[0;90m{{ .command | replace "\\n" "\\\\n" | replace "\r" "\\r" | quote }}\033[0m"
{{- end }}
if [ -p /dev/stdin ]; then
  cat - | {{ .command }} "$@"
  exit $?
elif [ -f /dev/stdin ]; then
  {{ .command }} "$@" </dev/stdin
  exit $?
elif [ -t 0 ]; then
  {{ .command }} "$@"
  exit $?
else
  echo "" >/dev/null | {{ .command }} "$@"
  exit $?
fi
`

// ShellAdapter adapts shell script from spec to a form that can be used in aliases.
type ShellAdapter yaml.Option

// Command returns a command to aliases script.
func (adpt *ShellAdapter) Command(client *docker.Client, overrideArgs []string, overrideOption docker.RunOption, debug bool) (*posix.ShellScript, error) {
	spec := (*yaml.Option)(adpt)
	bin := adaptDockerBinary(*spec)
	runner := adaptDockerRun(*spec)

	// extend entrypoint
	if overrideOption.Entrypoint == nil && spec.Entrypoint != nil && strings.HasPrefix(strings.Trim(*spec.Entrypoint, " \t\r\n"), "#!") {
		body := strings.Trim(*spec.Entrypoint, " \t\r\n")
		hash := types.MD5(body)
		entrypoint := fmt.Sprintf("/%s", hash)

		path := filepath.Join(context.ExportPath(), "entrypoint", hash)

		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(path, []byte(body), 0755); err != nil {
			return nil, err
		}

		overrideOption.Entrypoint = &entrypoint
		overrideOption.Volume = append(overrideOption.Volume, fmt.Sprintf("%s:%s", path, entrypoint))
	}
	// docker info
	if spec.Docker != nil {
		if overrideOption.Env == nil {
			overrideOption.Env = make(map[string]string)
		}
		overrideOption.Env["ALIASES_PWD"] = "${ALIASES_PWD:-$PWD}"
		if sock := client.Sock(); sock != nil {
			// unix socket
			overrideOption.Privileged = (func(str string) *string { return &str })("true")
			overrideOption.Volume = append(overrideOption.Volume, fmt.Sprintf("%s:%s", context.HomePath(), context.HomePath()))
			overrideOption.Volume = append(overrideOption.Volume, "${DOCKER_BINARY_PATH}:/usr/local/bin/docker")
			overrideOption.Volume = append(overrideOption.Volume, fmt.Sprintf("%s:/var/run/docker.sock", *sock))
		} else {
			// tcp, http...
			overrideOption.Network = (func(str string) *string { return &str })("host")
			overrideOption.Env["DOCKER_HOST"] = client.Host()
		}
	}
	//
	for _, prefix := range spec.EnvPrefix {
		overrideOption.EnvFile = append(overrideOption.EnvFile, "${TEMPDIR}/"+strings.Trim(strings.ToLower(prefix), "_")+".env")
	}
	// resolve dependency commands
	for _, dependency := range spec.Dependencies {
		path := filepath.Join(context.ExportPath(), dependency.Namespace(), dependency.Path.Base())
		if spec.Path == dependency.Path {
			// command1:
			//   ...
			//   dependencies:
			//   - command1
			continue
		}

		overrideOption.Volume = append(overrideOption.Volume, fmt.Sprintf("%s:%s", path, dependency.Path.Name()))

		if _, err := os.Stat(path); !os.IsNotExist(err) {
			// command1:
			//   ...
			//   dependencies:
			//   - command2:
			//       ...
			//       dependencies:
			//       - command1
			continue
		}

		if err := NewScript(*dependency).Write(client); err != nil {
			return nil, err
		}
	}

	funcs := template.FuncMap{
		"quote": strconv.Quote,
		"replace": func(old string, new string, s string) string {
			return strings.Replace(s, old, new, -1)
		},
	}

	tmpl := template.Must(template.New(adpt.Path.Name()).Funcs(sprig.TxtFuncMap()).Funcs(funcs).Parse(content))

	data := map[string]interface{}{
		"command":       runner.Command(client, overrideArgs, overrideOption).String(),
		"envPrefix":     spec.EnvPrefix,
		"dependencies":  len(spec.Dependencies) > 0,
		"binary":        nil,
		"debug":         debug,
		"temporaryPath": context.TemporaryPath(spec),
	}
	if spec.Docker != nil {
		data["binary"] = map[string]string{
			"command":  bin.Command(client).String(),
			"filename": bin.FileName(),
			"path":     context.BinaryPath(),
		}
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return nil, err
	}

	return posix.Shell(tpl.String()), nil
}

// adaptShell returns ShellAdapter.
func adaptShell(spec yaml.Option) *ShellAdapter {
	return (*ShellAdapter)(&spec)
}

package docker

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/k-kinzal/aliases/pkg/posix"
)

// VersionOption is an option for executing docker version.
type VersionOption struct {
	Format     *string
	Kubeconfig *string
}

// Version returns a command that can execute docker version.
func (client *Client) Version(option VersionOption) *posix.Cmd {
	cmd := posix.Command(client.path, "version")
	if v := option.Format; v != nil {
		cmd.Args = append(cmd.Args, "--format", strconv.Quote(*v))
	}
	if v := option.Kubeconfig; v != nil {
		cmd.Args = append(cmd.Args, "--kubeconfig", strconv.Quote(*v))
	}

	return cmd
}

// ClientVersion is client info.
//
// Client: Docker Engine - Community
//  Version:           18.09.1
//  API version:       1.39
//  Go version:        go1.10.6
//  Git commit:        4c52b90
//  Built:             Wed Jan  9 19:33:12 2019
//  OS/Arch:           darwin/amd64
//  Experimental:      false
type ClientVersion struct {
	Version    string
	APIVersion string
	GOVersion  string
	GitCommit  string
	Build      string
	OS         string
	Arch       string
}

// ClientVersion returns client info.
func (client *Client) ClientVersion() (*ClientVersion, error) {
	format := strings.Trim(`
{{ .Client.Version }}
{{ .Client.APIVersion }}
{{ .Client.GoVersion }}
{{ .Client.GitCommit }}
{{ .Client.BuildTime }}
{{ .Client.Os }}
{{ .Client.Arch }}
`, " \t\r\n")
	opt := VersionOption{
		Format: &format,
	}
	cmd := client.Version(opt)
	out, err := cmd.Output()
	if err != nil {
		e, ok := err.(*exec.ExitError)
		if !ok {
			return nil, err
		}
		if !strings.HasPrefix(string(e.Stderr), "Cannot connect to the Docker daemon") {
			return nil, err
		}
	}
	s := strings.Split(strings.Trim(string(out), "\"\n"), "\\n")

	return &ClientVersion{
		Version:    s[0],
		APIVersion: s[1],
		GOVersion:  s[2],
		GitCommit:  s[3],
		Build:      s[4],
		OS:         s[5],
		Arch:       s[6],
	}, nil
}

// ServerVersion is client info.
//
// Server: Docker Engine - Community
//  Engine:
//   Version:          18.09.1
//   API version:      1.39 (minimum version 1.12)
//   Go version:       go1.10.6
//   Git commit:       4c52b90
//   Built:            Wed Jan  9 19:41:49 2019
//   OS/Arch:          linux/amd64
//   Experimental:     true
type ServerVersion struct {
	Version    string
	APIVersion string
	GOVersion  string
	GitCommit  string
	Build      string
	OS         string
	Arch       string
}

// ServerVersion returns server info.
func (client *Client) ServerVersion() (*ServerVersion, error) {
	format := strings.Trim(`
{{ .Server.Version }}
{{ .Server.APIVersion }}
{{ .Server.GoVersion }}
{{ .Server.GitCommit }}
{{ .Server.BuildTime }}
{{ .Server.Os }}
{{ .Server.Arch }}
`, " \t\r\n")
	opt := VersionOption{
		Format: &format,
	}
	out, err := client.Version(opt).Output()
	if err != nil {
		if err != nil {
			e, ok := err.(*exec.ExitError)
			if !ok {
				return nil, err
			}
			return nil, fmt.Errorf("%s", strings.ToLower(string(e.Stderr)))
		}
	}
	s := strings.Split(strings.Trim(string(out), "\"\n"), "\\n")

	return &ServerVersion{
		Version:    s[0],
		APIVersion: s[1],
		GOVersion:  s[2],
		GitCommit:  s[3],
		Build:      s[4],
		OS:         s[5],
		Arch:       s[6],
	}, nil
}

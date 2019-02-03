package docker

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/k-kinzal/aliases/pkg/posix"
)

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
	Version      string
	APIVersion   string
	GOVersion    string
	GitCommit    string
	Build        time.Time
	OS           string
	Arch         string
	Experimental bool
}

// parseClientVersion parses client info.
func parseClientVersion(stdout string) (*ClientVersion, error) {
	res := &ClientVersion{}

	var matches []string
	matches = regexp.MustCompile(`Version:\s+(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.Version = matches[1]

	matches = regexp.MustCompile(`API version:\s+(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.APIVersion = matches[1]

	matches = regexp.MustCompile(`Go version:\s+(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.GOVersion = matches[1]

	matches = regexp.MustCompile(`Git commit:\s+(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.GitCommit = matches[1]

	matches = regexp.MustCompile(`Built:\s+(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	t, err := time.Parse("Mon Jan  2 15:04:05 2006", matches[1])
	if err != nil {
		return nil, err
	}
	res.Build = t

	matches = regexp.MustCompile(`OS/Arch:\s+(.*?)/(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 2 {
		return nil, fmt.Errorf("")
	}
	res.OS = matches[1]
	res.Arch = matches[2]

	matches = regexp.MustCompile(`Experimental:\s+(.*)`).FindStringSubmatch(stdout)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	b, err := strconv.ParseBool(matches[1])
	if err != nil {
		return nil, fmt.Errorf("")
	}
	res.Experimental = b

	return res, nil
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
	Version      string
	APIVersion   string
	GOVersion    string
	GitCommit    string
	Built        time.Time
	OS           string
	Arch         string
	Experimental bool
}

// parseServerVersion parses server info.
func parseServerVersion(stdout string) (*ServerVersion, error) {
	if !strings.Contains(stdout, "Server: Docker Engine") {
		return nil, nil
	}
	res := &ServerVersion{}

	var matches [][]string
	matches = regexp.MustCompile(`Version:\s+(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.Version = matches[1][1]

	matches = regexp.MustCompile(`API version:\s+(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.APIVersion = matches[1][1]

	matches = regexp.MustCompile(`Go version:\s+(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.GOVersion = matches[1][1]

	matches = regexp.MustCompile(`Git commit:\s+(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.GitCommit = matches[1][1]

	matches = regexp.MustCompile(`Built:\s+(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	t, err := time.Parse("Mon Jan  2 15:04:05 2006", matches[1][1])
	if err != nil {
		return nil, err
	}
	res.Built = t

	matches = regexp.MustCompile(`OS/Arch:\s+(.*?)/(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	res.OS = matches[1][1]
	res.Arch = matches[1][2]

	matches = regexp.MustCompile(`Experimental:\s+(.*)`).FindAllStringSubmatch(stdout, 2)
	if len(matches) <= 1 {
		return nil, fmt.Errorf("")
	}
	b, err := strconv.ParseBool(matches[1][1])
	if err != nil {
		return nil, fmt.Errorf("")
	}
	res.Experimental = b

	return res, nil
}

// VersionOption is an option for executing docker version.
type VersionOption struct {
	Format     *string
	Kubeconfig *string
}

// Version is docker info.
type Version struct {
	Client ClientVersion
	Server *ServerVersion
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

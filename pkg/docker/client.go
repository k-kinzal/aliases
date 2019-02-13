package docker

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	"github.com/k-kinzal/aliases/pkg/logger"
)

// Clients provide operations on Docker commands.
type Client struct {
	path string
	host string
	sock *string
}

// Path returns the path to binary of docker.
func (client *Client) Path() string {
	return client.path
}

// Host return DOCKER_HOST
func (client *Client) Host() string {
	return client.host
}

// Proto returns the protocol of DOCKER_HOST
func (client *Client) Proto() string {
	matches := regexp.MustCompile(`^([^:]*)://`).FindStringSubmatch(client.host)
	if len(matches) == 0 {
		return ""
	}
	return matches[1]
}

// Sock returns the path to socket for docker.
func (client *Client) Sock() *string {
	return client.sock
}

var (
	client *Client
	once   sync.Once
)

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	var err error
	once.Do(func() {
		c := &Client{}
		// docker binary path
		c.path = exec.Command("docker").Path
		// docker host
		c.host = os.Getenv("DOCKER_HOST")
		if c.host == "" {
			// https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-socket-option
			c.host = "unix:///var/run/docker.sock"
		}
		if c.Proto() != "unix" {
			logger.Warnf("%s may not working possibility. please same path that you use on the host and the host of `DOCKER_HOST`.", c.host)
		} else {
			sock := strings.TrimPrefix(c.host, "unix://")
			if _, err := os.Stat(sock); err != nil {
				err = fmt.Errorf("%s: no such file. please set DOCKER_HOST", sock)
				return
			}
			c.sock = &sock
		}
		// docker version
		clientVersion, e := c.ClientVersion()
		if e != nil {
			err = fmt.Errorf("docker is not installed. see https://docs.docker.com/install/")
			return
		}
		serverVersion, _ := c.ServerVersion()
		if serverVersion == nil {
			logger.Warnf("cannot connect to the Docker daemon at %s. Is the docker daemon running?", c.host)
		} else {
			if clientVersion.Version != serverVersion.Version {
				logger.Warnf("docker client version `%s` and server version `%s` are different", clientVersion.Version, serverVersion.Version)
			}
		}
		client = c
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

package docker_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/docker"
)

func ExampleClient_Version() {
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	cmd := client.Version(docker.VersionOption{})
	fmt.Println(cmd)
	// Output: docker version
}

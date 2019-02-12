package docker_test

import (
	"fmt"
	"strings"
	"testing"

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

func TestClient_ClientVersion(t *testing.T) {
	client, err := docker.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ClientVersion()
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ServerVersion(t *testing.T) {
	// FIXME: If docker daemon is stopped it needs to work.
	client, err := docker.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ServerVersion()
	if err != nil && !strings.HasPrefix(err.Error(), "cannot connect to the docker daemon") {
		t.Fatal(err)
	}
}

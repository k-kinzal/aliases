package docker_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/k-kinzal/aliases/pkg/docker"
)

func TestMain(m *testing.M) {
	sock, err := os.Create("/tmp/docker.sock")
	if err != nil {
		panic(err)
	}
	defer sock.Close()

	if err := os.Setenv("DOCKER_HOST", fmt.Sprintf("unix://%s", sock.Name())); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func ExampleNewClient() {
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}
	fmt.Println(client.Host())
	// Output: unix:///tmp/docker.sock
}

func ExampleClient_Host() {
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.Host())
	// Output: unix:///tmp/docker.sock
}

func ExampleClient_Proto() {
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.Proto())
	// Output: unix
}

func ExampleClient_Sock() {
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(*client.Sock())
	// Output: /tmp/docker.sock
}

package docker_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/docker"
)

func ptr(s string) *string {
	return &s
}

func ExampleClient_Run() {
	client, err := docker.NewClient()
	if err != nil {
		panic(err)
	}

	opt := docker.RunOption{
		Env: map[string]string{
			"FOO": "bar",
		},
		Interactive: ptr("true"),
		Rm:          ptr("true"),
		TTY:         ptr("true"),
	}
	cmd := client.Run("alpine", []string{"sh", "-c", "echo 1"}, opt)
	fmt.Println(cmd)
	// Output: docker run --env FOO="bar" --interactive --rm --tty alpine sh -c "echo 1"
}

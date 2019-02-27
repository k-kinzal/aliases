package posix_test

import (
	"fmt"
	"os"

	"github.com/k-kinzal/aliases/pkg/posix"
)

func ExampleShell() {
	cmd := posix.Shell("echo $FOO")
	cmd.Stdout = os.Stdout
	cmd.Env = append(cmd.Env, "FOO=1")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	// Output: 1
}

func ExampleShellScript_String() {
	cmd := posix.Shell("echo $FOO")
	fmt.Println(cmd)
	// Output: echo $FOO
}

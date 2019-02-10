package posix_test

import (
	"fmt"
	"os"

	"github.com/k-kinzal/aliases/pkg/posix"
)

func ExampleCommand() {
	cmd := posix.Command("echo", "1")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	// Output: 1
}

func ExampleCmd_String() {
	cmd := posix.Command("echo", "1")
	fmt.Println(cmd)
	// Output: echo 1
}

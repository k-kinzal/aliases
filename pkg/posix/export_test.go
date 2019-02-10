package posix_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/posix"
)

func ExampleExport() {
	cmd := posix.Export("FOO", "{\"foo\": 1}")
	fmt.Println(cmd)
	// Output: export FOO="{\"foo\": 1}"
}

func ExamplePathExport() {
	cmd := posix.PathExport("/bin:/sbin")
	fmt.Println(cmd)
	// Output: export PATH="/bin:/sbin:$PATH"
}

package posix_test

import (
	"fmt"

	"github.com/k-kinzal/aliases/pkg/posix"
)

func ExampleAlias() {
	cmd := posix.Alias("foo", "echo 'Hello World'")
	fmt.Println(cmd)
	// Output: alias foo='echo \'Hello World\''
}

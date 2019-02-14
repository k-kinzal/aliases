package context_test

import (
	"fmt"
	"os"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
)

func ExampleHomePath() {
	if err := context.ChangeHomePath("/tmp"); err != nil {
		panic(err)
	}
	fmt.Println(context.HomePath())
	// Output: /tmp
}

func ExampleConfPath() {
	file, err := os.Create("/tmp/aliases.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := context.ChangeConfPath(file.Name()); err != nil {
		panic(err)
	}
	fmt.Println(context.ConfPath())
	// Output: /tmp/aliases.yaml
}

func ExampleExportPath() {
	if err := context.ChangeExportPath("/tmp/export"); err != nil {
		panic(err)
	}
	fmt.Println(context.ExportPath())
	// Output: /tmp/export
}

func ExampleBinaryPath() {
	if err := context.ChangeHomePath("/tmp"); err != nil {
		panic(err)
	}
	fmt.Println(context.BinaryPath())
	// Output: /tmp/docker
}

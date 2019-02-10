package script_test

import (
	"fmt"
	"os"
	"strconv"

	"github.com/k-kinzal/aliases/pkg/aliases/script"
)

func ExampleExpandColonDelimitedStringListWithEnv() {
	os.Setenv("FOO", "1")
	val := []string{
		"$FOO:$FOO",
	}
	fmt.Println(script.ExpandColonDelimitedStringListWithEnv(val))
	// Output: [1:$FOO]
}

func ExampleExpandColonDelimitedStringWithEnv() {
	os.Setenv("FOO", "2")
	val := "$FOO:$FOO"
	fmt.Println(script.ExpandColonDelimitedStringWithEnv(val))
	// Output: 2:$FOO
}

func ExampleExpandStringKeyMapWithEnv() {
	os.Setenv("FOO", "1")
	val := map[string]string{
		"$FOO": "$FOO",
	}
	fmt.Println(script.ExpandStringKeyMapWithEnv(val))
	// Output: map[1:$FOO]
}

func ExampleExpandEnv() {
	os.Setenv("FOO_BAR_1", "1")
	// expand env
	fmt.Println(script.ExpandEnv("$FOO_BAR_1"))
	fmt.Println(script.ExpandEnv("${FOO_BAR_1}"))
	fmt.Println(script.ExpandEnv(strconv.Quote("$FOO_BAR_1")))
	fmt.Println(script.ExpandEnv(strconv.Quote("${FOO_BAR_1}")))

	// expand env in string
	fmt.Println(script.ExpandEnv("#$FOO_BAR_1#"))
	fmt.Println(script.ExpandEnv("#${FOO_BAR_1}#"))
	fmt.Println(script.ExpandEnv("#\"$FOO_BAR_1\"#"))
	fmt.Println(script.ExpandEnv("#\"${FOO_BAR_1}\"#"))

	// special environment
	fmt.Println(script.ExpandEnv("$PWD"))
	fmt.Println(script.ExpandEnv("${PWD}"))
	fmt.Println(script.ExpandEnv("##$PWD##"))
	fmt.Println(script.ExpandEnv("##${PWD}##"))

	// no expand env
	fmt.Println(script.ExpandEnv("'$FOO_BAR_1'"))
	fmt.Println(script.ExpandEnv("`echo $FOO_BAR_1`"))
	fmt.Println(script.ExpandEnv("$(echo $FOO_BAR_1)"))
	fmt.Println(script.ExpandEnv("$@")) // not expand shell special variables

	// empty
	fmt.Println(script.ExpandEnv("$"))
	fmt.Println(script.ExpandEnv("${}"))
	fmt.Println(script.ExpandEnv("\"\""))
	fmt.Println(script.ExpandEnv("''"))
	fmt.Println(script.ExpandEnv("``"))
	fmt.Println(script.ExpandEnv("$()"))

	// broken pair
	fmt.Println(script.ExpandEnv("${"))
	fmt.Println(script.ExpandEnv("\""))
	fmt.Println(script.ExpandEnv("'"))
	fmt.Println(script.ExpandEnv("`"))
	fmt.Println(script.ExpandEnv("$("))
	fmt.Println(script.ExpandEnv("${ABC"))
	fmt.Println(script.ExpandEnv("${..."))
	fmt.Println(script.ExpandEnv("\"..."))
	fmt.Println(script.ExpandEnv("'..."))
	fmt.Println(script.ExpandEnv("`..."))
	fmt.Println(script.ExpandEnv("$(..."))

	// Output:
	// 1
	// 1
	// "1"
	// "1"
	// #1#
	// #1#
	// #"1"#
	// #"1"#
	// ${ALIASES_PWD:-$PWD}
	// ${ALIASES_PWD:-$PWD}
	// ##${ALIASES_PWD:-$PWD}##
	// ##${ALIASES_PWD:-$PWD}##
	// '$FOO_BAR_1'
	// `echo $FOO_BAR_1`
	// $(echo $FOO_BAR_1)
	// $@
	// $
	// ${}
	// ""
	// ''
	// ``
	// $()
	// ${
	// "
	// '
	// `
	// $(
	// ${ABC
	// ${...
	// "...
	// '...
	// `...
	// $(...
}

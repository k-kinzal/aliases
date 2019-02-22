package posix

import (
	"fmt"
	"strconv"
	"strings"
)

type ShellScript Cmd

// Run shell script.
func (shell *ShellScript) Run() error {
	env := make([]string, 0)
	for _, e := range shell.Env {
		a := strings.Split(e, "=")
		name := a[0]
		val := strings.Join(a[1:], "=")
		env = append(env, fmt.Sprintf("%s=%s", name, strconv.Quote(val)))
	}
	shell.Args[2] = fmt.Sprintf("#!/bin/sh\n%s;\n%s", strings.Join(env, ";\n"), shell.Args[2])
	return shell.Cmd.Run()
}

// String returns body of shell script.
func (shell *ShellScript) String() string {
	return shell.Args[2]
}

// Shell runs the POSIX command as a shell
//
// http://pubs.opengroup.org/onlinepubs/009695399/utilities/sh.html
// TODO: Implement other options if there is a use case
func Shell(commandString string) *ShellScript {
	return (*ShellScript)(Command("sh", "-c", commandString))
}

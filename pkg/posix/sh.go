package posix

import (
	"fmt"
	"strings"
)

type ShellScript Cmd

func (script *ShellScript) Run() error {
	script.Args[2] = fmt.Sprintf("#!/bin/sh\n%s; %s", strings.Join(script.Env, ";\n"), script.Args[2])
	return ((*Cmd)(script)).Run()
}

func (script *ShellScript) String() string {
	return script.Args[2]
}

// http://pubs.opengroup.org/onlinepubs/009695399/utilities/sh.html
// TODO: Implement other options if there is a use case
func Shell(commandString string) *ShellScript {
	cmd := Command("sh", "-c", commandString)

	return (*ShellScript)(cmd)
}

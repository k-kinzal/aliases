package posix

import (
	"fmt"
	"strconv"
	"strings"
)

type ShellScript Cmd

func (script *ShellScript) Run() error {
	env := make([]string, 0)
	for _, e := range script.Env {
		a := strings.Split(e, "=")
		name := a[0]
		val := strings.Join(a[1:], "=")
		env = append(env, fmt.Sprintf("%s=%s", name, strconv.Quote(val)))
	}
	script.Args[2] = fmt.Sprintf("#!/bin/sh\n%s; %s", strings.Join(env, ";\n"), script.Args[2])
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

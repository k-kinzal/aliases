package posix

// http://pubs.opengroup.org/onlinepubs/009695399/utilities/sh.html
// TODO: Implement other options if there is a use case
func Shell(commandString string) *Cmd {
	cmd := Command("sh", "-c", commandString)

	return cmd
}

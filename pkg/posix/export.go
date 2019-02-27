package posix

import (
	"fmt"
	"strconv"
	"strings"
)

// Export is the export wrapper for the POSIX command.
//
// http://pubs.opengroup.org/onlinepubs/007904975/utilities/export.html
func Export(name string, word string) *Cmd {
	cmd := Command("export")

	if name != "" && word != "" {
		cmd.Args = append(cmd.Args, fmt.Sprintf("%s=%s", strings.ToUpper(name), strconv.Quote(word)))
	} else if name != "" {
		cmd.Args = append(cmd.Args, strings.ToUpper(name))
	}

	return cmd
}

// PathExport returns the wrapper command `export PATH = ...`.
func PathExport(word string) *Cmd {
	return Export("PATH", fmt.Sprintf("%s:$PATH", word))
}

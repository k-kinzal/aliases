package posix

import (
	"fmt"
	"strconv"
	"strings"
)

// http://pubs.opengroup.org/onlinepubs/007904975/utilities/export.html
func Export(name string, word string, print bool) *Cmd {
	cmd := Command("export")

	if print {
		cmd.Args = append(cmd.Args, "-p")
	}

	if name != "" && word != "" {
		cmd.Args = append(cmd.Args, fmt.Sprintf("%s=%s", strings.ToUpper(name), strconv.Quote(word)))
	} else if name != "" {
		cmd.Args = append(cmd.Args, strings.ToUpper(name))
	}

	return cmd
}

func PathExport(word string, print bool) *Cmd {
	return Export("PATH", fmt.Sprintf("%s:$PATH", word), print)
}

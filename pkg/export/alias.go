package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
)



func Aliases(cmds []aliases.AliasCommand) {
	for _, cmd := range cmds {
		fmt.Printf(fmt.Sprintf("alias %s='%s'\n", cmd.Filename, cmd.ToString()))
	}
}
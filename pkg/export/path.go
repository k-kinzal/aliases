package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
	"os"
)

func Path(conf *aliases.AliasesConf, ctx *aliases.Context) {
	dir := ctx.GetBinaryPath(conf.Hash)
	os.Remove(dir)
	os.Mkdir(dir, 0755)

	for _, conf := range conf.Commands {
		writeFiles(&conf, dir)
	}

	fmt.Printf("export PATH=\"%s:$PATH\"", dir)
}
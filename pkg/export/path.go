package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
	"github.com/k-kinzal/aliases/pkg/context"
	"os"
)

func Path(ctx *context.Context, conf *aliases.AliasesConf) {
	dir := ctx.GetBinaryPath()
	os.Remove(dir)
	os.Mkdir(dir, 0755)

	for _, conf := range conf.Commands {
		writeFiles(&conf, dir)
	}

	fmt.Printf("export PATH=\"%s:$PATH\"", dir)
}
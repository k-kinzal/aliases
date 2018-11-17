package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
)

func Path(conf aliases.AliasesConf, ctx aliases.Context) {
	fmt.Printf("export PATH=\"%s:$PATH\"", ctx.GetBinaryPath(conf.Hash))

}
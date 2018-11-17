package export

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg"
	"io/ioutil"
	"os"
)

var (
	tmpl = `#!/bin/sh

if [ -p /dev/stdin ]; then
    cat - | %s $@
	exit $?
else
    %s $@
	exit $?
fi
`
)

func Script(cmds []aliases.AliasCommand, conf aliases.AliasesConf, ctx aliases.Context) {
	os.Remove(ctx.GetBinaryPath(conf.Hash))
	os.Mkdir(ctx.GetBinaryPath(conf.Hash), 0755)

	for _, cmd := range cmds {
		str := cmd.ToString()
		path := fmt.Sprintf("%s/%s", ctx.GetBinaryPath(conf.Hash), cmd.Filename)
		content := fmt.Sprintf(tmpl, str, str)

		ioutil.WriteFile(path, []byte(content), 0755)
	}

	fmt.Printf("export PATH=\"%s:$PATH\"", ctx.GetBinaryPath(conf.Hash))

}
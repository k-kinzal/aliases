package aliases

import (
	"fmt"
	"os"
)

type Context struct {
	ConfPath string
}

func NewContext(path string) (*Context, error) {
	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("cannot get PWD `%q`", err)
		}

		path = fmt.Sprintf("%s/aliases.yaml", cwd)
	}
	context := new(Context)
	context.ConfPath = path

	return context, nil
}

package context

import (
	"fmt"
	"os"
	"path"

	"github.com/k-kinzal/aliases/pkg/types"
)

var (
	homePath   string
	confPath   string
	exportPath string
	binaryPath string
)

// makeDir make directory.
func makeDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

// replaceDir replace directory.
func replacetDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	return nil
}

// HomePath returns home directory path for aliases.
func HomePath() string {
	if homePath == "" {
		panic("logic error: home path must be set")
	}
	return homePath
}

// ChangeHomePath changes home directory path for aliases.
func ChangeHomePath(path string) error {
	if err := makeDir(path); err != nil {
		return err
	}
	homePath = path
	return nil
}

// ConfPath returns path of configuration file.
func ConfPath() string {
	if confPath == "" {
		panic("logic error: conf path must be set")
	}
	return confPath
}

// ChangeConfPath changes configuration file path for aliases.
func ChangeConfPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s: no such file or directory", path)
	}
	confPath = path
	return nil
}

// ConfPath returns path of export directory.
func ExportPath() string {
	if exportPath == "" {
		panic("logic error: export path must be set")
	}
	return exportPath
}

// ChangeExportPath changes export directory path for aliases.
func ChangeExportPath(path string) error {
	if err := replacetDir(path); err != nil {
		return err
	}
	exportPath = path
	return nil
}

// BinaryPath returns path of docker binary directory.
func BinaryPath() string {
	p := path.Join(HomePath(), "docker")
	if _, err := os.Stat(p); os.IsNotExist(err) {
		_ = makeDir(binaryPath)
	}

	return p
}

func TemporaryPath(spec interface{}) string {
	p := path.Join(HomePath(), "tmp", types.MD5(spec))
	if _, err := os.Stat(p); os.IsNotExist(err) {
		_ = makeDir(p)
	}

	return p
}

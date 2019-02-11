package context

import (
	"fmt"
	"os"
	"path"
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
		if err := os.Mkdir(path, 0755); err != nil {
			return fmt.Errorf("runtime error: %s", err)
		}
	}
	return nil
}

// replaceDir replace directory.
func replacetDir(path string) error {
	if err := os.RemoveAll(path); err != nil {
		return fmt.Errorf("runtime error: %s", err)
	}
	if err := os.Mkdir(path, 0755); err != nil {
		return fmt.Errorf("runtime error: %s", err)
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
		return fmt.Errorf("runtime error: %s: No such file or directory", path)
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
	if binaryPath == "" {
		binaryPath = path.Join(HomePath(), "docker")
		_ = makeDir(binaryPath)
	}
	return binaryPath
}

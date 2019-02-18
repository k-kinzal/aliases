package script

import (
	"fmt"
	"os"
	"path"

	"github.com/k-kinzal/aliases/pkg/aliases/context"
	"github.com/k-kinzal/aliases/pkg/docker"
)

// Write exports aliases script to a file.
func (script *Script) Write() (string, error) {
	return script.WriteWithOverride(nil, docker.RunOption{})
}

// Write exports aliases script to a file with override docker option.
func (script *Script) WriteWithOverride(args []string, option docker.RunOption) (string, error) {
	for _, cmd := range script.relative {
		if _, err := cmd.Write(); err != nil {
			return "", err
		}
	}

	if script.entrypoint.body != "" {
		if err := script.WriteExtendEntrypoint(); err != nil {
			return "", err
		}
	}

	targetPath := script.Path(context.ExportPath())

	if err := os.MkdirAll(path.Dir(targetPath), 0755); err != nil {
		return "", err
	}

	file, err := os.Create(targetPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := os.Chmod(targetPath, 0755); err != nil {
		return "", err
	}

	shell, err := script.Shell(args, option)
	if err != nil {
		return "", err
	}

	if _, err := file.Write([]byte(fmt.Sprintf("#!/bin/sh\n%s", shell.String()))); err != nil {
		return "", err
	}

	return targetPath, nil
}

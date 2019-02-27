package script

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/k-kinzal/aliases/pkg/logger"

	"github.com/k-kinzal/aliases/pkg/posix"

	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/k-kinzal/aliases/pkg/aliases/context"

	"github.com/k-kinzal/aliases/pkg/aliases/yaml"
	"github.com/k-kinzal/aliases/pkg/docker"
)

// Script is the actual of command alises.
type Script struct {
	spec yaml.Option
}

// Write exports aliases script to a file.
func (script *Script) Write(client *docker.Client) error {
	return script.WriteWithOverride(client, nil, docker.RunOption{})
}

// WriteWithOverride exports aliases script to a file with override docker option.
func (script *Script) WriteWithOverride(client *docker.Client, overrideArgs []string, overrideOption docker.RunOption) error {
	targetPath := filepath.Join(context.ExportPath(), script.spec.Namespace(), script.spec.Path.Base())

	if err := os.MkdirAll(path.Dir(targetPath), 0755); err != nil {
		return err
	}

	fp, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		return nil
	}
	defer fp.Close()

	shell := adaptShell(script.spec)
	cmd, err := shell.Command(client, overrideArgs, overrideOption, false)
	if err != nil {
		return err
	}

	if _, err := fp.WriteString(fmt.Sprintf("#!/bin/sh\n%s", cmd.String())); err != nil {
		return err
	}

	return nil
}

// Alias returns the command of alias.
func (script *Script) Alias(client *docker.Client) (*posix.Cmd, error) {
	if err := script.Write(client); err != nil {
		return nil, err
	}
	targetPath := filepath.Join(context.ExportPath(), script.spec.Path.Base())

	return posix.Alias(script.spec.Path.Base(), targetPath), nil
}

// Run aliases script.
func (script *Script) Run(client *docker.Client, overrideArgs []string, overrideOption docker.RunOption) error {
	targetPath := filepath.Join(context.ExportPath(), script.spec.Namespace(), script.spec.Path.Base())

	if err := os.MkdirAll(path.Dir(targetPath), 0755); err != nil {
		return err
	}

	fp, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755) // set lock
	if err != nil {
		return nil
	}
	defer fp.Close()

	shell := adaptShell(script.spec)
	cmd, err := shell.Command(client, overrideArgs, overrideOption, logger.LogLevel() == logger.DebugLevel)
	if err != nil {
		return err
	}
	cmd.Env = os.Environ()

	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if (info.Mode()&os.ModeNamedPipe) != 0 || info.Size() > 0 {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		cmd.Stdin = os.Stdin
	} else if (info.Mode() & os.ModeCharDevice) != 0 {
		cmd.Stdin = os.Stdin
	} else {
		cmd.Stdin = nil
	}

	cmd.Stdout = os.Stdout

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil
	}

	fn := func() error {
		if err := cmd.Start(); err != nil {
			return err
		}

		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			line := scanner.Text()
			// invalid argument error converts go error
			if strings.HasPrefix(line, "invalid argument") {
				r := regexp.MustCompile(`^invalid argument "(.*?)" for "(.*?)" flag: (.*)$`)
				matches := r.FindStringSubmatch(line)
				if len(matches) > 0 {
					return util.FlagError(matches[1], matches[2], matches[3])
				}
			}
			// remove docker prefix
			if strings.HasPrefix(line, "docker: Error response from daemon: ") {
				fmt.Fprintln(os.Stderr, "aliases: error response from daemon: "+strings.TrimPrefix(line, "docker: Error response from daemon: "))
				continue
			}
			// suppress guidance to help
			if line == "See 'docker run --help'." {
				continue
			}
			fmt.Fprintln(os.Stderr, line)
		}

		return cmd.Wait()
	}

	if cmd.SysProcAttr != nil {
		// FIXME: timeout get from command line arguments
		if err := util.Timeout(3*time.Second, fn); err != nil {
			if _, ok := err.(*util.TimeoutError); ok {
				_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
				return nil
			}
			return err
		}
	} else {
		return fn()
	}

	return nil
}

// NewScript creates a new Script.
func NewScript(spec yaml.Option) *Script {
	return &Script{spec}
}

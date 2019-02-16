package script

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/k-kinzal/aliases/pkg/docker"
)

// Run aliases script.
func (script *Script) Run(args []string, opt docker.RunOption) error {
	for _, relative := range script.relative {
		if _, err := relative.Write(); err != nil {
			return err
		}
	}

	dockerCmdString := script.docker(args, opt).String()
	logger.Debug(dockerCmdString)

	command := posix.Shell(dockerCmdString)
	command.Env = os.Environ()

	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}
	if (info.Mode()&os.ModeNamedPipe) != 0 || info.Size() > 0 {
		command.Stdin = os.Stdin
	} else {
		command.Stdin = nil
	}

	command.Stdout = os.Stdout

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil
	}

	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	fn := func() error {
		if err := command.Start(); err != nil {
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

		return command.Wait()
	}

	if (info.Mode()&os.ModeNamedPipe) != 0 && info.Size() == 0 {
		// FIXME: timeout get from command line arguments
		if err := util.Timeout(3*time.Second, fn); err != nil {
			if _, ok := err.(*util.TimeoutError); ok {
				_ = syscall.Kill(-command.Process.Pid, syscall.SIGKILL)
				return nil
			}
			return err
		}
	} else {
		return fn()
	}

	return nil
}

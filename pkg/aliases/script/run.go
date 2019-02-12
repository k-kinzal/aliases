package script

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/k-kinzal/aliases/pkg/logger"
	"github.com/k-kinzal/aliases/pkg/util"

	"github.com/k-kinzal/aliases/pkg/docker"

	"github.com/k-kinzal/aliases/pkg/posix"
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
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil
	}

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
		// suppress guidance to help
		if line == "See 'docker run --help'." {
			continue
		}
		fmt.Fprintln(os.Stderr, line)
	}

	return command.Wait()
}

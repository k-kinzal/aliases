package yaml

import (
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var (
	pathRegexp       = regexp.MustCompile(`^[^.]+(\.dependencies\[\d+\]\.[^.]+)*$`)
	dependencyRegexp = regexp.MustCompile(`dependencies\[(\d+)\]`)
	baseRegexp       = regexp.MustCompile(`(^|dependencies\[\d+\]\.)([^.]+)$`)
)

// SpecPath is the path that points to the location of OptionSpec.
type SpecPath string

// Parent returns the path that points to the location of the parent OptionSpec.
func (p *SpecPath) Parent() *SpecPath {
	matches := pathRegexp.FindStringSubmatch(p.String())
	if len(matches) <= 1 {
		return nil
	}
	if matches[len(matches)-1] == "" {
		return nil
	}
	val := (SpecPath)(strings.TrimSuffix(p.String(), matches[len(matches)-1]))
	return &val
}

// Dependencies returns the path point to the location of dependency OptionSpec.
func (p *SpecPath) Dependencies(i int, index string) *SpecPath {
	matches := pathRegexp.FindStringSubmatch(p.String())
	if len(matches) == 0 {
		return nil
	}
	val := SpecPath(fmt.Sprintf("%s.dependencies[%d].%s", p.String(), i, index))
	return &val
}

// Name returns name of OptionSpec.
func (p *SpecPath) Name() string {
	matches := baseRegexp.FindStringSubmatch(p.String())
	if len(matches) == 0 {
		return ""
	}
	return matches[2]
}

// Base returns filename of OptionSpec.
func (p *SpecPath) Base() string {
	return path.Base(p.Name())
}

// Index returns index of parent dependencies.
func (p *SpecPath) Index() int {
	matches := dependencyRegexp.FindStringSubmatch(*(*string)(p))
	if len(matches) == 0 {
		return -1
	}
	n, err := strconv.Atoi(matches[len(matches)-1])
	if err != nil {
		return -1
	}

	return n
}

// String converts string from SpecPath
func (p *SpecPath) String() string {
	return *(*string)(p)
}

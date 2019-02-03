package yaml

import (
	"fmt"
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
func (path *SpecPath) Parent() *SpecPath {
	matches := pathRegexp.FindStringSubmatch(path.String())
	if len(matches) <= 1 {
		return nil
	}
	if matches[len(matches)-1] == "" {
		return nil
	}
	p := (SpecPath)(strings.TrimSuffix(path.String(), matches[len(matches)-1]))
	return &p
}

// Dependencies returns the path point to the location of dependency OptionSpec.
func (path *SpecPath) Dependencies(i int, index string) *SpecPath {
	matches := pathRegexp.FindStringSubmatch(path.String())
	if len(matches) == 0 {
		return nil
	}
	p := SpecPath(fmt.Sprintf("%s.dependencies[%d].%s", path.String(), i, index))
	return &p
}

// Base returns name of OptionSpec.
func (path *SpecPath) Base() string {
	matches := baseRegexp.FindStringSubmatch(path.String())
	if len(matches) == 0 {
		return ""
	}
	return matches[2]
}

// Index returns index of parent dependencies.
func (path *SpecPath) Index() int {
	matches := dependencyRegexp.FindStringSubmatch(*(*string)(path))
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
func (path *SpecPath) String() string {
	return *(*string)(path)
}

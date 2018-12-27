package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

type helper struct {
	*cli.Context
}

func (h *helper) bool(names ...string) *string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return nil
	}

	val := fmt.Sprintf("%t", h.Context.Bool(name))
	return &val
}

func (h *helper) int(names ...string) *string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return nil
	}

	val := strconv.Itoa(h.Context.Int(name))
	return &val
}

func (h *helper) int64(names ...string) *string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return nil
	}

	val := strconv.FormatInt(h.Context.Int64(name), 64)
	return &val
}

func (h *helper) uint16(names ...string) *string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return nil
	}

	val := strconv.FormatUint(h.Context.Uint64(name), 16)
	return &val
}

func (h *helper) string(names ...string) *string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return nil
	}

	val := h.Context.String(name)
	return &val
}

func (h *helper) stringSlice(names ...string) []string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return []string{}
	}

	return h.Context.StringSlice(name)
}

func (h *helper) stringMap(names ...string) map[string]string {
	var name string
	for _, n := range names {
		if h.Context.IsSet(n) {
			name = n
		}
	}
	if name == "" {
		return map[string]string{}
	}

	val := map[string]string{}
	for _, v := range h.Context.StringSlice(name) {
		s := strings.Split(v, "=")
		val[s[0]] = strings.Join(s[1:], "=")
	}

	return val
}

func (h *helper) args() []string {
	var args []string
	for _, arg := range h.Context.Args()[1:] {
		if arg == "--" {
			continue
		}
		args = append(args, arg)
	}
	return args
}

func (h *helper) firstArg() *string {
	args := h.Context.Args()
	if len(args) == 0 {
		return nil
	}
	if args[0] == "--" {
		if len(args) == 1 {
			return nil
		}
		return &args[1]
	}
	return &args[0]
}

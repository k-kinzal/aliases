package aliases

import (
	"os"
	"strings"
)

func expandColonDelimitedStringWithEnv(s string) string {
	arr := strings.Split(s, ":")
	if len(arr) > 0 {
		arr[0] = os.ExpandEnv(arr[0])
	}
	return strings.Join(arr, ":")
}

func expandColonDelimitedStringListWithEnv(arr []string) []string {
	rets := make([]string, 0)
	for _, v := range arr {
		rets = append(rets, expandColonDelimitedStringWithEnv(v))
	}

	return rets
}

func expandStringKeyMapWithEnv(m map[string]string) map[string]string {
	rets := make(map[string]string, 0)
	for k, v := range m {
		rets[os.ExpandEnv(k)] = v
	}

	return rets
}
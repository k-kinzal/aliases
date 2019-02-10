package script

import (
	"os"
	"strings"
)

func expandEnv(str string) string {
	from := strings.Index(str, "$(")
	to := strings.LastIndex(str, ")")
	command := ""
	if from != -1 && to != -1 {
		command = str[from:to]
		str = strings.Replace(str, command, "{{ COMMAND }}", -1)
	}

	str = strings.Replace(str, "$PWD", "{{ ALIASES_PWD }}", -1)
	str = os.ExpandEnv(str)
	str = strings.Replace(str, "{{ ALIASES_PWD }}", "${ALIASES_PWD:-$PWD}", -1)
	str = strings.Replace(str, "{{ COMMAND }}", command, -1)

	return str
}

func expandColonDelimitedStringWithEnv(s string) string {
	arr := strings.Split(s, ":")
	if len(arr) > 0 {
		arr[0] = expandEnv(arr[0])
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
	rets := make(map[string]string, len(m))
	for k, v := range m {
		rets[expandEnv(k)] = v
	}

	return rets
}

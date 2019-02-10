package script

import (
	"os"
	"strings"
)

func ExpandEnv(str string) string {
	expanded := ""
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '"': // "..." expand env
			prefix := expanded
			expanded = expanded + "\""
			for j := i + 1; j < len(str); j++ {
				expanded = expanded + string(str[j])
				if str[j] == '"' {
					expanded = prefix + "\"" + ExpandEnv(str[(i+1):j]) + "\""
					i = j
					break
				}
				if j+1 == len(str) {
					i = j
					break
				}

			}
		case '\'': // '...' no expand env
			expanded = expanded + "'"
			for i++; i < len(str); i++ {
				expanded = expanded + string(str[i])
				if str[i] == '\'' {
					break
				}
			}
		case '`': // `...` no expand env
			expanded = expanded + "`"
			for i++; i < len(str); i++ {
				expanded = expanded + string(str[i])
				if str[i] == '`' {
					break
				}
			}
		case '$':
			prefix := expanded
			expanded = expanded + "$"
			if i+1 < len(str) && str[i+1] == '(' { // $(...) no expand env
				i++
				expanded = expanded + "("
				for i = i + 1; i < len(str); i++ {
					expanded = expanded + string(str[i])
					if str[i] == ')' {
						break
					}
				}
			} else if i+1 < len(str) && str[i+1] == '{' { // ${...} expand env
				i++
				expanded = expanded + "{"
				for j := i + 1; j < len(str); j++ {
					c := str[j]
					expanded = expanded + string(c)
					if c == '}' {
						s := str[(i - 1):(j + 1)]
						if s == "${}" {
							// no expand
						} else if s == "${PWD}" {
							s = "${ALIASES_PWD:-$PWD}"
						} else {
							s = os.ExpandEnv(s)
						}
						expanded = prefix + s
						i = j
						break
					}
					if !(c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
						i = j
						break
					}
					if j == len(str)-1 {
						i = j
						break
					}
				}
			} else { // $AZaz09 expand env
				for j := i + 1; j < len(str); j++ {
					c := str[j]
					expanded = expanded + string(c)
					if (c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') && j == len(str)-1 {
						// $Azaz09
						s := str[i:]
						if s == "$PWD" {
							s = "${ALIASES_PWD:-$PWD}"
						} else {
							s = os.ExpandEnv(s)
						}
						expanded = prefix + s
						i = j
						break
					} else if !(c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z') {
						// $Azaz09#
						s := str[i:j]
						if s == "$PWD" {
							s = "${ALIASES_PWD:-$PWD}"
						} else {
							s = os.ExpandEnv(s)
						}
						expanded = prefix + s
						i = j - 1
						break
					}
				}
			}
		default:
			expanded = expanded + string(str[i])
		}
	}
	//from := strings.Index(str, "$(")
	//to := strings.LastIndex(str, ")")
	//command := ""
	//if from != -1 && to != -1 {
	//	command = str[from:to]
	//	str = strings.Replace(str, command, "{{ COMMAND }}", -1)
	//}
	//
	//str = strings.Replace(str, "$PWD", "{{ ALIASES_PWD }}", -1)
	//str = os.ExpandEnv(str)
	//str = strings.Replace(str, "{{ ALIASES_PWD }}", "${ALIASES_PWD:-$PWD}", -1)
	//str = strings.Replace(str, "{{ COMMAND }}", command, -1)

	return expanded
}

func ExpandColonDelimitedStringWithEnv(s string) string {
	arr := strings.Split(s, ":")
	if len(arr) > 0 {
		arr[0] = ExpandEnv(arr[0])
	}
	return strings.Join(arr, ":")
}

func ExpandColonDelimitedStringListWithEnv(arr []string) []string {
	rets := make([]string, 0)
	for _, v := range arr {
		rets = append(rets, ExpandColonDelimitedStringWithEnv(v))
	}

	return rets
}

func ExpandStringKeyMapWithEnv(m map[string]string) map[string]string {
	rets := make(map[string]string, len(m))
	for k, v := range m {
		rets[ExpandEnv(k)] = v
	}

	return rets
}

package util

import "sort"

// StringKeys returns sorted keys.
func StringKeys(v map[string]string) []string {
	keys := make([]string, len(v))
	i := 0
	for k := range v {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

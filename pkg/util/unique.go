package util

// UniqueStringSlice returns unique slice.
func UniqueStringSlice(data []string) []string {
	slice := make([]string, 0)
	u := map[string]int{}
	for i, v := range data {
		if _, ok := u[v]; ok {
			continue
		}
		u[v] = i
		slice = append(slice, v)
	}
	return slice
}

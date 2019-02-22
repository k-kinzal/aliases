package util

func UniqueStringSlice(data []string) []string {
	var slice []string
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

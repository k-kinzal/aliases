package types

import "sort"

// Set of data structure.
type Set struct {
	hash Hasher
	data map[string]interface{}
}

// Add values to set.
func (set *Set) Add(i interface{}) {
	index := set.hash(i)
	if _, ok := set.data[index]; ok {
		return
	}
	set.data[index] = i
}

// Slice converts from set.
func (set *Set) Slice() []interface{} {
	i := 0
	keys := make([]string, len(set.data))
	for key, _ := range set.data {
		keys[i] = key
		i++
	}
	sort.Strings(keys)

	j := 0
	slice := make([]interface{}, len(set.data))
	for _, key := range keys {
		slice[j] = set.data[key]
		j++
	}
	return slice
}

// NewSet creates a new Set.
func NewSet(hasher Hasher) *Set {
	if hasher == nil {
		hasher = MD5
	}
	return &Set{hasher, make(map[string]interface{})}
}

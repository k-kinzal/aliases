package types

type Set struct {
	hash Hasher
	data map[string]interface{}
}

func (set *Set) Add(i interface{}) {
	index := set.hash(i)
	if _, ok := set.data[index]; ok {
		return
	}
	set.data[index] = i
}

func (set *Set) Slice() []interface{} {
	i := 0
	slice := make([]interface{}, len(set.data))
	for _, val := range set.data {
		slice[i] = val
		i++
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

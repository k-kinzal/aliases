package types

// Stack is the FILO data structure.
type Stack struct {
	hash  Hasher
	slice []interface{}
	index map[string]interface{}
}

// Push adds data to last of Stack.
func (stack *Stack) Push(v interface{}) {
	stack.slice = append(stack.slice, v)
	stack.index[stack.hash(v)] = &v
}

// Pop get data from first.
func (stack *Stack) Pop() interface{} {
	if len(stack.slice) == 0 {
		return nil
	}
	v := stack.slice[len(stack.slice)-1]
	stack.slice = stack.slice[:len(stack.slice)-1]
	delete(stack.index, stack.hash(v))
	return v
}

// Has returns whether there is the same data.
func (stack *Stack) Has(v interface{}) bool {
	_, ok := stack.index[stack.hash(v)]
	return ok
}

// Slice converts from Stack to Slice.
func (stack *Stack) Slice() []interface{} {
	slice := make([]interface{}, len(stack.slice))
	for i := len(stack.slice) - 1; i >= 0; i-- {
		slice[len(stack.slice)-i-1] = stack.slice[i]
	}
	return slice
}

// NewStack creates a new Stack.
func NewStack(hasher Hasher) *Stack {
	if hasher == nil {
		hasher = MD5
	}
	return &Stack{hasher, make([]interface{}, 0), make(map[string]interface{})}
}

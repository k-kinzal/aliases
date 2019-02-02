package types

type Stack struct {
	hash  func(v interface{}) string
	slice []interface{}
	index map[string]interface{}
}

func (stack *Stack) Push(v interface{}) {
	stack.slice = append(stack.slice, v)
	stack.index[stack.hash(v)] = v
}

func (stack *Stack) Pop() interface{} {
	if len(stack.slice) == 0 {
		return nil
	}
	v := stack.slice[len(stack.slice)-1]
	stack.slice = stack.slice[:len(stack.slice)-1]
	delete(stack.index, stack.hash(v))
	return v
}

func (stack *Stack) Has(v interface{}) bool {
	_, ok := stack.index[stack.hash(v)]
	return ok
}

func (stack *Stack) Slice() []interface{} {
	slice := make([]interface{}, len(stack.slice))
	for i := len(stack.slice) - 1; i >= 0; i-- {
		slice[len(stack.slice)-i-1] = stack.slice[i]
	}
	return slice
}

func NewStack(hasher *func(v interface{}) string) *Stack {
	if hasher == nil {
		fn := SHA256
		hasher = &fn
	}
	return &Stack{*hasher, make([]interface{}, 0), make(map[string]interface{})}
}

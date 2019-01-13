package util

import (
	"crypto/md5"
	"encoding/base64"
	"unsafe"
)

type Stack struct {
	slice []interface{}
	index map[string]interface{}
}

func (stack *Stack) Index(v interface{}) string {
	size := unsafe.Sizeof(v)
	b := (*[1 << 10]uint8)(unsafe.Pointer(&v))[:size:size]
	h := md5.New()
	return base64.StdEncoding.EncodeToString(h.Sum(b))
}

func (stack *Stack) Push(v interface{}) {
	stack.slice = append(stack.slice, v)
	stack.index[stack.Index(v)] = v
}

func (stack *Stack) Pop() interface{} {
	if len(stack.slice) == 0 {
		return nil
	}
	v := stack.slice[len(stack.slice)-1]
	stack.slice = stack.slice[:len(stack.slice)-1]
	delete(stack.index, stack.Index(v))
	return v
}

func (stack *Stack) Has(v interface{}) bool {
	_, ok := stack.index[stack.Index(v)]
	return ok
}

func (stack *Stack) Slice() []interface{} {
	slice := make([]interface{}, len(stack.slice))
	for i := len(stack.slice) - 1; i >= 0; i-- {
		slice[len(stack.slice)-i-1] = stack.slice[i]
	}
	return slice
}

func NewStack() *Stack {
	return &Stack{make([]interface{}, 0), make(map[string]interface{}, 0)}
}

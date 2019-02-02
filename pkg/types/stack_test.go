package types

import "fmt"

func ExampleNewStack() {
	var hasher Hasher = MD5
	stack := NewStack(hasher)

	stack.Push(1)
	fmt.Println(stack.Pop())
	// Output: 1
}

func ExampleStack_Push() {
	var hasher Hasher = MD5
	stack := NewStack(hasher)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop())
	// Output: 3 2 1
}

func ExampleStack_Pop() {
	var hasher Hasher = MD5
	stack := NewStack(hasher)

	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())

	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
	// Output: 2
	// 4 3 1 <nil>
}

func ExampleStack_Has() {
	var hasher Hasher = MD5
	stack := NewStack(hasher)

	fmt.Println(stack.Has(1))

	stack.Push(1)
	fmt.Println(stack.Has(1))
	// Output: false
	// true
}

func ExampleStack_Slice() {
	var hasher Hasher = MD5
	stack := NewStack(hasher)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Slice())
	// Output: [3 2 1]
}

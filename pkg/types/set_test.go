package types

import "fmt"

func ExampleNewSet() {
	set := NewSet(nil)
	set.Add(1)
	set.Add(2)
	set.Add(2)
	set.Add(3)

	slice := set.Slice()
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	// Output:
	// 1
	// 2
	// 3
}

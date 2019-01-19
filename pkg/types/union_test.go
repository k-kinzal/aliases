package types_test

import (
	"fmt"
	"github.com/k-kinzal/aliases/pkg/types"
	"reflect"
	"testing"
)

func ExampleNewUnion() {
	union := types.NewUnion(reflect.Int, reflect.String, "abc")
	fmt.Printf("%s", union)
	// Output: string("abc")
}

func ExampleUnion_Type() {
	union := types.NewUnion(reflect.Int, reflect.String, "abc")
	switch union.Type() {
	case union.Left():
		fmt.Printf("Union is `%s`", union.Left())
		break
	case union.Right():
		fmt.Printf("Union is `%s`", union.Right())
		break
	}
	// Output: Union is `string`
}

func ExampleUnion_Left() {
	union := types.NewUnion(reflect.Int, reflect.String, "abc")
	fmt.Println(union.Left())
	// Output: int
}

func ExampleUnion_Right() {
	union := types.NewUnion(reflect.Int, reflect.String, "abc")
	fmt.Println(union.Right())
	// Output: string
}

func ExampleUnion_Value() {
	union := types.NewUnion(reflect.Int, reflect.String, "abc")
	fmt.Println(union.Value())
	// Output: abc
}

func TestNewUnionFailed(t *testing.T) {
	defer func() {
		err := recover()
		if err != "value is expected to be type `int` or `string`, but the actual is `bool`" {
			t.Errorf("not expect message of \"%s\"", err)
		}
	}()
	types.NewUnion(reflect.Int, reflect.String, true)
	t.Error("expected that `panic()` but did not occur")
}

package types

import (
	"fmt"
	"reflect"
)

// Union has the value of type Left or Right.
type Union struct {
	left  reflect.Kind
	right reflect.Kind
	value interface{}
}

// Type returns the type of Left or Right.
func (u *Union) Type() reflect.Kind {
	return reflect.TypeOf(u.value).Kind()
}

// Left returns the type of Left.
func (u *Union) Left() reflect.Kind {
	return u.left
}

// Right returns the type of Right.
func (u *Union) Right() reflect.Kind {
	return u.right
}

// Value returns the value of type Left or Right.
// But you should extend the function returning the appropriate type without calling Value.
func (u *Union) Value() interface{} {
	return u.value
}

// Union to String
func (u *Union) String() string {
	if reflect.TypeOf(u.value).Kind() == u.left {
		return fmt.Sprintf("%s(%#v)", u.left, u.value)
	} else {
		return fmt.Sprintf("%s(%#v)", u.right, u.value)
	}
}

// NewUnion returns Union
func NewUnion(left reflect.Kind, right reflect.Kind, value interface{}) *Union {
	kind := reflect.TypeOf(value).Kind()
	if kind != left && kind != right {
		panic(fmt.Sprintf("value is expected to be type `%s` or `%s`, but the actual is `%T`", left, right, value))
	}
	return &Union{left, right, value}
}

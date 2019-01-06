package stacks

import (
	"DataStruct/array"
	"fmt"
)

/*
this stacks is base for array to Realization
这个栈的底层实现是使用的数组
*/
type MyStacks struct {
	array *array.MyArray
}

func NewMyStacks(cap int) *MyStacks {
	return &MyStacks{array.NewMyArray(cap)}
}

func (ms MyStacks) IsEmpty() bool {
	return ms.array.IsEmpty()
}

func (ms MyStacks) GetSize() int {
	return ms.array.GetSize()
}

func (ms MyStacks) Push(val interface{}) error {
	return ms.array.AddAtLast(val)
}

func (ms MyStacks) Pop() (interface{}, error) {
	return ms.array.DelByIndex(ms.array.GetSize() - 1)
}

func (ms MyStacks) Peek() (interface{}, error) {
	return ms.array.GetByIndex(ms.array.GetSize() - 1)
}

func MyStacksDemo() {
	stacks := NewMyStacks(2)
	boolStacks := stacks.IsEmpty()
	fmt.Println(boolStacks, stacks.array)
	err := stacks.Push(1)
	err = stacks.Push(2)
	err = stacks.Push(3)
	fmt.Println(err, stacks.array)
	val, err := stacks.Pop()
	fmt.Println(val, err, stacks.array)
}

package stacks

import (
	"DataStruct/linkList"
	"fmt"
)

type LinkListStacks struct {
	data *linkList.MyLinkList
}

func (lls LinkListStacks) String() string {
	return lls.data.String()
}

func NewLinkListStacks() *LinkListStacks {
	return &LinkListStacks{
		data: linkList.NewMyLinkList(),
	}
}

func (lls LinkListStacks) IsEmpty() bool {
	return lls.data.IsEmpty()
}

func (lls LinkListStacks) GetSize() int {
	return lls.data.GetSize()
}

func (lls *LinkListStacks) Push(val interface{}) error {
	return lls.data.AddFirst(val)
}

func (lls *LinkListStacks) Pop() (interface{}, error) {
	return lls.data.RemoveFrist()
}

func (lls *LinkListStacks) Peek() (interface{}, error) {
	return lls.data.GetByIndex(0)
}

func LlsDemo() {
	stacks := NewLinkListStacks()
	boolStacks := stacks.IsEmpty()
	fmt.Println(boolStacks, stacks.data)
	err := stacks.Push(1)
	err = stacks.Push(2)
	err = stacks.Push(3)
	fmt.Println(err, stacks.data)
	val, err := stacks.Pop()
	fmt.Println(val, err, stacks.data)
	val, err = stacks.Pop()
	fmt.Println(val, err, stacks.data)
	val, err = stacks.Pop()
	fmt.Println(val, err, stacks.data)
	val, err = stacks.Pop()
	fmt.Println(val, err, stacks.data)
}
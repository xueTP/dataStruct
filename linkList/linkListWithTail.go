package linkList

import (
	"DataStruct/util"
	"errors"
	"fmt"
)

type LinkListWithTail struct {
	base *MyLinkList
	tail *Node
}

func (llwt LinkListWithTail) String() string {
	s := llwt.base.String()
	s += "tail: " + util.InterfaceToString(llwt.tail.Val)
	return s
}

func NewLinkListWithTail() *LinkListWithTail {
	llwt := &LinkListWithTail{
		base: NewMyLinkList(),
	}
	llwt.tail = llwt.base.head
	return llwt
}

func (llwt LinkListWithTail) IsEmpty() bool {
	return llwt.tail == llwt.base.head
}

func (llwt LinkListWithTail) GetSize() int {
	return llwt.base.GetSize()
}

func (llwt *LinkListWithTail) AddByIndex(index int, val interface{}) error {
	err := llwt.base.checkIndex(index)
	if err != nil {
		return err
	}
	prev := llwt.base.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	if index == llwt.base.GetSize() {
		llwt.tail = NewNodeWithNext(val, nil)
		prev.Next = llwt.tail
	} else {
		prev.Next = NewNodeWithNext(val, prev.Next)
	}
	llwt.base.size += 1
	return nil
}

func (llwt *LinkListWithTail) AddFirst(val interface{}) error {
	return llwt.AddByIndex(0, val)
}

func (llwt *LinkListWithTail) AddLast(val interface{}) error {
	return llwt.AddByIndex(llwt.base.GetSize(), val)
}

func (llwt LinkListWithTail) GetByIndex(index int) (interface{}, error) {
	return llwt.base.GetByIndex(index)
}

func (llwt *LinkListWithTail) RemoveByIndex(index int) (interface{}, error) {
	if llwt.base.IsEmpty() {
		return nil, errors.New("this linkList is empty")
	}
	err := llwt.base.checkIndex(index)
	if err != nil {
		return nil, err
	}
	prev := llwt.base.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	if index == llwt.base.GetSize() {
		llwt.tail = prev
	}
	delNode := prev.Next
	prev.Next = prev.Next.Next
	delNode.Next = nil
	llwt.base.size -= 1
	return delNode.Val, nil
}

func (llwt *LinkListWithTail) RemoveFrist() (interface{}, error) {
	return llwt.RemoveByIndex(0)
}

func (llwt *LinkListWithTail) RemoveLast() (interface{}, error) {
	return llwt.RemoveByIndex(llwt.base.GetSize() - 1)
}

func LinkListWithTailDemo() {
	llwt := NewLinkListWithTail()
	for i := 1; i <= 3; i++ {
		llwt.AddFirst(i)
	}
	fmt.Println(llwt)
	err := llwt.AddByIndex(2, 4)
	fmt.Println(llwt, err)
	val, err := llwt.RemoveByIndex(1)
	fmt.Println(llwt, val, err)
	val, err = llwt.RemoveLast()
	fmt.Println(llwt, val, err)
	err = llwt.AddLast(6)
	fmt.Println(llwt, err)
}

package linkList

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
	"DataStruct/util"
)

type MyLinkList struct {
	head *Node
	size int
}

func NewMyLinkList() *MyLinkList {
	return &MyLinkList{head: NewNode(nil), size: 0}
}

func (mll MyLinkList) String() string {
	s := "this LinkList : size: "+ strconv.Itoa(mll.GetSize()) +" data: ["
	cur := mll.head.Next
	for cur != nil {
		s += util.InterfaceToString(cur.Val) + "->"
		cur = cur.Next
	}
	s = strings.Trim(s, "->")
	s += "]"
	return s
}

func (mll MyLinkList) IsEmpty() bool {
	if mll.size == 0 {
		return true
	}
	return false
}

func (mll MyLinkList) GetSize() int {
	return mll.size
}

func (mll MyLinkList) checkIndex(index int) error {
	if index < 0 || index > mll.size  {
		return errors.New("this index is not legitimate")
	}
	return nil
}

func (mll *MyLinkList) AddByIndex(index int, val interface{}) error {
	err := mll.checkIndex(index)
	if err != nil {
		return err
	}
	prev := mll.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = NewNodeWithNext(val, prev.Next)
	mll.size += 1
	return nil
}

func (mll *MyLinkList) AddFirst(val interface{}) error {
	return mll.AddByIndex(0, val)
}

func (mll *MyLinkList) AddLast(val interface{}) error {
	return mll.AddByIndex(mll.size, val)
}

func (mll MyLinkList) GetByIndex(index int) (interface{}, error) {
	err := mll.checkIndex(index)
	if mll.IsEmpty() {
		return nil, errors.New("this linkList is empty")
	}
	if err != nil {
		return nil, err
	}
	cur := mll.head.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val, nil
}

func (mll *MyLinkList) RemoveByIndex(index int) (interface{}, error) {
	if mll.IsEmpty() {
		return nil, errors.New("this linkList is empty")
	}
	err := mll.checkIndex(index)
	if err != nil {
		return nil, err
	}
	prev := mll.head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	delNode := prev.Next
	prev.Next = prev.Next.Next
	delNode.Next = nil
	mll.size -= 1
	return delNode.Val, nil
}

func (mll *MyLinkList) RemoveFrist() (interface{}, error) {
	return mll.RemoveByIndex(0)
}

func (mll *MyLinkList) RemoveLast() (interface{}, error) {
	return mll.RemoveByIndex(mll.size-1)
}

func MyLinkListDemo() {
	mll := NewMyLinkList()
	val, err := mll.GetByIndex(0)
	fmt.Println(mll, val, err)
	for i := 1; i <= 3; i++ {
		mll.AddFirst(i)
	}
	fmt.Println(mll)
	err = mll.AddByIndex(2, 4)
	fmt.Println(mll, err)
	val, err = mll.RemoveByIndex(1)
	fmt.Println(mll, val, err)
	val, err = mll.RemoveByIndex(2)
	fmt.Println(mll, val, err)
}

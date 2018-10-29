package queue

import (
	"DataStruct/linkList"
	"fmt"
)

type LinkListQueue struct {
	data *linkList.LinkListWithTail
}

func (llq LinkListQueue) String() string {
	return llq.data.String()
}

func NewLinkListQueue() *LinkListQueue {
	return &LinkListQueue{
		data: linkList.NewLinkListWithTail(),
	}
}

func (llq LinkListQueue) IsEmpty() bool {
	return llq.data.IsEmpty()
}

func (llq LinkListQueue) GetSize() int {
	return llq.data.GetSize()
}

func (llq LinkListQueue) Peek() (interface{}, error) {
	return llq.data.GetByIndex(0)
}

func (llq *LinkListQueue) EnQueue(val interface{}) {
	llq.data.AddLast(val)
}

func (llq *LinkListQueue) DeQueue() (interface{}, error) {
	return llq.data.RemoveFrist()
}

func LinkListQueueDemo() {
	lq := NewLinkListQueue()
	fmt.Println(lq.data, lq.GetSize())
	val, err := lq.Peek()
	fmt.Println(lq.data, lq.GetSize(), val, err)
	lq.EnQueue(1)
	fmt.Println(lq.data, lq.GetSize())
	lq.EnQueue(2)
	fmt.Println(lq.data, lq.GetSize())
	lq.EnQueue(3)
	fmt.Println(lq.data, lq.GetSize())
	lq.EnQueue(4)
	fmt.Println(lq.data, lq.GetSize())
	val, err = lq.Peek()
	fmt.Println(lq.data, lq.GetSize(), val, err)
	val, err = lq.DeQueue()
	fmt.Println(lq.data, lq.GetSize(), val, err)
	lq.EnQueue(5)
	fmt.Println(lq.data, lq.GetSize())
	val, err = lq.DeQueue()
	fmt.Println(lq.data, lq.GetSize(), val, err)
	lq.EnQueue(6)
	fmt.Println(lq.data, lq.GetSize())
	val, err = lq.DeQueue()
	fmt.Println(lq.data, lq.GetSize(), val, err)
	lq.EnQueue(7)
	fmt.Println(lq.data, lq.GetSize())
	lq.EnQueue(8)
	fmt.Println(lq.data, lq.GetSize())
	lq.EnQueue(9)
	fmt.Println(lq.data, lq.GetSize())
	lq.EnQueue(10)
	fmt.Println(lq.data, lq.GetSize())
}

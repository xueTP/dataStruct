package queue

import (
	"errors"
	"fmt"
)

type LoopQueue struct {
	top int
	tail int
	data []interface{}
}

func NewLoopQueue(cap int) *LoopQueue {
	if cap <= 2 {
		cap = 2
	}
	return &LoopQueue{
		data: make([]interface{}, cap),
	}
}

func (lq LoopQueue)IsEmpty() bool {
	if lq.top == lq.tail {
		return true
	}
	return false
}

func (lq LoopQueue)GetSize() int {
	var size int
	if lq.top < lq.tail {
		size = lq.tail - lq.top
	}else if lq.top > lq.tail {
		size = len(lq.data) - (lq.top - lq.tail)
	}
	return size
}

func (lq LoopQueue)Peek() (interface{}, error) {
	if lq.IsEmpty() {
		return nil, errors.New("this loopQueue is empty")
	}
	return lq.data[lq.top], nil
}

func (lq *LoopQueue) EnQueue(val interface{}) {
	lq.data[lq.tail] = val
	lq.tail = (lq.tail + 1) % len(lq.data)
	fmt.Println("***********", lq.tail, lq.top, lq.GetSize())
	if (lq.tail + 1) % len(lq.data) == lq.top {
		lq.resize((lq.GetSize() + 1) * 2)
	}
}

func (lq *LoopQueue) DeQueue() (interface{}, error) {
	if lq.IsEmpty() {
		return nil, errors.New("this loopQueue is empty")
	}
	res := lq.data[lq.top]
	lq.top = (lq.top + 1) % len(lq.data)
	if lq.GetSize() < len(lq.data)/4 {
		lq.resize(len(lq.data)/2)
	}
	return res, nil
}

func (lq *LoopQueue)resize(cap int) {
	data := make([]interface{}, cap)
	for i := 0; i < lq.GetSize(); i++ {
		data[i] = lq.data[(lq.top + i) % len(lq.data)]
	}
	lq.tail = lq.GetSize()
	lq.top = 0
	lq.data = data
}

func LoopQueueDemo() {
	lq := NewLoopQueue(1)
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

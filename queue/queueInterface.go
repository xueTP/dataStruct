package queue

type Queue interface {
	IsEmpty() bool
	GetSize() int
	Peek() (interface{}, error)
	EnQueue(val interface{})
	DeQueue() (interface{}, error)
}

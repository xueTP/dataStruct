package stacks

type stacksInterface interface {
	IsEmpty() bool
	GetSize() int
	Push(val interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}

package linkList

type linkListInterface interface {
	IsEmpty() bool
	GetSize() int
	AddByIndex(index int, val interface{}) error
	AddFirst(val interface{}) error
	AddLast(val interface{}) error
	GetByIndex(index int) (interface{}, error)
	RemoveByIndex(index int) (interface{}, error)
	RemoveFrist() (interface{}, error)
	RemoveLast() (interface{}, error)
}

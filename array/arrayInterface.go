package array

type arrayInterface interface {
	IsEmpty() bool
	GetCap() int
	GetSize() int
	AddByIndex(index int, val interface{}) error
	AddAtLast(val interface{}) error
	AddAtHead(val interface{}) error
	GetByIndex(index int) (interface{}, error)
	GetIndexByValL(val interface{}) (int, error)
	GetIndexByValR(val interface{}) (int, error)
	DelByIndex(index int) (interface{}, error)
	DelByValL(val interface{}) (int, error)
	DelByValR(val interface{}) (int, error)
	ReplaceByIndex(index int, val interface{}) error
}

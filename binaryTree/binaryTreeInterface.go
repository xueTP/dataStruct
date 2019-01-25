package binaryTree

type BinaryTreeInterface interface {
	IsEmpty() bool
	GetSize() int
	AddNode(val Compared)
	DelMin() (interface{}, error)
	DelMax() (interface{}, error)
	DelNode(val Compared) error
	FindNode(val Compared) bool
}

type BinaryTreeToMapInterface interface {
	IsEmpty() bool
	GetSize() int
	AddNode(key Compared, val interface{})
	DelMin() (interface{}, interface{}, error)
	DelMax() (interface{}, interface{}, error)
	DelNode(key Compared) error
	FindNode(key Compared) bool
	GetVal(key Compared) interface{}
	SetNodeVal(key Compared, val interface{}) bool
}

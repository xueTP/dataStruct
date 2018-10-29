package heap

import "DataStruct/binaryTree"

/*
	this interface is regulation this Maximum heap has function
 */
type MaximumHeapInterface interface {
	IsEmpty() bool
	GetSize() int
	Get() interface{}
	Add(e binaryTree.Compared)
	Replace(e binaryTree.Compared) interface{}
	Heapify(es []binaryTree.Compared)
}

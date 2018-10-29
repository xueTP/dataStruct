package set

import "MyPro/dataStruct/binaryTree"

type SetInterface interface {
	IsEmpty() bool
	GetSize() int
	Find(e binaryTree.Compared) bool
	Add(e binaryTree.Compared)
	Remove(e binaryTree.Compared) bool
}

package mmap

import "MyPro/dataStruct/binaryTree"

type MmapInterface interface {
	IsEmpty() bool
	GetSize() int
	Set(k binaryTree.Compared, v interface{})
	Get(k binaryTree.Compared) interface{}
	Del(k binaryTree.Compared) interface{}
}

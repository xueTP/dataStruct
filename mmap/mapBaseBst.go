package mmap

import (
	"MyPro/dataStruct/binaryTree"
	"fmt"
)

type MapBaseBst struct {
	data *binaryTree.BinarySearchTreeToMap
}

func (mbb MapBaseBst) String() string {
	return "map bast in binarySearchTree: " + mbb.data.String()
}

func NewMapBaseBst() *MapBaseBst {
	return &MapBaseBst{
		data: binaryTree.NewBinarySearchTreeToMap(),
	}
}

func (mbb MapBaseBst) IsEmpty() bool {
	return mbb.data.IsEmpty()
}

func (mbb MapBaseBst) GetSize() int {
	return mbb.data.GetSize()
}

func (mbb *MapBaseBst) Set(k binaryTree.Compared, v interface{}) {
	if mbb.data.FindNode(k) {
		mbb.data.SetNodeVal(k, v)
	}else {
		mbb.data.AddNode(k, v)
	}
}

func (mbb *MapBaseBst) Get(k binaryTree.Compared) interface{} {
	return mbb.data.GetVal(k)
}

func (mbb *MapBaseBst) Del(k binaryTree.Compared) interface{} {
	val := mbb.data.GetVal(k)
	if val == nil {
		return nil
	}
	mbb.data.DelNode(k)
	return val
}

func MapBaseBstDemo() {
	mbb := NewMapBaseBst()
	mbb.Set(inter(12), 2)
	fmt.Println(mbb)
	mbb.Set(inter(6), 12)
	fmt.Println(mbb)
	mbb.Set(inter(35), 3)
	fmt.Println(mbb)
	val := mbb.Get(inter(6))
	fmt.Println(mbb, val)
	val = mbb.Get(inter(34))
	fmt.Println(mbb, val)
	val = mbb.Del(inter(12))
	fmt.Println(mbb, val)
	val = mbb.Del(inter(23))
	fmt.Println(mbb, val)
}

type inter int

func (i inter) Comparison(val interface{}) int {
	v, ok := val.(inter)
	if ok && i > v {
		return 1
	} else if ok && i < v {
		return -1
	}
	return 0
}


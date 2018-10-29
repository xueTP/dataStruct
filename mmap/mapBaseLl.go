package mmap

import (
	"MyPro/dataStruct/linkList"
	"MyPro/dataStruct/binaryTree"
	"fmt"
)

type MapBaseLl struct {
	data *linkList.LinkListToMap
}

func (mbl MapBaseLl) String() string {
	return "this map base linkList: " + mbl.data.ToString()
}

func NewMapBaseLl() *MapBaseLl {
	return &MapBaseLl{
		data: linkList.NewLinkListToMap(),
	}
}

func (mbl MapBaseLl) IsEmpty() bool {
	return mbl.data.IsEmpty()
}

func (mbl MapBaseLl) GetSize() int {
	return mbl.data.GetSize()
}

func (mbl *MapBaseLl) Set(k binaryTree.Compared, v interface{}) {
	mbl.data.SetByKey(k, v)
}

func (mbl MapBaseLl) Get(k binaryTree.Compared) interface{} {
	return mbl.data.GetByKey(k)
}

func (mbl *MapBaseLl) Del(k binaryTree.Compared) interface{} {
	return mbl.data.RemoveByKey(k)
}

func MapBaseLlDemo() {
	mbl := NewMapBaseLl()
	mbl.Set(inter(12), 2)
	fmt.Println(mbl)
	mbl.Set(inter(6), 12)
	fmt.Println(mbl)
	mbl.Set(inter(35), 3)
	fmt.Println(mbl)
	val := mbl.Get(inter(6))
	fmt.Println(mbl, val)
	val = mbl.Get(inter(34))
	fmt.Println(mbl, val)
	val = mbl.Del(inter(12))
	fmt.Println(mbl, val)
	val = mbl.Del(inter(23))
	fmt.Println(mbl, val)
}

package set

import (
	"MyPro/dataStruct/binaryTree"
	"log"
	"fmt"
)

type SetBaseBst struct {
	data *binaryTree.BinarySearchTree
}

func (sbst SetBaseBst) String() string {
	return "this set is base for binarySearchTree: " + sbst.data.String()
}

func NewSetBaseBst() *SetBaseBst {
	return &SetBaseBst{
		data: binaryTree.NewBinarySearchTree(),
	}
}

func (sbst SetBaseBst) IsEmpty() bool {
	return sbst.data.IsEmpty()
}

func (sbst SetBaseBst) GetSize() int {
	return sbst.data.GetSize()
}

func (sbst SetBaseBst) Find(e binaryTree.Compared) bool {
	return sbst.data.FindNode(e)
}

func (sbst SetBaseBst) Add(e binaryTree.Compared) {
	sbst.data.AddNode(e)
}

func (sbst SetBaseBst) Remove(e binaryTree.Compared) bool {
	err := sbst.data.DelNode(e)
	if err != nil {
		log.Printf("error : %v", err)
		return false
	}
	return true
}

func SetBaseBstDemo() {
	sbst := NewSetBaseBst()
	bool := sbst.Find(inter(1))
	fmt.Println(sbst, bool)
	sbst.Add(inter(1))
	sbst.Add(inter(2))
	sbst.Add(inter(1))
	fmt.Println(sbst)
	bool = sbst.Remove(inter(1))
	fmt.Println(sbst, bool)
	sbst.Add(inter(1))
	fmt.Println(sbst)
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

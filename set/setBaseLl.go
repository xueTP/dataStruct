package set

import (
	"MyPro/dataStruct/linkList"
	"fmt"
	"log"
	"MyPro/dataStruct/binaryTree"
)

type SetBaseLl struct {
	data *linkList.RecursionLinkList
}

func (sbl SetBaseLl) String() string {
	return "set base for Linklist " + sbl.data.ToString()
}

func NewSetBaseLl() *SetBaseLl {
	return &SetBaseLl{
		data: linkList.NewRecursionLinkList(),
	}
}

func (sbl SetBaseLl) IsEmpty() bool {
	return sbl.data.IsEmpty()
}

func (sbl SetBaseLl) GetSize() int {
	return sbl.data.GetSize()
}

func (sbl SetBaseLl) Find(e binaryTree.Compared) bool {
	return sbl.data.Find(e)
}

func (sbl *SetBaseLl) Add(e binaryTree.Compared) {
	if !sbl.data.Find(e) {
		err := sbl.data.AddFirst(e)
		if err != nil {
			log.Printf("error: %v", err)
		}
	}
}

func (sbl *SetBaseLl) Remove(e binaryTree.Compared) bool {
	return sbl.data.RemoveByVal(e)
}

func SetBaseLlDemo() {
	sbst := NewSetBaseLl()
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

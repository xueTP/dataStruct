package unionFind

import (
	"dataStruct/util"
	"fmt"
	"strings"
)

type UnionFindBySlice struct {
	data    map[interface{}]int
	mapping []int
	rang    []int
	size    int
}

func (ufbs UnionFindBySlice) String() string {
	var s string
	for i, v := range ufbs.data {
		s += util.InterfaceToString(i) + "(" + util.InterfaceToString(v) + ")" + "=>" + util.InterfaceToString(ufbs.mapping[v]) + ","
	}
	s = strings.Trim(s, ",")
	return s
}

func NewUnionFindBySlice(param []interface{}) *UnionFindBySlice {
	obj := &UnionFindBySlice{size: len(param) - 1}
	data := make(map[interface{}]int, len(param))
	mapping := make([]int, len(param))
	for i, v := range param {
		data[v] = i
		mapping[i] = i
	}
	obj.rang = make([]int, len(param))
	obj.data = data
	obj.mapping = mapping
	return obj
}

func (ufbs UnionFindBySlice) GetSize() int {
	return ufbs.size
}

func (ufbs UnionFindBySlice) find(index int) int {
	for ufbs.mapping[index] != index {
		if ufbs.mapping[index] != index {
			ufbs.mapping[index] = ufbs.mapping[ufbs.mapping[index]]
		}
		index = ufbs.mapping[index]
	}
	return index
}

func (ufbs UnionFindBySlice) checkParam(param interface{}) int {
	if _, ok := ufbs.data[param]; !ok {
		ufbs.size++
		ufbs.data[param] = ufbs.size
		ufbs.mapping = append(ufbs.mapping, ufbs.size)
		ufbs.rang = append(ufbs.rang, 0)
	}
	return ufbs.data[param]
}

func (ufbs *UnionFindBySlice) Union(a, b interface{}) {
	av := ufbs.checkParam(a)
	bv := ufbs.checkParam(b)
	aroot := ufbs.find(av)
	broot := ufbs.find(bv)
	if ufbs.rang[aroot] > ufbs.rang[broot] {
		ufbs.mapping[broot] = aroot
	} else if ufbs.rang[broot] > ufbs.rang[aroot] {
		ufbs.mapping[aroot] = broot
	} else {
		ufbs.rang[aroot]++
		ufbs.mapping[aroot] = broot
	}
}

func (ufbs *UnionFindBySlice) isConn(a, b interface{}) bool {
	av := ufbs.checkParam(a)
	bv := ufbs.checkParam(b)
	aroot := ufbs.find(av)
	broot := ufbs.find(bv)
	return aroot == broot
}

func UnionFindBySliceDemo() {
	ufbs := NewUnionFindBySlice([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(ufbs)
	bo := ufbs.isConn(1, 3)
	fmt.Println(ufbs, bo)
	ufbs.Union(4, 6)
	fmt.Println(ufbs)
	ufbs.Union(4, 7)
	fmt.Println(ufbs)
	ufbs.Union(4, 1)
	fmt.Println(ufbs)
}

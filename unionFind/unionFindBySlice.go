package unionFind

import (
	"dataStruct/util"
	"strings"
)

type UnionFindBySlice struct {
	data map[interface{}]int
	mapping []int
	size int
}

func (ufbs UnionFindBySlice) String() string {
	var s string
	for i, v := range ufbs.data {
		s += util.InterfaceToString(i) + "=>" + util.InterfaceToString(ufbs.mapping[v]) + ","
	}
	s = strings.Trim(s, ",")
	return s
}

func NewUnionFindBySlice(param []interface{}) *UnionFindBySlice {
	obj := &UnionFindBySlice{size: len(param) - 1}
	data := make(map[interface{}]int, len(param))
	mapping := make([]int, len(param))
	for i , v := range param {
		data[v] = i
		mapping[i] = i
	}
	obj.data = data
	obj.mapping = mapping
	return obj
}

func (ufbs UnionFindBySlice) GetSize() int{
	return ufbs.size
}

func (ufbs UnionFindBySlice) find(index int) (int, int) {
	var count int
	for ufbs.mapping[index] != index {
		count ++
		index = ufbs.mapping[index]
	}
	return index, count
}

func (ufbs UnionFindBySlice) checkParam(param interface{}) int {
	if _, ok := ufbs.data[param]; !ok {
		ufbs.size ++
		ufbs.data[param] = ufbs.size
		ufbs.mapping = append(ufbs.mapping, ufbs.size)
	}
	return ufbs.data[param]
}

func (ufbs *UnionFindBySlice) Union(a, b interface{}) {
	av := ufbs.checkParam(a)
	bv := ufbs.checkParam(b)
	aroot, ac := ufbs.find(av)
	broot, bc := ufbs.find(bv)
	if ac > bc {
		ufbs.mapping[broot] = aroot
	}else {
		ufbs.mapping[aroot] = broot
	}
}

func (ufbs *UnionFindBySlice) isConn(a, b interface{}) bool {
	av := ufbs.checkParam(a)
	bv := ufbs.checkParam(b)
	aroot, _ := ufbs.find(av)
	broot, _ := ufbs.find(bv)
	return aroot == broot
}

func UnionFindBySliceDemo() {

}

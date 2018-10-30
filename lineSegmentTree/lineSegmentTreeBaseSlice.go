package lineSegmentTree

import "fmt"

// LineSegmentTreeBaseSlice 线段树结构，基于golang的切片实现
type LineSegmentTreeBaseSlice struct {
	data []interface{}
	segmentTree []interface{}
	margin func(a, b interface{}) interface{}
}

// overwrite interface Stringer , to user by fmt customize output
func (lst LineSegmentTreeBaseSlice) String() string {
	return fmt.Sprintf("this data: %v, \n segmentTree: %v", lst.data, lst.segmentTree)
}

// Line Segment Tree 线段树的初始化方法，放回一个LineSegmentTreeBaseSlice 的对象
func NewLSTBaseSlice(param []interface{}, fun func(a, b interface{}) interface{}) *LineSegmentTreeBaseSlice {
	obj := &LineSegmentTreeBaseSlice{
		data: param,
		margin: fun,
	}
	seg := make([]interface{}, 4*len(param))
	obj.getSegTree(0, 0, len(obj.data) - 1, &seg)
	obj.segmentTree = seg
	return obj
}

// 获取给定下标的左节点index, 公式 i * 2 + 1
func (lst LineSegmentTreeBaseSlice) getLeftChild(x int) int {
	return x * 2 + 1
}

// 获取给定下标的右节点index, 公式 i * 2 + 2
func (lst LineSegmentTreeBaseSlice) getRightChild(x int) int {
	return x * 2 + 2
}

// 主要用于lineSegmentTree 初始化内部存储的树结构时调用
func (lst LineSegmentTreeBaseSlice) getSegTree(index, l, r int, seg *[]interface{}) {
	if l == r {
		(*seg)[index] = lst.data[l]
		return
	}

	mid := l + (r - l) / 2
	leftIndex := lst.getLeftChild(index)
	lst.getSegTree(leftIndex, l, mid, seg)
	rightIndex := lst.getRightChild(index)
	lst.getSegTree(rightIndex, mid + 1, r, seg)
	(*seg)[index] = lst.margin((*seg)[leftIndex], (*seg)[rightIndex])
	return
}

func (lst LineSegmentTreeBaseSlice) QueryInterval(l, r int) interface{} {
	if l < 0 || r > lst.GetSize() - 1 || r < l {
		return nil
	}

	return lst.queryInterval(0, 0, lst.GetSize(), l, r)
}

func (lst LineSegmentTreeBaseSlice) queryInterval(index, l, r , queryL, queryR int) interface{} {
	if l == queryL && r == queryR {
		return lst.segmentTree[index]
	}
	if r == l {
		return lst.data[r]
	}

	var rval, lval interface{}
	mid := l + (r - l) / 2
	leftIndex := lst.getLeftChild(index)
	rightIndex := lst.getRightChild(index)
	if queryR <= mid {
		lval = lst.queryInterval(leftIndex, l, mid, queryL, queryR)
	}else if queryL > mid {
		rval = lst.queryInterval(rightIndex, mid + 1, r, queryL, queryR)
	}else {
		lval = lst.queryInterval(leftIndex, l, mid, queryL, mid)
		rval = lst.queryInterval(rightIndex, mid + 1, r, mid + 1, queryR)
	}
	return lst.margin(lval, rval)
}

// 获取线段树内的成员个数
func (lst LineSegmentTreeBaseSlice) GetSize() int {
	return len(lst.data) - 1
}

func LSTBaseSliceDemo() {
	fun := func(a, b interface{}) interface{}{
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		return a.(int) + b.(int)
	}
	lst := NewLSTBaseSlice([]interface{}{1,2,3,45,6,47,34,3,3,4}, fun)
	fmt.Println(lst)
	val := lst.QueryInterval(1, 6)
	fmt.Println(lst, val)
}
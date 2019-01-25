package lineSegmentTree

import (
	"errors"
	"fmt"
)

// LineSegmentTreeBaseSlice 线段树结构，基于golang的切片实现
type LineSegmentTreeBaseSlice struct {
	data        []interface{}
	segmentTree []interface{}
	margin      func(a, b interface{}) interface{}
}

// overwrite interface Stringer , to user by fmt customize output
func (lst LineSegmentTreeBaseSlice) String() string {
	return fmt.Sprintf("this data: %v, \n segmentTree: %v", lst.data, lst.segmentTree)
}

// Line Segment Tree 线段树的初始化方法，放回一个LineSegmentTreeBaseSlice 的对象
func NewLSTBaseSlice(param []interface{}, fun func(a, b interface{}) interface{}) *LineSegmentTreeBaseSlice {
	obj := &LineSegmentTreeBaseSlice{
		data:   param,
		margin: fun,
	}
	seg := make([]interface{}, 4*len(param))
	obj.getSegTree(0, 0, len(obj.data)-1, &seg)
	obj.segmentTree = seg
	return obj
}

// 获取给定下标的左节点index, 公式 i * 2 + 1
func (lst LineSegmentTreeBaseSlice) getLeftChild(x int) int {
	return x*2 + 1
}

// 获取给定下标的右节点index, 公式 i * 2 + 2
func (lst LineSegmentTreeBaseSlice) getRightChild(x int) int {
	return x*2 + 2
}

// 主要用于lineSegmentTree 初始化内部存储的树结构时调用
func (lst LineSegmentTreeBaseSlice) getSegTree(index, l, r int, seg *[]interface{}) {
	if l == r {
		(*seg)[index] = lst.data[l]
		return
	}

	mid := l + (r-l)/2
	leftIndex := lst.getLeftChild(index)
	lst.getSegTree(leftIndex, l, mid, seg)
	rightIndex := lst.getRightChild(index)
	lst.getSegTree(rightIndex, mid+1, r, seg)
	(*seg)[index] = lst.margin((*seg)[leftIndex], (*seg)[rightIndex])
	return
}

// 获取一个区间内的数据操作值， 主要通过递归获取，这个值的操作定义为 LineSegmentTreeBaseSlice 的 margin 方法
func (lst LineSegmentTreeBaseSlice) QueryInterval(l, r int) interface{} {
	if l < 0 || r > lst.GetSize()-1 || r < l {
		return nil
	}

	return lst.queryInterval(0, 0, lst.GetSize(), l, r)
}

// 递归获取区间内的值， demo: 4-9 => 4-5 + (6-8 + 9)
func (lst LineSegmentTreeBaseSlice) queryInterval(index, l, r, queryL, queryR int) interface{} {
	if l == queryL && r == queryR {
		return lst.segmentTree[index]
	}
	if r == l {
		return lst.data[r]
	}

	var rval, lval interface{}
	mid := l + (r-l)/2
	leftIndex := lst.getLeftChild(index)
	rightIndex := lst.getRightChild(index)
	if queryR <= mid {
		lval = lst.queryInterval(leftIndex, l, mid, queryL, queryR)
	} else if queryL > mid {
		rval = lst.queryInterval(rightIndex, mid+1, r, queryL, queryR)
	} else {
		lval = lst.queryInterval(leftIndex, l, mid, queryL, mid)
		rval = lst.queryInterval(rightIndex, mid+1, r, mid+1, queryR)
	}
	return lst.margin(lval, rval)
}

// 获取线段树内的成员个数
func (lst LineSegmentTreeBaseSlice) GetSize() int {
	return len(lst.data) - 1
}

// 对于线段树内部的存储数据进行更新，同时维护线段树结构
func (lst LineSegmentTreeBaseSlice) Update(index int, val interface{}) error {
	if index < 0 || index > lst.GetSize() {
		return errors.New("this index is overflow")
	}

	lst.data[index] = val
	lst.update(0, 0, lst.GetSize(), index, val)
	return nil
}

// 递归对index数据变动造成的线段树结构进行维护
func (lst LineSegmentTreeBaseSlice) update(root, l, r, index int, val interface{}) {
	if l == r {
		lst.segmentTree[root] = val
		return
	}

	mid := l + (r-l)/2
	leftIndex := lst.getLeftChild(root)
	rightIndex := lst.getRightChild(root)
	if index >= mid+1 {
		lst.update(rightIndex, mid+1, l, index, val)
	} else {
		lst.update(leftIndex, l, mid, index, val)
	}
	lst.segmentTree[root] = lst.margin(lst.segmentTree[leftIndex], lst.segmentTree[rightIndex])
}

func LSTBaseSliceDemo() {
	fun := func(a, b interface{}) interface{} {
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		return a.(int) + b.(int)
	}
	lst := NewLSTBaseSlice([]interface{}{1, 2, 3, 45, 6, 47, 34, 3, 3, 4}, fun)
	fmt.Println(lst)
	val := lst.QueryInterval(1, 6)
	fmt.Println(lst, val)
	err := lst.Update(2, 23)
	fmt.Println(lst, err)
}

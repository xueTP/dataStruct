package heap

import (
	"DataStruct/binaryTree"
	"math"
	"fmt"
	"strconv"
)

type MaximumHeap struct {
	data []binaryTree.Compared
	size int
}

func NewMaximumHeap() *MaximumHeap {
	return &MaximumHeap{
		data: []binaryTree.Compared{},
	}
}

/*
	继承 Stringer 的string方法，fmt时会走这个方法
 */
func (mh MaximumHeap) String() string {
	return "this maximum heap data: " + fmt.Sprintf("%v", mh.data) + "size :" + strconv.Itoa(mh.GetSize())
}

/*
	获取给定下标的父节点index, 公式 (i - 1)/2 向下取整
 */
func (mh MaximumHeap) getParent(x int) int {
	return int(math.Floor(float64(x - 1) / 2))
}

/*
	获取给定下标的左节点index, 公式 i * 2 + 1
 */
func (mh MaximumHeap) getLeftChild(x int) int {
	return x * 2 + 1
}

/*
	获取给定下标的右节点index, 公式 i * 2 + 2
 */
func (mh MaximumHeap) getRightChild(x int) int {
	return x * 2 + 2
}

/*
	数据上浮 最大堆内部方法，主要是对于传入的数据位置依次向上浮动
	直到父节点大于数据或到顶了
 */
func (mh *MaximumHeap) siftUp(x int) {
	var nextX int
	for x > 0 {
		nextX = mh.getParent(x)
		if mh.data[x].Comparison(mh.data[nextX]) > 0 {
			mh.data[x], mh.data[nextX] = mh.data[nextX], mh.data[x]
			x = nextX
		}else {
			break
		}
	}
}

/*
	数据下沉 最大堆内部方法，主要是对于传入的数据位置依次向下浮动
	直到左右子节点都小于数据或到没有孩子节点了
 */
func (mh *MaximumHeap) siftDown(x int) {
	var max, r int
	for mh.getLeftChild(x) <= mh.GetSize() - 1 {
		max, r = mh.getLeftChild(x), mh.getRightChild(x)
		// fmt.Println(x, max, r)
		if r <= mh.GetSize() - 1 {
			if mh.data[r].Comparison(mh.data[max]) > 0 {
				max = r
			}
		}
		if mh.data[x].Comparison(mh.data[max]) < 0 {
			mh.data[x], mh.data[max] = mh.data[max], mh.data[x]
			x = max
		}else {
			break
		}
	}
}

/*
	返回当前最大堆是否为空，为空返回true
*/
func (mh MaximumHeap) IsEmpty() bool {
	return mh.size == 0
}

/*
	返回这个最大堆结构的目前堆内数据个数
*/
func (mh MaximumHeap) GetSize() int {
	return mh.size
}

/*
	获取最大堆中的最大成员，index 为0的数据并取最后一个数进行下沉操作
	保证最大堆的每个父节点都大于孩子节点
 */
func (mh *MaximumHeap) Get() interface{} {
	if mh.IsEmpty() {
		return nil
	}
	res := mh.data[0]
	mh.data[0] = mh.data[mh.GetSize() - 1]
	mh.size --
	mh.siftDown(0)
	mh.data = mh.data[:mh.GetSize()]
	return res
}

/*
	对最大堆添加成员，直接在数据最后添加，对当前数据最后的一位，index 为
	size - 1, 进行上浮操作保证最大堆的每个父节点都大于孩子节点
 */
func (mh *MaximumHeap) Add(e binaryTree.Compared) {
	mh.data = append(mh.data, e)
	mh.size ++
	mh.siftUp(mh.GetSize() - 1)
}

/*
	获取堆顶成员及添加成员，通过获取 0 位置的数据同时将参数放入 0 ，进行下
	沉操作
 */
func (mh *MaximumHeap) Replace(e binaryTree.Compared) interface{} {
	res := mh.data[0]
	mh.data[0] = e
	mh.siftDown(0)
	return res
}

/*
	一次性将多个数据放入，找到最后一个根节点开始下沉，直到堆顶
 */
func (mh *MaximumHeap) Heapify(es []binaryTree.Compared) {
	mh.data = append(mh.data, es...)
	mh.size = len(mh.data)
	for i := mh.getParent(mh.GetSize()-1); i >= 0; i-- {
		mh.siftDown(i)
	}
}

func MaximumHeapDemo() {
	mh := NewMaximumHeap()
	mh.Add(inter(45))
	mh.Add(inter(23))
	mh.Add(inter(33))
	mh.Add(inter(55))
	mh.Add(inter(1))
	mh.Add(inter(44))
	fmt.Println(mh)
	v := mh.Get()
	fmt.Println(mh, v)
	v = mh.Get()
	fmt.Println(mh, v)
	v = mh.Get()
	fmt.Println(mh, v)
	v = mh.Replace(inter(2))
	fmt.Println(mh, v)
	mh.Heapify([]binaryTree.Compared{inter(24), inter(222), inter(34), inter(55), inter(1)})
	fmt.Println(mh)
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
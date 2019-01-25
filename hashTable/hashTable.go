package hashTable

import (
	"dataStruct/binaryTree"
)

const maxRestSize = 9
const minRestSize = 1

type HashTable struct {
	size int
	data []binaryTree.BinarySearchTreeToMap
	hashKey int
	hashSeed []int
}

func NewHashTable() *HashTable {
	return &HashTable{
		hashKey: 0,
		data: make([]binaryTree.BinarySearchTreeToMap, 17),
		hashSeed: []int{17, 37, 79, 163, 331, 673, 1361, 2729, 5471, 10949, 21911, 43853, 87719, 175447, 350899, 701819, 1403641, 2807303, 5614657, 11229331, 22458671, 44917381, 89834777, 179669557, 359339171, 718678369, 1437356741, 2147483647},
	}
}

func (ht *HashTable) getHashCode(v interface{}) int64 {
	return GetHashCode(v) % int64(ht.hashSeed[ht.hashKey])
}

func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

func (ht *HashTable) resetSize(newKey int) {
	if newKey < 0 {
		newKey = 0
	}else if newKey > len(ht.hashSeed) - 1 {
		newKey = len(ht.hashSeed) - 1
	}
	// newData := make([]binaryTree.BinarySearchTreeToMap, newKey)
	// oldKey := ht.hashKey
	// ht.hashKey = newKey
	// for k, list := range ht.data {
	// 	for j, v := range list {
	//
	// 	}
	// }
	// TODO 树结构体 自建结构体没有循环获取内部数据导致无法resetSize
}

func (ht *HashTable) Set(k binaryTree.Compared, v interface{}) {
	hashCode := ht.getHashCode(k)
	ht.data[hashCode].AddNode(k, v)
	ht.size ++
}

func (ht *HashTable) Get(k binaryTree.Compared) interface{} {
	hashCode := ht.getHashCode(k)
	return ht.data[hashCode].GetVal(k)
}

func (ht *HashTable) Del(k binaryTree.Compared) error {
	hashCode := ht.getHashCode(k)
	return ht.data[hashCode].DelNode(k)
}





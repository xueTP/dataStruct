package redBlackTree

import "dataStruct/binaryTree"

const Red = true
const Black = false

// redBlackTreeNode 红黑树基础节点， 主要是拥有 color 进行界定红还是黑
type redBlackTreeNode struct {
	key   binaryTree.Compared
	val   interface{}
	color bool
	left  *redBlackTreeNode
	right *redBlackTreeNode
}

// NewRedBlackNode 实例化红黑树节点，默认创建红节点
func NewRedBlackNode(k binaryTree.Compared, v interface{}) *redBlackTreeNode {
	return &redBlackTreeNode{
		key:   k,
		val:   v,
		color: Red,
	}
}

func NewRedBlackEmptyNode() *redBlackTreeNode {
	return nil
}

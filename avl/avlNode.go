package avl

import "dataStruct/binaryTree"

// 给 avl树节点， 带有key val high
type AvlTreeNode struct {
	Val         interface{}
	Key         interface{}
	High        int
	Left, Right *AvlTreeNode
}

func NewAvlNodeEmpty() *AvlTreeNode {
	return nil
}

func NewAvlNodeOnly(k binaryTree.Compared, v interface{}) *AvlTreeNode {
	return &AvlTreeNode{Key: k, Val: v}
}

func NewAvlNode(k binaryTree.Compared, v interface{}, l, r *AvlTreeNode) *AvlTreeNode {
	return &AvlTreeNode{
		Val:   v,
		Key:   k,
		Right: r,
		Left:  l,
	}
}

package binaryTree

type TreeNode struct {
	Val         interface{}
	Left, Right *TreeNode
}

func NewTreeNodeEmpty() *TreeNode {
	return nil
}

func NewTreeNodeOnly(v Compared) *TreeNode {
	return &TreeNode{Val: v}
}

func NewTreeNode(v Compared, l, r *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   v,
		Right: r,
		Left:  l,
	}
}

/**
 * 使用这个二叉搜索树的结构需要实现这个比较接口
 * > 0 ==> this > beComp
 * < 0 ==> this < beComp
 * = 0 ==> this == beComp
 */
type Compared interface {
	Comparison(beComp interface{}) int
}

// 给 map 使用的二叉搜索树节点， 带有key val
type TreeNodeToMap struct {
	Val         interface{}
	Key         interface{}
	Left, Right *TreeNodeToMap
}

func NewTreeNodeToMEmpty() *TreeNodeToMap {
	return nil
}

func NewTreeNodeToMOnly(k Compared, v interface{}) *TreeNodeToMap {
	return &TreeNodeToMap{Key: k, Val: v}
}

func NewTreeNodeToM(k Compared, v interface{}, l, r *TreeNodeToMap) *TreeNodeToMap {
	return &TreeNodeToMap{
		Val:   v,
		Key:   k,
		Right: r,
		Left:  l,
	}
}

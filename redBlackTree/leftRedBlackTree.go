package redBlackTree

import (
	"dataStruct/binaryTree"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// LeftRBTree 红黑树内部实现的结构体，本次实现主要是基础的红节点左倾
type LeftRBTree struct {
	size int
	node *redBlackTreeNode // 红黑树根节点
}

// NewLeftRBTree 实例化左倾红黑树对象
func NewLeftRBTree() *LeftRBTree {
	return &LeftRBTree{
		node: NewRedBlackEmptyNode(),
	}
}

func (rbt LeftRBTree) String() string {
	return "tree in_order_traversal is [" + strings.Trim(rbt.inOrderTraversal(rbt.node, ""), ",") + "] size is " + strconv.Itoa(rbt.GetSize())
}

// inOrderTraversal 中序遍历红黑树内部存储数据
func (rbt *LeftRBTree) inOrderTraversal(node *redBlackTreeNode, s string) string {
	if node == nil {
		return s
	}
	s = rbt.inOrderTraversal(node.left, s)
	s += fmt.Sprintf("%v-%v", node.key, node.val) + "|"
	if node.color {
		s += "红,"
	} else {
		s += "黑,"
	}
	s = rbt.inOrderTraversal(node.right, s)
	return s
}

// isRed 辅助函数，用于判断传入节点是否为红节点
func (rbt *LeftRBTree) isRed(node *redBlackTreeNode) bool {
	return node.color == Red
}

// flipColor 辅助函数，用于当添加节点在当前为三节点的右边，根节点为黑，左节点为红，右节点为红
func (rbt *LeftRBTree) flipColor(node *redBlackTreeNode) {
	node.color = Red
	node.left.color = Black
	node.right.color = Black
}

// leftRotate 辅助函数左旋，用于当节点添加在二节点右边时使用
func (rbt *LeftRBTree) leftRotate(node *redBlackTreeNode) *redBlackTreeNode {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = Red
	return x
}

func (rbt *LeftRBTree) addNode(val binaryTree.Compared, Node *redBlackTreeNode) *redBlackTreeNode {
	if Node == nil {
		rbt.size++
		return NewRedBlackNode(val, nil)
	}

	if val.Comparison(Node.val) > 0 {
		Node.right = rbt.addNode(val, Node.right)
	} else if val.Comparison(Node.val) < 0 {
		Node.left = rbt.addNode(val, Node.left)
	}
	return Node
}

func (rbt LeftRBTree) IsEmpty() bool {
	return rbt.size == 0
}

func (rbt LeftRBTree) GetSize() int {
	return rbt.size
}

func (rbt *LeftRBTree) AddNode(val binaryTree.Compared) {
	rbt.node = rbt.addNode(val, rbt.node)
}

func (rbt *LeftRBTree) delMin(node *redBlackTreeNode) (*redBlackTreeNode, interface{}) {
	if node.left == nil {
		return node.right, node.val
	}
	var v interface{}
	node.left, v = rbt.delMin(node.left)
	return node, v
}

func (rbt *LeftRBTree) DelMin() (interface{}, error) {
	if rbt.IsEmpty() {
		return nil, errors.New("this binarySearchTree is empty")
	}
	var res interface{}
	rbt.node, res = rbt.delMin(rbt.node)
	rbt.size--
	return res, nil
}

func (rbt *LeftRBTree) delMax(node *redBlackTreeNode) (*redBlackTreeNode, interface{}) {
	if node.right == nil {
		return node.left, node.val
	}
	var v interface{}
	node.right, v = rbt.delMax(node.right)
	return node, v
}

func (rbt *LeftRBTree) DelMax() (interface{}, error) {
	if rbt.IsEmpty() {
		return nil, errors.New("this binarySearchTree is empty")
	}
	var res interface{}
	rbt.node, res = rbt.delMax(rbt.node)
	rbt.size--
	return res, nil
}

func (rbt *LeftRBTree) delNode(node *redBlackTreeNode, val binaryTree.Compared) (*redBlackTreeNode, error) {
	if node == nil {
		return node, errors.New("this node is not in this LeftRBTree")
	}
	var err error
	if val.Comparison(node.val) > 0 {
		node.right, err = rbt.delNode(node.right, val)
	} else if val.Comparison(node.val) < 0 {
		node.left, err = rbt.delNode(node.left, val)
	} else {
		var newVal interface{}
		if node.right == nil {
			return nil, nil
		}
		node.right, newVal = rbt.delMin(node.right)
		node.val = newVal
	}
	return node, err
}

func (rbt *LeftRBTree) DelNode(val binaryTree.Compared) error {
	if rbt.IsEmpty() {
		return errors.New("this binarySearchTree is empty")
	}
	var err error
	rbt.node, err = rbt.delNode(rbt.node, val)
	rbt.size--
	return err
}

func (rbt *LeftRBTree) FindNode(val binaryTree.Compared) bool {
	tempNode := rbt.node
	for tempNode != nil {
		if val.Comparison(tempNode.val) > 0 {
			tempNode = tempNode.right
		} else if val.Comparison(tempNode.val) < 0 {
			tempNode = tempNode.left
		} else {
			return true
		}
	}
	return false
}

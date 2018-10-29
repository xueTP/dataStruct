package binaryTree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BinarySearchTree struct {
	size int
	node *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		node: NewTreeNodeEmpty(),
	}
}

func (bst *BinarySearchTree) String() string {
	return "tree in_order_traversal is [" + strings.Trim(bst.inOrderTraversal(bst.node, ""), ",") + "] size is " + strconv.Itoa(bst.GetSize())
}

func (bst *BinarySearchTree) inOrderTraversal(node *TreeNode, s string) string {
	if node == nil {
		return s
	}
	s = bst.inOrderTraversal(node.Left, s)
	s += fmt.Sprintf("%v", node.Val) + ","
	s = bst.inOrderTraversal(node.Right, s)
	return s
}

func (bst *BinarySearchTree) addNode(val Compared, Node *TreeNode) *TreeNode {
	if Node == nil {
		bst.size ++
		return NewTreeNodeOnly(val)
	}

	if val.Comparison(Node.Val) > 0 {
		Node.Right = bst.addNode(val, Node.Right)
	} else if val.Comparison(Node.Val) < 0 {
		Node.Left = bst.addNode(val, Node.Left)
	}
	return Node
}

func (bst BinarySearchTree) IsEmpty() bool {
	return bst.size == 0
}

func (bst BinarySearchTree) GetSize() int {
	return bst.size
}

func (bst *BinarySearchTree) AddNode(val Compared) {
	bst.node = bst.addNode(val, bst.node)
}

func (bst *BinarySearchTree) delMin(node *TreeNode) (*TreeNode, interface{}) {
	if node.Left == nil {
		return node.Right, node.Val
	}
	var v interface{}
	node.Left, v = bst.delMin(node.Left)
	return node, v
}

func (bst *BinarySearchTree) DelMin() (interface{}, error) {
	if bst.IsEmpty() {
		return nil, errors.New("this binarySearchTree is empty")
	}
	var res interface{}
	bst.node, res = bst.delMin(bst.node)
	bst.size--
	return res, nil
}

func (bst *BinarySearchTree) delMax(node *TreeNode) (*TreeNode, interface{}) {
	if node.Right == nil {
		return node.Left, node.Val
	}
	var v interface{}
	node.Right, v = bst.delMax(node.Right)
	return node, v
}

func (bst *BinarySearchTree) DelMax() (interface{}, error) {
	if bst.IsEmpty() {
		return nil, errors.New("this binarySearchTree is empty")
	}
	var res interface{}
	bst.node, res = bst.delMax(bst.node)
	bst.size--
	return res, nil
}

func (bst *BinarySearchTree) delNode(node *TreeNode, val Compared) (*TreeNode, error) {
	if node == nil {
		return node, errors.New("this node is not in this BinarySearchTree")
	}
	var err error
	if val.Comparison(node.Val) > 0 {
		node.Right, err = bst.delNode(node.Right, val)
	}else if val.Comparison(node.Val) < 0 {
		node.Left, err = bst.delNode(node.Left, val)
	}else {
		var newVal interface{}
		if node.Right == nil {
			return nil, nil
		}
		node.Right, newVal = bst.delMin(node.Right)
		node.Val = newVal
	}
	return node, err
}

func (bst *BinarySearchTree) DelNode(val Compared) error {
	if bst.IsEmpty() {
		return errors.New("this binarySearchTree is empty")
	}
	var err error
	bst.node, err =  bst.delNode(bst.node, val)
	bst.size--
	return err
}

func (bst *BinarySearchTree) FindNode(val Compared) bool {
	tempNode := bst.node
	for tempNode != nil {
		if val.Comparison(tempNode.Val) > 0 {
			tempNode = tempNode.Right
		}else if val.Comparison(tempNode.Val) < 0 {
			tempNode = tempNode.Left
		}else {
			return true
		}
	}
	return false
}

func BinarySearchTreeDemo() {
	bst := NewBinarySearchTree()
	fmt.Println(bst)
	bst.AddNode(inter(12))
	bst.AddNode(inter(23))
	bst.AddNode(inter(4))
	fmt.Println(bst)
	bst.AddNode(inter(1))
	bst.AddNode(inter(7))
	bst.AddNode(inter(2))
	bst.AddNode(inter(9))
	bst.AddNode(inter(15))
	fmt.Println(bst)
	val, err := bst.DelMax()
	fmt.Println(bst, "--res", val, err)
	err = bst.DelNode(inter(4))
	fmt.Println(bst, "--res", err)
	bool := bst.FindNode(inter(15))
	fmt.Println(bst, "--res", bool)
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

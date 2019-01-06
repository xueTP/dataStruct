package binaryTree

import (
	"dataStruct/util"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BinarySearchTreeToMap struct {
	size int
	node *TreeNodeToMap
}

func NewBinarySearchTreeToMap() *BinarySearchTreeToMap {
	return &BinarySearchTreeToMap{
		node: NewTreeNodeToMEmpty(),
	}
}

func (bst *BinarySearchTreeToMap) String() string {
	return "tree in_order_traversal is [" + strings.Trim(bst.inOrderTraversal(bst.node, ""), ",") + "] size is " + strconv.Itoa(bst.GetSize())
}

func (bst *BinarySearchTreeToMap) inOrderTraversal(node *TreeNodeToMap, s string) string {
	if node == nil {
		return s
	}
	s = bst.inOrderTraversal(node.Left, s)
	s += fmt.Sprintf("%v", node.Key) + ":" + util.InterfaceToString(node.Val) + ","
	s = bst.inOrderTraversal(node.Right, s)
	return s
}

func (bst BinarySearchTreeToMap) IsEmpty() bool {
	return bst.size == 0
}

func (bst BinarySearchTreeToMap) GetSize() int {
	return bst.size
}

func (bst *BinarySearchTreeToMap) addNode(key Compared, val interface{}, Node *TreeNodeToMap) *TreeNodeToMap {
	if Node == nil {
		bst.size++
		return NewTreeNodeToMOnly(key, val)
	}

	if key.Comparison(Node.Key) > 0 {
		Node.Right = bst.addNode(key, val, Node.Right)
	} else if key.Comparison(Node.Key) < 0 {
		Node.Left = bst.addNode(key, val, Node.Left)
	}

	return Node
}

func (bst *BinarySearchTreeToMap) AddNode(key Compared, val interface{}) {
	bst.node = bst.addNode(key, val, bst.node)
}

func (bst *BinarySearchTreeToMap) delMin(node *TreeNodeToMap) (*TreeNodeToMap, interface{}, interface{}) {
	if node.Left == nil {
		return node.Right, node.Key, node.Val
	}
	var v, k interface{}
	node.Left, k, v = bst.delMin(node.Left)
	return node, k, v
}

func (bst *BinarySearchTreeToMap) DelMin() (interface{}, interface{}, error) {
	if bst.IsEmpty() {
		return nil, nil, errors.New("this binarySearchTree is empty")
	}
	var val, key interface{}
	bst.node, key, val = bst.delMin(bst.node)
	bst.size--
	return val, key, nil
}

func (bst *BinarySearchTreeToMap) delMax(node *TreeNodeToMap) (*TreeNodeToMap, interface{}, interface{}) {
	if node.Right == nil {
		return node.Left, node.Key, node.Val
	}
	var v, k interface{}
	node.Right, k, v = bst.delMax(node.Right)
	return node, k, v
}

func (bst *BinarySearchTreeToMap) DelMax() (interface{}, interface{}, error) {
	if bst.IsEmpty() {
		return nil, nil, errors.New("this binarySearchTree is empty")
	}
	var val, key interface{}
	bst.node, key, val = bst.delMax(bst.node)
	bst.size--
	return val, key, nil
}

func (bst *BinarySearchTreeToMap) delNode(node *TreeNodeToMap, key Compared) (*TreeNodeToMap, error) {
	if node == nil {
		return node, errors.New("this node is not in this BinarySearchTree")
	}
	var err error
	if key.Comparison(node.Key) > 0 {
		node.Right, err = bst.delNode(node.Right, key)
	} else if key.Comparison(node.Key) < 0 {
		node.Left, err = bst.delNode(node.Left, key)
	} else {
		var newVal, newKey interface{}
		if node.Right == nil {
			return nil, nil
		}
		node.Right, newKey, newVal = bst.delMin(node.Right)
		node.Val = newVal
		node.Key = newKey
	}
	return node, err
}

func (bst *BinarySearchTreeToMap) DelNode(key Compared) error {
	if bst.IsEmpty() {
		return errors.New("this binarySearchTree is empty")
	}
	var err error
	bst.node, err = bst.delNode(bst.node, key)
	bst.size--
	return err
}

func (bst BinarySearchTreeToMap) FindNode(key Compared) bool {
	return bst.GetVal(key) != nil
}

func (bst BinarySearchTreeToMap) GetVal(key Compared) interface{} {
	tempNode := bst.node
	for tempNode != nil {
		if key.Comparison(tempNode.Key) > 0 {
			tempNode = tempNode.Right
		} else if key.Comparison(tempNode.Key) < 0 {
			tempNode = tempNode.Left
		} else {
			return tempNode.Val
		}
	}
	return nil
}

func (bst *BinarySearchTreeToMap) SetNodeVal(key Compared, val interface{}) bool {
	tempNode := bst.node
	for tempNode != nil {
		if key.Comparison(tempNode.Val) > 0 {
			tempNode = tempNode.Right
		} else if key.Comparison(tempNode.Val) < 0 {
			tempNode = tempNode.Left
		} else {
			tempNode.Val = val
			return true
		}
	}
	return false
}

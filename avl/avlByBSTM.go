package avl

import (
	"dataStruct/binaryTree"
	"dataStruct/util"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type AvlByBSTM struct {
	size int
	node *AvlTreeNode
}

func NewAvlByBSTM() *AvlByBSTM {
	return &AvlByBSTM{
		node: NewAvlNodeEmpty(),
	}
}

func (abst *AvlByBSTM) String() string {
	return "tree in_order_traversal is [" + strings.Trim(abst.inOrderTraversal(abst.node, ""), ",") + "] size is " + strconv.Itoa(abst.GetSize())
}

// inOrderTraversal 中序遍历 avl 树,输出内部存储的结构
func (abst *AvlByBSTM) inOrderTraversal(node *AvlTreeNode, s string) string {
	if node == nil {
		return s
	}
	s += "("
	s = abst.inOrderTraversal(node.Left, s)
	s += ")"
	s += fmt.Sprintf("%v", node.Key) + ":" + util.InterfaceToString(abst.getHigh(node)) + ","
	s += "("
	s = abst.inOrderTraversal(node.Right, s)
	s += ")"
	return s
}

// IsEmpty 返回当前调用 avl 对象是否为空
func (abst AvlByBSTM) IsEmpty() bool {
	return abst.size == 0
}

// GetSize 返回整个 avl 的节点个数
func (abst AvlByBSTM) GetSize() int {
	return abst.size
}

// getHigh 获取传入节点的当前高度
func (abst AvlByBSTM) getHigh(node *AvlTreeNode) int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(abst.getHigh(node.Left)), float64(abst.getHigh(node.Right)))) + 1
}

// getBalanceFactor 返回传入分支的平衡因子 (左节点高度 - 右节点高度)
func (abst AvlByBSTM) getBalanceFactor(node *AvlTreeNode) int {
	return abst.getHigh(node.Left) - abst.getHigh(node.Right)
}

// ggRotate avl 辅助方法，主要用来维持 avl 平衡
func (abst *AvlByBSTM) ggRotate(node *AvlTreeNode) *AvlTreeNode {
	centerNode := node.Left
	tempy := centerNode.Right
	node.Left = tempy
	// node.High = abst.getHigh(node)
	centerNode.Right = node
	// centerNode.High = abst.getHigh(centerNode)
	return centerNode
}

func (abst *AvlByBSTM) llRotate(node *AvlTreeNode) *AvlTreeNode {
	centerNode := node.Right
	tempy := centerNode.Left
	node.Right = tempy
	centerNode.Left = node
	return centerNode
}

// addNode 通过递归向 avl 中添加节点（已经存在则修改节点对应val）
func (abst *AvlByBSTM) addNode(key binaryTree.Compared, val interface{}, Node *AvlTreeNode) *AvlTreeNode {
	if Node == nil {
		abst.size++
		return NewAvlNodeOnly(key, val)
	}

	if key.Comparison(Node.Key) > 0 {
		Node.Right = abst.addNode(key, val, Node.Right)
	} else if key.Comparison(Node.Key) < 0 {
		Node.Left = abst.addNode(key, val, Node.Left)
	} else {
		Node.Val = val
	}
	// 获取当前的高度(左右子节点的最大高度+1)
	Node.High = abst.getHigh(Node)
	// 获取当前节点的平衡因子
	balanceFactor := abst.getBalanceFactor(Node)
	if balanceFactor > 1 && abst.getBalanceFactor(Node.Left) >= 0 {
		return abst.ggRotate(Node)
	}
	if balanceFactor > 1 && abst.getBalanceFactor(Node.Left) < 0 {
		Node.Left = abst.llRotate(Node.Left)
		return abst.ggRotate(Node)
	}
	if balanceFactor < -1 && abst.getBalanceFactor(Node.Right) <= 0 {
		return abst.llRotate(Node)
	}
	if balanceFactor < -1 && abst.getBalanceFactor(Node.Right) > 0 {
		Node.Right = abst.ggRotate(Node.Right)
		return abst.llRotate(Node)
	}

	return Node
}

// AddNode 主要调用 addNode 实现添加节点
func (abst *AvlByBSTM) AddNode(key binaryTree.Compared, val interface{}) {
	abst.node = abst.addNode(key, val, abst.node)
}

func (abst *AvlByBSTM) delMin(node *AvlTreeNode) (*AvlTreeNode, interface{}, interface{}) {
	if node.Left == nil {
		return node.Right, node.Key, node.Val
	}
	var v, k interface{}
	node.Left, k, v = abst.delMin(node.Left)
	return node, k, v
}

func (abst *AvlByBSTM) DelMin() (interface{}, interface{}, error) {
	if abst.IsEmpty() {
		return nil, nil, errors.New("this binarySearchTree is empty")
	}
	var val, key interface{}
	abst.node, key, val = abst.delMin(abst.node)
	abst.size--
	return val, key, nil
}

func (abst *AvlByBSTM) delMax(node *AvlTreeNode) (*AvlTreeNode, interface{}, interface{}) {
	if node.Right == nil {
		return node.Left, node.Key, node.Val
	}
	var v, k interface{}
	node.Right, k, v = abst.delMax(node.Right)
	return node, k, v
}

func (abst *AvlByBSTM) DelMax() (interface{}, interface{}, error) {
	if abst.IsEmpty() {
		return nil, nil, errors.New("this binarySearchTree is empty")
	}
	var val, key interface{}
	abst.node, key, val = abst.delMax(abst.node)
	abst.size--
	return val, key, nil
}

func (abst *AvlByBSTM) delNode(node *AvlTreeNode, key binaryTree.Compared) (*AvlTreeNode, error) {
	if node == nil {
		return node, errors.New("this node is not in this BinarySearchTree")
	}
	var err error
	if key.Comparison(node.Key) > 0 {
		node.Right, err = abst.delNode(node.Right, key)
	} else if key.Comparison(node.Key) < 0 {
		node.Left, err = abst.delNode(node.Left, key)
	} else {
		var newVal, newKey interface{}
		if node.Right == nil {
			return nil, nil
		}
		node.Right, newKey, newVal = abst.delMin(node.Right)
		node.Val = newVal
		node.Key = newKey
	}

	// 获取当前的高度(左右子节点的最大高度+1)
	node.High = abst.getHigh(node)
	// 获取当前节点的平衡因子
	balanceFactor := abst.getBalanceFactor(node)
	if balanceFactor > 1 && abst.getBalanceFactor(node.Left) >= 0 {
		return abst.ggRotate(node), err
	}
	if balanceFactor > 1 && abst.getBalanceFactor(node.Left) < 0 {
		node.Left = abst.llRotate(node.Left)
		return abst.ggRotate(node), err
	}
	if balanceFactor < -1 && abst.getBalanceFactor(node.Right) <= 0 {
		return abst.llRotate(node), err
	}
	if balanceFactor < -1 && abst.getBalanceFactor(node.Right) > 0 {
		node.Right = abst.ggRotate(node.Right)
		return abst.llRotate(node), err
	}

	return node, err
}

func (abst *AvlByBSTM) DelNode(key binaryTree.Compared) error {
	if abst.IsEmpty() {
		return errors.New("this binarySearchTree is empty")
	}
	var err error
	abst.node, err = abst.delNode(abst.node, key)
	abst.size--
	return err
}

func (abst AvlByBSTM) FindNode(key binaryTree.Compared) bool {
	return abst.GetVal(key) != nil
}

func (abst AvlByBSTM) GetVal(key binaryTree.Compared) interface{} {
	tempNode := abst.node
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

func (abst *AvlByBSTM) SetNodeVal(key binaryTree.Compared, val interface{}) bool {
	tempNode := abst.node
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

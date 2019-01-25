package redBlackTree

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

// 模拟一个自定义类型用来测试 Compared
type comparedInter int

func (this comparedInter) Comparison(beComp interface{}) int {
	beCompVal, ok := beComp.(comparedInter)
	if !ok {
		logrus.Errorf("this beComp is error not implement Compared")
	}
	return int(this - beCompVal)
}

func TestLeftRBTree_leftRotate(t *testing.T) {
	rbtNode := NewRedBlackNode(comparedInter(1), nil)
	rbtNode.color = Black
	rbtNode.right = NewRedBlackNode(comparedInter(3), nil)
	rbt := &LeftRBTree{node: rbtNode, size: 2}
	logrus.Infoln(rbt)
	rbt.node = rbt.leftRotate(rbt.node)
	logrus.Infoln(rbt)
}

func TestLeftRBTree_rightRotate(t *testing.T) {
	rbtNode := NewRedBlackNode(comparedInter(3), nil)
	rbtNode.color = Black
	rbtNode.left = NewRedBlackNode(comparedInter(2), nil)
	rbtNode.left.left = NewRedBlackNode(comparedInter(1), nil)
	rbt := &LeftRBTree{node: rbtNode, size: 2}
	logrus.Infoln(rbt)
	rbt.node = rbt.rightRotate(rbt.node)
	logrus.Infoln(rbt)
}

func TestLeftRBTree_addNode(t *testing.T) {
	rbt := NewLeftRBTree()
	rbt.node = rbt.addNode(comparedInter(4), rbt.node)
	logrus.Info(rbt)
	rbt.node = rbt.addNode(comparedInter(2), rbt.node)
	logrus.Info(rbt)
	rbt.node = rbt.addNode(comparedInter(3), rbt.node)
	logrus.Info(rbt)
	rbt.node = rbt.addNode(comparedInter(1), rbt.node)
	logrus.Infof("root node is %v", rbt.node)
	logrus.Info(rbt)
}

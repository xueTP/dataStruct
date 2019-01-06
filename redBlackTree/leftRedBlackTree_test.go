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

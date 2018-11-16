package avl

import (
	"testing"
	"github.com/Sirupsen/logrus"
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

func TestAvlByBSTM_ggRotate(t *testing.T) {
	testNode := NewAvlNode(
		comparedInter(4), nil, NewAvlNode(
			comparedInter(3), nil, NewAvlNodeOnly(
				comparedInter(1), nil),
			nil),
		nil)
	avl := &AvlByBSTM{
		node: testNode,
		size: 3,
	}
	t.Logf("%#v, %v", testNode, avl.getBalanceFactor(testNode))
	resNode := avl.ggRotate(testNode)
	if avl.getBalanceFactor(resNode) > 1 {
		t.Error("this func is fail")
	}
	t.Logf("%#v", resNode)
}

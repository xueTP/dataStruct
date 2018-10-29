package linkList

type Node struct {
	Val interface{}
	Next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{Val: val, Next: nil}
}

func NewNodeWithNext(val interface{}, next *Node) *Node {
	return &Node{Val: val, Next: next}
}


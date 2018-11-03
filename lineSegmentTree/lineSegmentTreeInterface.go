package lineSegmentTree

type LineSegmentTreeInterface interface {
	GetSize() int
	QueryInterval(l, r int) interface{}
	Update(index int, val interface{}) error
}

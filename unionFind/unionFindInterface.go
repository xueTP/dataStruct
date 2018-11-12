package unionFind

type UnionFindInterface interface {
	GetSize() int
	Union(a, b interface{})
	isConn(a, b interface{}) bool
}

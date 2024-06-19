package types

type Comparable interface {
	//	equal
	EQ(Comparable) bool
	//	not equal
	NE(Comparable) bool
	//	less than
	LT(Comparable) bool
	//	less than or equal
	LE(Comparable) bool
	//	greater than
	GT(Comparable) bool
	//	greater than or equal
	GE(Comparable) bool
}

type Comparator[T Comparable] interface {
	Compare(T, T) int
}

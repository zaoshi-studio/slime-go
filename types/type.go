package types

type Ordered interface {
	Integer | Float | Complex | ~string
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

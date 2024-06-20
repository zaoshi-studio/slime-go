package types

type Uint64 struct {
	value uint64
}

func (v1 *Uint64) EQ(v Comparable) bool {
	v2, ok := v.(*Uint64)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Uint64) NE(v Comparable) bool {
	v2, ok := v.(*Uint64)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Uint64) LT(v Comparable) bool {
	v2, ok := v.(*Uint64)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Uint64) LE(v Comparable) bool {
	v2, ok := v.(*Uint64)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Uint64) GT(v Comparable) bool {
	v2, ok := v.(*Uint64)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Uint64) GE(v Comparable) bool {
	v2, ok := v.(*Uint64)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

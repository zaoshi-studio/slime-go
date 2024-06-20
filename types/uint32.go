package types

type Uint32 struct {
	value uint
}

func (v1 *Uint32) EQ(v Comparable) bool {
	v2, ok := v.(*Uint32)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Uint32) NE(v Comparable) bool {
	v2, ok := v.(*Uint32)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Uint32) LT(v Comparable) bool {
	v2, ok := v.(*Uint32)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Uint32) LE(v Comparable) bool {
	v2, ok := v.(*Uint32)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Uint32) GT(v Comparable) bool {
	v2, ok := v.(*Uint32)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Uint32) GE(v Comparable) bool {
	v2, ok := v.(*Uint32)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

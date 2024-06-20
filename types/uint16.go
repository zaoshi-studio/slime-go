package types

type Uint16 struct {
	value uint
}

func (v1 *Uint16) EQ(v Comparable) bool {
	v2, ok := v.(*Uint16)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Uint16) NE(v Comparable) bool {
	v2, ok := v.(*Uint16)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Uint16) LT(v Comparable) bool {
	v2, ok := v.(*Uint16)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Uint16) LE(v Comparable) bool {
	v2, ok := v.(*Uint16)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Uint16) GT(v Comparable) bool {
	v2, ok := v.(*Uint16)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Uint16) GE(v Comparable) bool {
	v2, ok := v.(*Uint16)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

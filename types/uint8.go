package types

type Uint8 struct {
	value uint8
}

func (v1 *Uint8) EQ(v Comparable) bool {
	v2, ok := v.(*Uint8)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Uint8) NE(v Comparable) bool {
	v2, ok := v.(*Uint8)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Uint8) LT(v Comparable) bool {
	v2, ok := v.(*Uint8)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Uint8) LE(v Comparable) bool {
	v2, ok := v.(*Uint8)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Uint8) GT(v Comparable) bool {
	v2, ok := v.(*Uint8)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Uint8) GE(v Comparable) bool {
	v2, ok := v.(*Uint8)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

package types

type Uint struct {
	value uint
}

func (v1 *Uint) EQ(v Comparable) bool {
	v2, ok := v.(*Uint)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Uint) NE(v Comparable) bool {
	v2, ok := v.(*Uint)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Uint) LT(v Comparable) bool {
	v2, ok := v.(*Uint)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Uint) LE(v Comparable) bool {
	v2, ok := v.(*Uint)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Uint) GT(v Comparable) bool {
	v2, ok := v.(*Uint)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Uint) GE(v Comparable) bool {
	v2, ok := v.(*Uint)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

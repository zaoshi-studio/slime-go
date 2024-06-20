package types

type Int8 struct {
	value int8
}

func (v1 *Int8) EQ(v Comparable) bool {
	v2, ok := v.(*Int8)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Int8) NE(v Comparable) bool {
	v2, ok := v.(*Int8)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Int8) LT(v Comparable) bool {
	v2, ok := v.(*Int8)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Int8) LE(v Comparable) bool {
	v2, ok := v.(*Int8)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Int8) GT(v Comparable) bool {
	v2, ok := v.(*Int8)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Int8) GE(v Comparable) bool {
	v2, ok := v.(*Int8)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

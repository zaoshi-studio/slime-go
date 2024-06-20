package types

type Int16 struct {
	value int16
}

func (v1 *Int16) EQ(v Comparable) bool {
	v2, ok := v.(*Int16)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Int16) NE(v Comparable) bool {
	v2, ok := v.(*Int16)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Int16) LT(v Comparable) bool {
	v2, ok := v.(*Int16)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Int16) LE(v Comparable) bool {
	v2, ok := v.(*Int16)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Int16) GT(v Comparable) bool {
	v2, ok := v.(*Int16)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Int16) GE(v Comparable) bool {
	v2, ok := v.(*Int16)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

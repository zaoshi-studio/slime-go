package types

type Int32 struct {
	value int32
}

func (v1 *Int32) EQ(v Comparable) bool {
	v2, ok := v.(*Int32)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Int32) NE(v Comparable) bool {
	v2, ok := v.(*Int32)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Int32) LT(v Comparable) bool {
	v2, ok := v.(*Int32)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Int32) LE(v Comparable) bool {
	v2, ok := v.(*Int32)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Int32) GT(v Comparable) bool {
	v2, ok := v.(*Int32)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Int32) GE(v Comparable) bool {
	v2, ok := v.(*Int32)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

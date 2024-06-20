package types

type Int64 struct {
	value int64
}

func (v1 *Int64) EQ(v Comparable) bool {
	v2, ok := v.(*Int64)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Int64) NE(v Comparable) bool {
	v2, ok := v.(*Int64)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Int64) LT(v Comparable) bool {
	v2, ok := v.(*Int64)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Int64) LE(v Comparable) bool {
	v2, ok := v.(*Int64)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Int64) GT(v Comparable) bool {
	v2, ok := v.(*Int64)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Int64) GE(v Comparable) bool {
	v2, ok := v.(*Int64)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

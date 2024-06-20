package types

type Int struct {
	value int
}

func (v1 *Int) EQ(v Comparable) bool {
	v2, ok := v.(*Int)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Int) NE(v Comparable) bool {
	v2, ok := v.(*Int)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Int) LT(v Comparable) bool {
	v2, ok := v.(*Int)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Int) LE(v Comparable) bool {
	v2, ok := v.(*Int)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Int) GT(v Comparable) bool {
	v2, ok := v.(*Int)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Int) GE(v Comparable) bool {
	v2, ok := v.(*Int)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

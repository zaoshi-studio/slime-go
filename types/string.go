package types

type String struct {
	value string
}

func (v1 *String) EQ(v Comparable) bool {
	v2, ok := v.(*String)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *String) NE(v Comparable) bool {
	v2, ok := v.(*String)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *String) LT(v Comparable) bool {
	v2, ok := v.(*String)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *String) LE(v Comparable) bool {
	v2, ok := v.(*String)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *String) GT(v Comparable) bool {
	v2, ok := v.(*String)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *String) GE(v Comparable) bool {
	v2, ok := v.(*String)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

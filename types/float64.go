package types

type Float64 struct {
	value float64
}

func (v1 *Float64) EQ(v Comparable) bool {
	v2, ok := v.(*Float64)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Float64) NE(v Comparable) bool {
	v2, ok := v.(*Float64)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Float64) LT(v Comparable) bool {
	v2, ok := v.(*Float64)
	if !ok {
		return false
	}
	return v1.value < v2.value
}

func (v1 *Float64) LE(v Comparable) bool {
	v2, ok := v.(*Float64)
	if !ok {
		return false
	}
	return v1.value <= v2.value
}

func (v1 *Float64) GT(v Comparable) bool {
	v2, ok := v.(*Float64)
	if !ok {
		return false
	}
	return v1.value > v2.value
}

func (v1 *Float64) GE(v Comparable) bool {
	v2, ok := v.(*Float64)
	if !ok {
		return false
	}
	return v1.value >= v2.value
}

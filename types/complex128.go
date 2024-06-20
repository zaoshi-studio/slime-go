package types

type Complex128 struct {
	value complex128
}

func (v1 *Complex128) EQ(v Comparable) bool {
	v2, ok := v.(*Complex128)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Complex128) NE(v Comparable) bool {
	v2, ok := v.(*Complex128)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Complex128) LT(v Comparable) bool {
	v2, ok := v.(*Complex128)
	if !ok {
		return false
	}
	return real(v1.value) < real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) < imag(v2.value))
}

func (v1 *Complex128) LE(v Comparable) bool {
	v2, ok := v.(*Complex128)
	if !ok {
		return false
	}
	return real(v1.value) < real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) <= imag(v2.value))
}

func (v1 *Complex128) GT(v Comparable) bool {
	v2, ok := v.(*Complex128)
	if !ok {
		return false
	}
	return real(v1.value) > real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) > imag(v2.value))
}

func (v1 *Complex128) GE(v Comparable) bool {
	v2, ok := v.(*Complex128)
	if !ok {
		return false
	}
	return real(v1.value) > real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) >= imag(v2.value))
}

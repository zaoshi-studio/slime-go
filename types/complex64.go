package types

type Complex64 struct {
	value complex64
}

func (v1 *Complex64) EQ(v Comparable) bool {
	v2, ok := v.(*Complex64)
	if !ok {
		return false
	}
	return v1.value == v2.value
}

func (v1 *Complex64) NE(v Comparable) bool {
	v2, ok := v.(*Complex64)
	if !ok {
		return false
	}
	return v1.value != v2.value
}

func (v1 *Complex64) LT(v Comparable) bool {
	v2, ok := v.(*Complex64)
	if !ok {
		return false
	}
	return real(v1.value) < real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) < imag(v2.value))
}

func (v1 *Complex64) LE(v Comparable) bool {
	v2, ok := v.(*Complex64)
	if !ok {
		return false
	}
	return real(v1.value) < real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) <= imag(v2.value))
}

func (v1 *Complex64) GT(v Comparable) bool {
	v2, ok := v.(*Complex64)
	if !ok {
		return false
	}
	return real(v1.value) > real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) > imag(v2.value))
}

func (v1 *Complex64) GE(v Comparable) bool {
	v2, ok := v.(*Complex64)
	if !ok {
		return false
	}
	return real(v1.value) > real(v2.value) || (real(v1.value) == real(v2.value) && imag(v1.value) >= imag(v2.value))
}

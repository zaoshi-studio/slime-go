package ringbuffer

type RingBuffer[T any] struct {
	data     []T
	size     int
	readIdx  int
	writeIdx int
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	buffer := &RingBuffer[T]{
		size: size,
		data: make([]T, 0, size),
	}
	return buffer
}

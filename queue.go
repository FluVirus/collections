package collections

type Queue[T any] struct {
	buf   []T
	read  int
	write int
	len   int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) resize(targetCapacity int) {
	if len(q.buf) >= targetCapacity {
		return
	}

	newBuf := make([]T, targetCapacity)
	newBufP := 0

	for i := q.read; i < q.len; i++ {
		newBuf[newBufP] = q.buf[i]
		newBufP++
	}
	q.read = 0

	for i := 0; i < q.write; i++ {
		newBuf[newBufP] = q.buf[i]
		newBufP++
	}

	q.write = len(q.buf)

	q.buf = newBuf
}

func (q *Queue[T]) desiredCap() int {
	var oldCap = len(q.buf)
	var newCap int

	switch {
	case oldCap == 0:
		newCap = 1
	case oldCap < 1024:
		newCap = 2 * oldCap
	default:
		newCap = 5 * oldCap / 4
	}

	return newCap
}

func (q *Queue[T]) Enqueue(value T) {
	// resize is necessary when:
	// 1. len(buf) = 0
	// 2. q.write == q.read && q.len > 0

	if (len(q.buf) == 0) || (q.write == q.read && q.len > 0) {
		q.resize(q.desiredCap())
	}

	q.buf[q.write] = value
	q.write = (q.write + 1) % len(q.buf)

	q.len++
}

func (q *Queue[T]) Dequeue() T {
	if q.len == 0 {
		panic("trying to dequeue from empty queue")
	}

	val := q.buf[q.read]
	q.read = (q.read + 1) % len(q.buf)

	q.len--

	return val
}

func (q *Queue[T]) Peek() T {
	if q.len == 0 {
		panic("trying to peek from empty queue")
	}

	return q.buf[q.read]
}

func (q *Queue[T]) Len() int {
	return q.len
}

func (q *Queue[T]) Cap() int {
	return len(q.buf)
}

func (q *Queue[T]) Grow(targetCapacity int) {
	if targetCapacity < 0 {
		panic("trying to grow from negative capacity")
	}

	q.resize(targetCapacity)
}

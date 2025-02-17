package collections

import (
	"errors"
)

// Queue implements queue in golang
type Queue[T any] struct {
	buf   []T
	first int
	next  int
}

// NewQueue method creates queue of type T and returns pointers to it
func NewQueue[T any]() *Queue[T] {
	var q = new(Queue[T])
	q.buf = make([]T, 0)
	q.first, q.next = -1, -1

	return q
}

// resize method must be called when first and next pointers are equal;
// position is the same value of next and first
func (q *Queue[T]) resize(targetCapacity int) {
	var newBuf []T
	var newBufP int

	if len(q.buf) >= targetCapacity {
		return
	}

	newBuf = make([]T, targetCapacity)
	newBufP = 0

	if q.first >= 0 {
		for i := q.first; i < len(q.buf); i++ {
			newBuf[newBufP] = q.buf[i]
			newBufP++
		}
		q.first = 0
	}

	for i := 0; i < q.next; i++ {
		newBuf[newBufP] = q.buf[i]
		newBufP++
	}

	q.next = len(q.buf)

	q.buf = newBuf
}

func (q *Queue[T]) desiredCap() int {
	var oldCap = cap(q.buf)
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

func (q *Queue[T]) Enqueue(val T) {
	if q.next == q.first {
		q.resize(q.desiredCap())
	}

	q.buf[q.next] = val
	if q.first < 0 {
		q.first = q.next
	}

	q.next = (q.next + 1) % len(q.buf)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.first < 0 {
		var zero T
		return zero, errors.New("trying to dequeue from empty queue")
	}

	val := q.buf[q.first]
	q.first = (q.first + 1) % len(q.buf)
	if q.first == q.next {
		q.first = -1
	}

	return val, nil
}

func (q *Queue[T]) Peak() (T, error) {
	if q.first < 0 {
		var zero T
		return zero, errors.New("trying to peek from empty queue")
	}

	return q.buf[q.first], nil
}

func (q *Queue[T]) Len() int {
	if q.first < 0 {
		return 0
	}

	diff := q.next - q.first

	if diff > 0 {
		return diff
	}

	return len(q.buf) + diff
}

func (q *Queue[T]) Cap() int {
	return len(q.buf)
}

func (q *Queue[T]) Grow(targetCapacity int) {
	q.resize(targetCapacity)
}

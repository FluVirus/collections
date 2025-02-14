package collections

type Heap[T any] struct {
	slice   []T
	compare func(T, T) int
}

func NewHeap[T any](compare func(T, T) int) *Heap[T] {
	heap := Heap[T]{
		slice:   make([]T, 0),
		compare: compare,
	}

	return &heap
}

func (h *Heap[T]) Push(value T) {
	h.slice = append(h.slice, value)
	h.heapifyUp(len(h.slice) - 1)
}

func (h *Heap[T]) Pop() T {
	if len(h.slice) == 0 {
		panic("pop from empty heap")
	}

	value := h.slice[0]
	h.slice[0] = h.slice[len(h.slice)-1]
	h.slice = h.slice[:len(h.slice)-1]

	if len(h.slice) > 0 {
		h.heapifyDown(0)
	}

	return value
}

func (h *Heap[T]) Peek() T {
	if len(h.slice) == 0 {
		panic("peek from empty heap")
	}

	return h.slice[0]
}

func (h *Heap[T]) Len() int {
	return len(h.slice)
}

func (h *Heap[T]) heapifyUp(index int) {
	for i, p := index, parent(index); p >= 0 && p < i; i, p = p, parent(p) {
		cmp := h.compare(h.slice[i], h.slice[p])
		if cmp < 0 {
			h.slice[i], h.slice[p] = h.slice[p], h.slice[i]
		}
	}
}

func (h *Heap[T]) heapifyDown(index int) {
	for {
		left := leftChildren(index)
		right := rightChildren(index)
		smallest := index

		if left < len(h.slice) && h.compare(h.slice[left], h.slice[smallest]) < 0 {
			smallest = left
		}

		if right < len(h.slice) && h.compare(h.slice[right], h.slice[smallest]) < 0 {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.slice[smallest], h.slice[index] = h.slice[index], h.slice[smallest]
		index = smallest
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChildren(index int) int {
	return 2*index + 1
}

func rightChildren(index int) int {
	return 2*index + 2
}

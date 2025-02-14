package collections

type LinkedList[T any] struct {
	Head  *LinkedListNode[T]
	Tail  *LinkedListNode[T]
	count int
}

type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
	Prev  *LinkedListNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	var linkedList LinkedList[T]
	return &linkedList
}

func (ll *LinkedList[T]) PushFront(value T) {
	var node LinkedListNode[T]
	node.Value = value

	if ll.Head == nil {
		ll.Tail = &node
	} else {
		node.Next = ll.Head
		ll.Head.Prev = &node
	}

	ll.Head = &node
	ll.count++
}

func (ll *LinkedList[T]) PushBack(value T) {
	var node LinkedListNode[T]
	node.Value = value

	if ll.Tail == nil {
		ll.Head = &node
	} else {
		node.Prev = ll.Tail
		ll.Tail.Next = &node
	}

	ll.Tail = &node
	ll.count++
}

func (ll *LinkedList[T]) PopFront() T {
	var value T

	if ll.Head == nil {
		panic("cannot remove from empty list")
	}

	value = ll.Head.Value

	if ll.Head == ll.Tail {
		ll.Head, ll.Tail = nil, nil
	} else {
		ll.Head.Next.Prev = nil
		ll.Head = ll.Head.Next
	}

	ll.count--

	return value
}

func (ll *LinkedList[T]) PopBack() T {
	var value T

	if ll.Head == nil {
		panic("cannot remove from empty list")
	}

	value = ll.Tail.Value

	if ll.Head == ll.Tail {
		ll.Head, ll.Tail = nil, nil
	} else {
		ll.Tail.Prev.Next = nil
		ll.Tail = ll.Tail.Prev
	}

	ll.count--

	return value
}

func (ll *LinkedList[T]) Len() int {
	return ll.count
}

func (ll *LinkedList[T]) Get(index int) T {
	if index < 0 || index >= ll.count {
		panic("index out of range")
	}

	var node = ll.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node.Value
}

func (ll *LinkedList[T]) Set(index int, value T) {
	if index < 0 || index >= ll.count {
		panic("index out of range")
	}

	var node = ll.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	node.Value = value
}

func (ll *LinkedList[T]) Remove(index int) T {
	if index < 0 || index >= ll.count {
		panic("index out of range")
	}

	var node = ll.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	if node == ll.Head && node == ll.Tail {
		ll.Head, ll.Tail = nil, nil
	} else if node == ll.Head {
		ll.Head = node.Next
		ll.Head.Prev = nil
	} else if node == ll.Tail {
		ll.Tail = node.Prev
		ll.Tail.Next = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}

	ll.count--

	return node.Value
}

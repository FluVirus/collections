package collections

import "errors"

type LinkedList[T any] struct {
	Head  *linkedListNode[T]
	Tail  *linkedListNode[T]
	count int
}

type linkedListNode[T any] struct {
	Value T
	Next  *linkedListNode[T]
	Prev  *linkedListNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	var linkedList LinkedList[T]
	return &linkedList
}

func (ll *LinkedList[T]) Prepend(value T) {
	var node linkedListNode[T]
	node.Value = value

	if ll.Head == nil {
		ll.Tail = &node
	} else {
		node.Next = ll.Head
	}

	ll.Head = &node
	ll.count++
}

func (ll *LinkedList[T]) Append(value T) {
	var node linkedListNode[T]
	node.Value = value

	if ll.Tail == nil {
		ll.Head = &node
	} else {
		ll.Tail.Next = &node
		node.Prev = ll.Tail
	}

	ll.Tail = &node
	ll.count++
}

func (ll *LinkedList[T]) RemoveFromHead() (T, error) {
	var value T

	if ll.Head == nil {
		return value, errors.New("cannot remove from empty list")
	}

	value = ll.Head.Value

	if ll.Head == ll.Tail {
		ll.Head, ll.Tail = nil, nil
	} else {
		ll.Head.Next.Prev = nil
		ll.Head = ll.Head.Next
	}

	ll.count--
	return value, nil
}

func (ll *LinkedList[T]) RemoveFromTail() (T, error) {
	var value T

	if ll.Head == nil {
		return value, errors.New("cannot remove from empty list")
	}

	value = ll.Tail.Value

	if ll.Head == ll.Tail {
		ll.Head, ll.Tail = nil, nil
	} else {
		ll.Tail.Prev.Next = nil
		ll.Tail = ll.Tail.Prev
	}

	ll.count--
	return value, nil
}

func (ll *LinkedList[T]) Len() int {
	return ll.count
}

func (ll *LinkedList[T]) Get(index int) (T, error) {
	var value T

	if index < 0 || index >= ll.count {
		return value, errors.New("index out of range")
	}

	var node = ll.Head
	for i := 0; i <= index; i++ {
		node = node.Next
	}

	return node.Value, nil
}

func (ll *LinkedList[T]) Set(index int, value T) error {
	if index < 0 || index >= ll.count {
		return errors.New("index out of range")
	}

	var node = ll.Head
	for i := 0; i <= index; i++ {
		node = node.Next
	}

	node.Value = value

	return nil
}

func (ll *LinkedList[T]) Remove(index int) (T, error) {
	var value T

	if index < 0 || index >= ll.count {
		return value, errors.New("index out of range")
	}

	var node = ll.Head
	for i := 0; i <= index; i++ {
		node = node.Next
	}

	value = node.Value

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

	return value, nil
}

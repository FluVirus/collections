package collections

import "errors"

type LinkedList[T any] struct {
	Head *LinkedListNode[T]
	Tail *LinkedListNode[T]
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

func (ll *LinkedList[T]) Prepend(value T) {
	var node LinkedListNode[T]
	node.Value = value

	if ll.Head == nil {
		ll.Tail = &node
	} else {
		node.Next = ll.Head
	}

	ll.Head = &node
}

func (ll *LinkedList[T]) Append(value T) {
	var node LinkedListNode[T]
	node.Value = value

	if ll.Tail == nil {
		ll.Head = &node
	} else {
		ll.Tail.Next = &node
		node.Prev = ll.Tail
	}

	ll.Tail = &node
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

	return value, nil
}

package queue

import dsa "github.com/nacknime-official/kata-machine-go/src/DSA"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Queue[T comparable] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func newNode[T comparable](item T) *Node[T] {
	return &Node[T]{value: item, next: nil}
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{Length: 0, head: nil, tail: nil}
}

func (q *Queue[T]) Enqueue(item T) {
	node := newNode(item)
	q.Length++
	if q.head == nil {
		q.head = node
		q.tail = node
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.head != nil {
		removed := q.head
		q.head = removed.next
		removed.next = nil
		q.Length--
		return removed.value, true
	}

	return dsa.ZeroValue[T](), false
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head != nil {
		return q.head.value, true
	}
	return dsa.ZeroValue[T](), false
}

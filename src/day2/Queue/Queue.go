package queue

import (
	"math"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func NewNode[T comparable](item T) *Node[T] {
	return &Node[T]{value: item, next: nil}
}

type Queue[T comparable] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{Length: 0, head: nil}
}

func (q *Queue[T]) Enqueue(item T) {
	node := NewNode(item)
	q.Length++
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.head != nil {
		removed := q.head
		q.head = removed.next
		removed.next = nil
		q.Length = int(math.Max(float64(q.Length-1), 0))
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

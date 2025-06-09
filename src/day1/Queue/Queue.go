package queue

import (
	"math"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

// func (n *Node[T]) print() {
// 	fmt.Printf("Node: %v, %v\n", n.value, n.next)
// }

type Queue[T comparable] struct {
	Length int
	head   *Node[T]
	tail   *Node[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{Length: 0, head: nil, tail: nil}

}

func (q *Queue[T]) Enqueue(item T) {
	q.Length++
	node := &Node[T]{value: item, next: nil}
	if q.tail != nil {
		q.tail.next = node
		q.tail = node
		return
	}
	q.head = node
	q.tail = node
}

func (q *Queue[T]) Deque() (T, bool) {
	if q.head != nil {
		q.Length = int(math.Max(0, float64(q.Length-1)))
		cur := q.head
		q.head = cur.next
		if q.head == nil {
			q.tail = q.head
		}
		cur.next = nil
		return cur.value, true
	}
	return dsa.ZeroValue[T](), false

}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head != nil {
		return q.head.value, true
	}
	return dsa.ZeroValue[T](), false

}

// func (q *Queue[T]) print() {
// 	fmt.Printf("Length: %d\n", q.Length)
// 	if q.head == nil {
// 		return
// 	}
// 	fmt.Print("[")
// 	current := q.head
// 	for current.next != nil {
// 		fmt.Printf("%v, ", current.value)
// 		current = current.next
// 	}
// 	fmt.Printf("%v]\n\n", current.value)
//
// }

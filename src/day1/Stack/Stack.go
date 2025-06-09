package stack

import (
	"math"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Stack[T comparable] struct {
	Length int
	head   *Node[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{Length: 0, head: nil}

}

func (s *Stack[T]) Push(item T) {
	s.Length++
	node := &Node[T]{value: item, next: nil}
	if s.head != nil {
		node.next = s.head
	}
	s.head = node
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.head != nil {
		s.Length = int(math.Max(0, float64(s.Length-1)))
		cur := s.head
		s.head = s.head.next
		cur.next = nil
		return cur.value, true
	}
	return dsa.ZeroValue[T](), false

}

func (s *Stack[T]) Peek() (T, bool) {
	if s.head != nil {
		return s.head.value, true
	}

	return dsa.ZeroValue[T](), false

}

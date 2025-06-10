package stack

import (
	"math"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	prev  *Node[T]
}

func NewNode[T comparable](item T) *Node[T] {
	return &Node[T]{value: item, prev: nil}
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
	node := NewNode(item)
	if s.head != nil {
		node.prev = s.head
	}
	s.head = node
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.head != nil {
		removed := s.head
		s.head = removed.prev
		removed.prev = nil
		s.Length = int(math.Max(float64(s.Length-1), 0))
		return removed.value, true
	}

	return dsa.ZeroValue[T](), false

}

func (s *Stack[T]) Peek() (T, bool) {
	if s.head != nil {
		return s.head.value, true
	}
	return dsa.ZeroValue[T](), false

}

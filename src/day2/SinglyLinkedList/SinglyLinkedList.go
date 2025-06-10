package singlylinkedlist

import (
	"math"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{value: val, next: nil}
}

type SinglyLinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (l *SinglyLinkedList[T]) Len() int { return l.length }

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{length: 0, head: nil}

}

func (l *SinglyLinkedList[T]) traverseTo(idx int) *Node[T] {
	if l.head == nil || idx == 0 {
		return l.head
	}

	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}

	return current
}

func (l *SinglyLinkedList[T]) Prepend(item T) {
	l.length++
	node := NewNode(item)

	if l.head != nil {
		node.next = l.head
	}
	if l.head == nil {
		l.tail = node
	}
	l.head = node
}

func (l *SinglyLinkedList[T]) InsertAt(item T, idx int) {
	if idx == 0 {
		l.Prepend(item)
		return
	}
	l.length++
	node := NewNode(item)
	current := l.traverseTo(idx - 1)
	node.next = current.next
	current.next = node
}

func (l *SinglyLinkedList[T]) Append(item T) {
	if l.tail == nil {
		l.Prepend(item)
		return
	}
	node := NewNode(item)
	l.length++
	l.tail.next = node
	l.tail = node
}

func (l *SinglyLinkedList[T]) Remove(item T) (T, bool) {
	current := l.head
	if current.value == item {
		l.length = int(math.Max(float64(l.length-1), 0))
		l.head = l.head.next
		return current.value, true
	}

	var newTail *Node[T]

	for current.next != nil {
		if current.next == l.tail {
			newTail = current
		}
		if current.next.value == item {
			l.length = int(math.Max(float64(l.length-1), 0))
			current.next = current.next.next
			return current.next.value, true

		}
		current = current.next
	}
	if newTail != nil {
		l.tail = newTail
	}

	return dsa.ZeroValue[T](), false
}

func (l *SinglyLinkedList[T]) Get(idx int) (T, bool) {
	node := l.traverseTo(idx)
	if node != nil {
		return node.value, true
	}

	return dsa.ZeroValue[T](), false
}

func (l *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {
	if idx == 0 {
		l.length = int(math.Max(float64(l.length-1), 0))
		removed := l.head
		l.head = removed.next
		removed.next = nil
		return removed.value, true
	}
	current := l.traverseTo(idx - 1)
	if current.next != nil {
		l.length = int(math.Max(float64(l.length-1), 0))
		removed := current.next
		current.next = removed.next
		removed.next = nil
		return removed.value, true
	} else {
		l.length = int(math.Max(float64(l.length-1), 0))
		newTail := l.traverseTo(idx - 2)
		removed := l.tail
		l.tail = newTail
		return removed.value, true
	}

	return dsa.ZeroValue[T](), false

}

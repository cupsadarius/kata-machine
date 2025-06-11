package singlylinkedlist

import (
	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func newNode[T comparable](item T) *Node[T] {
	return &Node[T]{value: item, next: nil}
}

type SinglyLinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (l *SinglyLinkedList[T]) Len() int { return l.length }

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{length: 0, head: nil, tail: nil}
}

func (l *SinglyLinkedList[T]) traverseTo(idx int) *Node[T] {
	if l.head == nil {
		return nil
	}
	current := l.head
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current
}

func (l *SinglyLinkedList[T]) Prepend(item T) {
	node := newNode(item)
	l.length++
	if l.head != nil {
		node.next = l.head
	}
	if l.tail == nil {
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
	node := newNode(item)
	current := l.traverseTo(idx - 1)
	node.next = current.next
	current.next = node

}

func (l *SinglyLinkedList[T]) Append(item T) {
	if l.tail == nil {
		l.Prepend(item)
		return
	}
	node := newNode(item)
	l.length++

	l.tail.next = node
	l.tail = node

}

func (l *SinglyLinkedList[T]) Remove(item T) (T, bool) {
	current := l.head
	if current.value == item {
		l.length--
		l.head = current.next
		current.next = nil
		return current.value, true
	}
	for current.next != nil {
		if current.next.value == item {
			l.length--
			if current.next == l.tail {
				l.tail = current
			}
			removed := current.next
			current.next = removed.next
			removed.next = nil
			return removed.value, true
		}

	}

	return dsa.ZeroValue[T](), false

}

func (l *SinglyLinkedList[T]) Get(idx int) (T, bool) {
	node := l.traverseTo(idx)
	if node != nil {
		return node.value, true
	}

	return dsa.ZeroValue[T](), true
}

func (l *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {
	if idx == 0 {
		removed := l.head
		l.head = removed.next
		removed.next = nil
		l.length--
		return removed.value, true
	}
	node := l.traverseTo(idx - 1)
	if node.next != nil {
		removed := node.next
		node.next = removed.next
		removed.next = nil
		l.length--
		return removed.value, true
	} else {

		l.length--
		removed := l.tail
		l.tail = l.traverseTo(idx - 2)
		removed.next = nil
		return removed.value, true
	}

	return dsa.ZeroValue[T](), false

}

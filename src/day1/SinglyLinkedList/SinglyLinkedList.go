package singlylinkedlist

import (
	"fmt"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func (n *Node[T]) print() {
	fmt.Printf("Node: %v, %v\n", n.value, n.next)
}

type SinglyLinkedList[T comparable] struct {
	length int
	head   *Node[T]
}

func (l *SinglyLinkedList[T]) Len() int { return l.length }

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{length: 0, head: nil}
}

func (l *SinglyLinkedList[T]) Prepend(item T) {
	l.length++

	node := Node[T]{value: item, next: nil}
	if l.head != nil {
		node.next = l.head
	}
	l.head = &node
}

func (l *SinglyLinkedList[T]) InsertAt(item T, idx int) {
	l.length++
	node := Node[T]{value: item, next: nil}
	if idx == 0 {
		node.next = l.head
		l.head = &node
		return
	}
	current := l.head
	for i := 0; i < idx-1; i++ {
		current = current.next
	}

	node.next = current.next
	current.next = &node

}

func (l *SinglyLinkedList[T]) Append(item T) {
	l.length++
	node := Node[T]{value: item, next: nil}
	if l.head == nil {
		l.head = &node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = &node
}

func (l *SinglyLinkedList[T]) Remove(item T) (T, bool) {
	current := l.head
	if current.value == item {
		l.head = l.head.next
		l.length--
		return current.value, true
	}
	for current.next != nil {
		if current.next.value == item {
			l.length--
			current.next = current.next.next
			return current.next.value, true

		}
		current = current.next
	}

	return dsa.ZeroValue[T](), false

}

func (l *SinglyLinkedList[T]) Get(idx int) (T, bool) {
	curent := l.head
	for i := 0; i < idx; i++ {
		curent = curent.next
	}
	if curent != nil {
		return curent.value, true
	}
	return dsa.ZeroValue[T](), false
}

func (l *SinglyLinkedList[T]) RemoveAt(idx int) (T, bool) {
	if idx == 0 {
		removed := l.head
		l.head = removed.next
		l.length--
		return removed.value, true
	}

	curent := l.head
	for i := 0; i < idx-1; i++ {
		curent = curent.next
	}
	if curent.next != nil {
		l.length--
		removed := curent.next
		curent.next = removed.next
		return removed.value, true
	}

	return dsa.ZeroValue[T](), false
}

func (l *SinglyLinkedList[T]) print() {
	fmt.Printf("Length: %d\n", l.length)
	fmt.Print("[")
	current := l.head
	for current.next != nil {
		fmt.Printf("%v, ", current.value)
		current = current.next
	}
	fmt.Printf("%v]\n\n", current.value)

}

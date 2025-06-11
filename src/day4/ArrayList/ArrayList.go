package arraylist

import (
	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
)

type ArrayList[T comparable] struct {
	length   int
	capacity int
	entries  []T
}

func (l *ArrayList[T]) Len() int { return l.length }

func NewArrayList[T comparable](capacity int) *ArrayList[T] {
	return &ArrayList[T]{length: 0, entries: make([]T, capacity), capacity: capacity}
}

func (l *ArrayList[T]) shiftUpFrom(idx int) {
	for i := l.length; i > idx; i-- {
		l.entries[i] = l.entries[i-1]
	}
	l.length++
	if l.length >= l.capacity {
		l.growCapacity()
	}
}

func (l *ArrayList[T]) shiftDownTo(idx int) {
	for i := idx; i < l.length-1; i++ {
		l.entries[i] = l.entries[i+1]
	}
	l.length--
}

func (l *ArrayList[T]) growCapacity() {
	newCap := l.capacity * 2
	newEntries := make([]T, newCap)
	copy(newEntries, l.entries)
	l.capacity = newCap
	l.entries = newEntries
}

func (l *ArrayList[T]) Prepend(item T) {
	l.shiftUpFrom(0)
	l.entries[0] = item
}

func (l *ArrayList[T]) InsertAt(item T, idx int) {
	l.shiftUpFrom(idx)
	l.entries[idx] = item

}

func (l *ArrayList[T]) Append(item T) {
	if l.length == l.capacity {
		l.growCapacity()
	}
	l.entries[l.length] = item
	l.length++
}

func (l *ArrayList[T]) Remove(item T) (T, bool) {
	for idx := 0; idx < l.length; idx++ {
		if l.entries[idx] == item {
			removed := l.entries[idx]
			l.shiftDownTo(idx)
			return removed, true
		}
	}
	return dsa.ZeroValue[T](), false
}

func (l *ArrayList[T]) Get(idx int) (T, bool) {
	if idx > l.length {
		return dsa.ZeroValue[T](), false
	}
	return l.entries[idx], true
}

func (l *ArrayList[T]) RemoveAt(idx int) (T, bool) {
	if idx >= l.length {
		return dsa.ZeroValue[T](), false
	}
	removed := l.entries[idx]
	l.shiftDownTo(idx)
	return removed, true
}

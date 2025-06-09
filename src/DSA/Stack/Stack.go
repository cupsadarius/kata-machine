package stack

type Stack[T comparable] struct {
	Length int
}

func NewStack[T comparable]() *Stack[T] {

}

func (s *Stack[T]) Push(item T) {

}

func (s *Stack[T]) Pop() (T, bool) {

}

func (s *Stack[T]) Peek() (T, bool) {

}

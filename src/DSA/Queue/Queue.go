package queue

type Queue[T comparable] struct {
	Length int
}

func NewQueue[T comparable]() *Queue[T] {

}

func (q *Queue[T]) Enqueue(item T) {

}

func (q *Queue[T]) Deque() (T, bool) {

}

func (q *Queue[T]) Peek() (T, bool) {

}

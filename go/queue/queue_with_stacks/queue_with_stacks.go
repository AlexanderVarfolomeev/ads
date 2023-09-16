package queue_with_stacks

import "errors"

type Queue[T any] struct {
	s1 Stack[T]
	s2 Stack[T]
}

func (q *Queue[T]) Size() int {
	return q.s1.Size() + q.s2.Size()
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T

	if q.s1.Size() == 0 && q.s2.Size() == 0 {
		return result, errors.New("queue is empty")
	}

	if q.s2.Size() != 0 {
		return q.s2.Pop()
	}

	for q.s1.Size() != 0 {
		val, _ := q.s1.Pop()
		q.s2.Push(val)
	}

	return q.s2.Pop()
}

func (q *Queue[T]) Enqueue(itm T) {
	q.s1.Push(itm)
}

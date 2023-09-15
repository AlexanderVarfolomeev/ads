package main

import "errors"

//import (
//	// "os"
//	//	"fmt" включите если используете
//)

type Node[T any] struct {
	next  *Node[T]
	value T
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T

	if q.head == nil {
		return result, errors.New("queue is empty")
	}

	result = q.head.value
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	q.size--

	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
	n := &Node[T]{value: itm}
	if q.head == nil {
		q.head = n
	} else {
		q.tail.next = n
	}

	q.tail = n
	q.size++
}

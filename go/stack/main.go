package main

import "errors"

type Node[T any] struct {
	next  *Node[T]
	value T
}

type Stack[T any] struct {
	head *Node[T]
}

func (st *Stack[T]) Size() int {
	size := 0
	n := st.head

	for n != nil {
		n = n.next
		size++
	}

	return size
}

func (st *Stack[T]) Peek() (T, error) {
	var result T

	if st.head == nil {
		return result, errors.New("stack is empty")
	}

	return st.head.value, nil
}

func (st *Stack[T]) Pop() (T, error) {
	var result T

	if st.head == nil {
		return result, errors.New("stack is empty")
	}

	result = st.head.value
	st.head = st.head.next

	return result, nil
}

func (st *Stack[T]) Push(itm T) {
	n := &Node[T]{value: itm}
	n.next = st.head
	st.head = n
}

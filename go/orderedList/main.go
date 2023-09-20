package main

import (
	"asd/orderedList/constraints"
	"errors"
)

func main() {
}

type Node[T constraints.Ordered] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type OrderedList[T constraints.Ordered] struct {
	head       *Node[T]
	tail       *Node[T]
	_ascending bool

	count int
}

func (l *OrderedList[T]) Count() int {
	return l.count
}

func (l *OrderedList[T]) Add(item T) {
	newNode := Node[T]{value: item}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	n := l.head
	for n != nil {
		if l.Compare(n.value, item) >= 0 {

			if n == l.head {
				newNode.next = l.head
				l.head.prev = &newNode
				l.head = &newNode
			} else {
				newNode.prev = n.prev
				n.prev = &newNode

				n.prev.next = &newNode
				newNode.next = n
			}
			return
		}
		n = n.next
	}

	newNode.prev = l.tail
	l.tail.next = &newNode
	l.tail = &newNode

	l.count++

}

func (l *OrderedList[T]) addAscending(item T) {
	newNode := Node[T]{value: item}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	n := l.head
	for n != nil {
		if l.Compare(n.value, item) >= 0 {

			if n == l.head {
				newNode.next = l.head
				l.head.prev = &newNode
				l.head = &newNode
			} else {
				newNode.prev = n.prev
				n.prev = &newNode

				n.prev.next = &newNode
				newNode.next = n
			}
			return
		}
		n = n.next
	}

	newNode.prev = l.tail
	l.tail.next = &newNode
	l.tail = &newNode
}

func (l *OrderedList[T]) addDescending(item T) {
	newNode := Node[T]{value: item}
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	n := l.head
	for n != nil {
		if l.Compare(n.value, item) < 0 {

			if n == l.head {
				newNode.next = l.head
				l.head.prev = &newNode
				l.head = &newNode
			} else {
				newNode.prev = n.prev
				n.prev = &newNode

				n.prev.next = &newNode
				newNode.next = n
			}
			return
		}
		n = n.next
	}

	newNode.prev = l.tail
	l.tail.next = &newNode
	l.tail = &newNode
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
	if l._ascending {
		return l.findAscending(n)
	}
	return l.findDescending(n)
}

func (l *OrderedList[T]) findAscending(n T) (Node[T], error) {
	node := l.head

	for node != nil {
		if l.Compare(node.value, n) > 0 {
			return Node[T]{value: n, next: nil, prev: nil}, errors.New("node not found")
		}

		if l.Compare(node.value, n) == 0 {
			return *node, nil
		}

		node = node.next
	}

	return Node[T]{value: n, next: nil, prev: nil}, errors.New("node not found")
}

func (l *OrderedList[T]) findDescending(n T) (Node[T], error) {
	node := l.head

	for node != nil {
		if l.Compare(node.value, n) < 0 {
			return Node[T]{value: n, next: nil, prev: nil}, errors.New("node not found")
		}

		if l.Compare(node.value, n) == 0 {
			return *node, nil
		}

		node = node.next
	}

	return Node[T]{value: n, next: nil, prev: nil}, errors.New("node not found")
}

func (l *OrderedList[T]) Delete(n T) {
	node := l.head
	l.count--

	for node != nil {
		if l.Compare(node.value, n) == 0 {

			if node == l.head {
				l.head = l.head.next
				if l.head != nil {
					l.head.prev = nil
				}
			} else {
				node.prev.next = node.next

				if node.next != nil {
					node.next.prev = node.prev
				} else {
					l.tail = node.prev
				}
			}

			if l.head == nil {
				l.tail = nil
			}
			return
		}

		node = node.next
	}

}

func (l *OrderedList[T]) Clear(asc bool) {
	l._ascending = asc
	l.head, l.tail = nil, nil
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
	if v1 < v2 {
		return -1
	}
	if v1 > v2 {
		return +1
	}
	return 0
	// -1 если v1 < v2
	// 0 если v1 == v2
	// +1 если v1 > v2
}

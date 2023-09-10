package main

import "errors"

type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) Count() int {
	count := 0
	node := l.head

	for node != nil {
		node = node.next
		count++
	}

	return count
}

func (l *LinkedList2) Find(n int) (Node, error) {
	node := l.head

	for node != nil {
		if node.value == n {
			return *node, nil
		}
		node = node.next
	}

	return Node{}, errors.New("node not found")
}

func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node

	node := l.head

	for node != nil {
		if node.value == n {
			nodes = append(nodes, *node)
		}
		node = node.next
	}

	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	deleted := false

	for l.head != nil && l.head.value == n {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		deleted = true
		if !all {
			break
		}
	}

	if l.head == nil {
		l.tail = nil
	}

	if deleted && !all {
		return
	}

	for node := l.head; node != nil; node = node.next {
		if node.value == n {
			node.prev.next = node.next

			if node.next != nil {
				node.next.prev = node.prev
			} else {
				l.tail = node.prev
			}

			if !all {
				return
			}
		}
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	node := l.head

	for node != nil {
		if node == after {
			add.next = node.next

			if node.next != nil {
				node.next.prev = &add
			}

			node.next = &add
			add.prev = node
			if node == l.tail {
				l.tail = &add
			}

			break
		}

		node = node.next
	}
}

func (l *LinkedList2) InsertFirst(first Node) {
	if l.head != nil {
		l.head.prev = &first
	} else {
		l.tail = &first
	}

	first.next = l.head
	l.head = &first
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
}
